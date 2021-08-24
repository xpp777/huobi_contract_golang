package api

import (
	"container/list"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/xiaomingping/huobi_contract_golang/src/config"
	"github.com/xiaomingping/huobi_contract_golang/src/logh"
	"github.com/xiaomingping/huobi_contract_golang/src/utils/reqbuilder"
	"github.com/xiaomingping/huobi_contract_golang/src/utils/zip"
	"github.com/xiaomingping/huobi_contract_golang/src/ws/request"
	"reflect"
	"strings"
	"sync"
	"time"
)

type closeChanCallback chan struct{}

type CallbackFunc struct {
	fun   interface{}
	param reflect.Type
}

func (wsOp *CallbackFunc) Init(fun interface{}, param reflect.Type) *CallbackFunc {
	wsOp.fun = fun
	wsOp.param = param
	return wsOp
}

type WebSocketOp struct {
	host              string                   // 连接地址
	path              string                   // 路由
	conn              *websocket.Conn          // 基础连接
	accessKey         string                   // a key
	secretKey         string                   // s key
	closeChan         chan byte                // 关闭chan
	isClose           bool                     // 是否关闭
	authOk            bool                     // 是否权限验证
	mutex             sync.Mutex               // 锁
	readChan          chan []byte              //读消息管道
	writeChan         chan []byte              //写消息管道
	SubTopicMap       map[string]*CallbackFunc // 订阅历史
	ReqTopicMap       map[string]*CallbackFunc // 请求历史
	authConnect       int                      // 自动重连次数
	closeChanCallback chan struct{}            // 关闭通知外部
	allSubStrs        list.List
}

func (ws *WebSocketOp) open(path, host, accessKey, secretKey string, closeChanCallback closeChanCallback) bool {
	if host == "" {
		ws.host = config.LINEAR_SWAP_DEFAULT_HOST
	}
	ws.host = host
	ws.path = path
	ws.accessKey = accessKey
	ws.secretKey = secretKey
	ws.closeChan = make(chan byte, 1)
	ws.readChan = make(chan []byte, 30)
	ws.writeChan = make(chan []byte, 10)
	ws.allSubStrs = list.List{}
	ws.SubTopicMap = make(map[string]*CallbackFunc)
	ws.ReqTopicMap = make(map[string]*CallbackFunc)
	ws.authConnect = 5
	ws.closeChanCallback = closeChanCallback
	ret := ws.connServer()
	return ret
}
func (ws *WebSocketOp) Close() {
	ws.conn.Close()
	ws.mutex.Lock()
	defer ws.mutex.Unlock()
	if !ws.isClose {
		ws.isClose = true
		ws.authOk = false
		close(ws.closeChan)
	}
	ws.closeChanCallback <- struct{}{}
}

func (ws *WebSocketOp) sendAuth(host string, path string, accessKey string, secretKey string) bool {
	timestamp := time.Now().UTC().Format("2006-01-02T15:04:05")
	req := new(reqbuilder.GetRequest).Init()
	req.AddParam("AccessKeyId", accessKey)
	req.AddParam("SignatureMethod", "HmacSHA256")
	req.AddParam("SignatureVersion", "2")
	req.AddParam("Timestamp", timestamp)
	sign := new(reqbuilder.Signer).Init(secretKey)
	signature := sign.Sign("GET", host, path, req.BuildParams())
	auth := request.WSAuthData{
		Op:               "auth",
		AtType:           "api",
		AccessKeyId:      accessKey,
		SignatureMethod:  "HmacSHA256",
		SignatureVersion: "2",
		Timestamp:        timestamp,
		Signature:        signature}

	data, error := json.Marshal(&auth)
	if error != nil {
		logh.Error("Auth to json error.")
		return false
	}
	ws.SendMsg(data)
	return true
}

func (ws *WebSocketOp) connServer() bool {
	url := fmt.Sprintf("wss://%s%s", ws.host, ws.path)
	var err error
	ws.conn, _, err = websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		logh.Error("WebSocket connected error: %s", err)
		return false
	}
	logh.Info("WebSocket connected")
	go ws.readLoop()
	go ws.writeLoop()
	go ws.DoHand()
	if ws.accessKey == "" && ws.secretKey == "" {
		ws.authOk = true
		return true
	}
	ws.authOk = false
	return ws.sendAuth(ws.host, ws.path, ws.accessKey, ws.secretKey)
}

// 读协程
func (ws *WebSocketOp) readLoop() {
	for {
		msgType, buf, err := ws.conn.ReadMessage()
		if err != nil {
			logh.Debug(err.Error())
			goto ERR
		}
		var message []byte
		if msgType == websocket.BinaryMessage {
			message, err = zip.GZipDecompress(buf)
			if err != nil {
				logh.Error("UnGZip data error: %s", err)
				continue
			}
		} else if msgType == websocket.TextMessage {
			message = buf
		}
		select {
		case ws.readChan <- message:
		case <-ws.closeChan:
			goto ERR
		}
	}
ERR:
	logh.Debug("read 已关闭 %s",ws.accessKey)
	ws.Close()
	close(ws.readChan)
}

//写协程
func (ws *WebSocketOp) writeLoop() {
	for {
		select {
		case msg, ok := <-ws.writeChan:
			if !ok {
				goto ERR
			}
			if err := ws.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				logh.Debug(err.Error())
				goto ERR
			}
		case <-ws.closeChan:
			goto ERR
		}

	}
ERR:
	logh.Debug("write 已关闭 %s",ws.accessKey)
	ws.Close()
	close(ws.writeChan)
}

// 通用订阅
func (ws *WebSocketOp) sub(subReq []byte, ch string, fun interface{}, param reflect.Type) bool {
	for !ws.authOk {
		time.Sleep(10)
	}
	ch = strings.ToLower(ch)
	ws.mutex.Lock()
	defer ws.mutex.Unlock()
	var mi *CallbackFunc = nil
	if _, found := ws.SubTopicMap[ch]; found {
		mi = new(CallbackFunc).Init(fun, param)
		ws.SubTopicMap[ch] = mi
		return true
	}
	logh.Info("websocket has send data: %s", string(subReq))
	mi = new(CallbackFunc).Init(fun, param)
	ws.SubTopicMap[ch] = mi
	ws.allSubStrs.PushBack(string(subReq))
	return ws.SendMsg(subReq)
}
func (ws *WebSocketOp) req(subReq []byte, ch string, fun interface{}, param reflect.Type) bool {
	for !ws.authOk {
		time.Sleep(10)
	}
	ch = strings.ToLower(ch)
	ws.mutex.Lock()
	defer ws.mutex.Unlock()
	if mi, found := ws.ReqTopicMap[ch]; found {
		mi = new(CallbackFunc).Init(fun, param)
		ws.ReqTopicMap[ch] = mi
		return true
	}
	mi := new(CallbackFunc).Init(fun, param)
	ws.ReqTopicMap[ch] = mi
	return ws.SendMsg(subReq)
}

// 通用取消订阅
func (ws *WebSocketOp) unsub(unsubReq []byte, ch string) bool {
	for !ws.authOk {
		time.Sleep(10)
	}
	ch = strings.ToLower(ch)
	ws.mutex.Lock()
	defer ws.mutex.Unlock()
	if _, found := ws.SubTopicMap[ch]; !found {
		return true
	}
	delete(ws.SubTopicMap, ch)
	logh.Info("websocket has send data: %s", unsubReq)
	var next *list.Element
	for e := ws.allSubStrs.Front(); e != nil; e = next {
		next = e.Next()
		val := e.Value.(string)
		if val == string(unsubReq) {
			ws.allSubStrs.Remove(e)
		}
	}
	return ws.SendMsg(unsubReq)
}

func (ws *WebSocketOp) SendMsg(data []byte) bool {
	ws.writeChan <- data
	return true
}

func (ws *WebSocketOp) DoHand() {
	for {
		select {
		case msg, ok := <-ws.readChan:
			if !ok {
				continue
			}
			ws.msgHand(msg)
		case <-ws.closeChan:
			goto ERR
		}
	}
ERR:
	logh.Debug("do 已关闭 key:%s",ws.accessKey)
	ws.Close()
}

func (ws *WebSocketOp) msgHand(message []byte) {
	var data map[string]interface{}
	err := json.Unmarshal(message, &data)
	if err != nil {
		logh.Error("msg to map[string]json.RawMessage error: %s", err)
		return
	}
	if ts, found := data["ping"]; found {
		ts = int64(ts.(float64))
		pongData := fmt.Sprintf("{\"pong\":%d }", ts)
		ws.SendMsg([]byte(pongData))
		return
	}
	if _, found := data["op"]; found {
		ws.opHand(data, message)
		return
	} else if topic, found := data["subbed"]; found { // sub success reply
		logh.Info("\"subbed\": \"%s\"", topic)
	} else if topic, found := data["unsubbed"]; found { // unsub success reply
		logh.Info("\"unsubbed\": \"%s\"", topic)
		if _, found := ws.SubTopicMap[topic.(string)]; found {
			delete(ws.SubTopicMap, topic.(string))
		}
	} else if topic, found := data["ch"]; found { // market sub reply data
		ws.handleSubCallbackFun(topic.(string), message, data)
	} else if topic, found := data["rep"]; found { // market request reply data
		ws.handleReqCallbackFun(topic.(string), message, data)
	} else if code, found := data["err-code"]; found { // market request reply data
		code = code
		msg := data["err-msg"]
		logh.Error("%d:%s", code, msg)
	} else {
		logh.Info("WebSocket received unknow data: %s", data)
	}
}

func (ws *WebSocketOp) opHand(data map[string]interface{}, message []byte) {
	switch data["op"] {
	case "ping":
		ts := data["ts"]
		pongData := fmt.Sprintf("{ \"op\":\"pong\", \"ts\": \"%s\" }", ts)
		ws.SendMsg([]byte(pongData))
	case "close":
		logh.Error("Some error occurres when authentication in server side.")
	case "error":
		logh.Error("Illegal op or internal error, but websoket is still connected.")
	case "auth":
		code := int64(data["err-code"].(float64))
		if code == 0 {
			logh.Info("Authentication success.")
			ws.authOk = true
			for e := ws.allSubStrs.Front(); e != nil; e = e.Next() {
				ws.SendMsg([]byte(e.Value.(string)))
			}
		} else {
			msg := data["err-msg"].(string)
			logh.Error("Authentication failure: %d/%s", code, msg)
			ws.Close()
		}
	case "notify":
		topic := data["topic"].(string)
		ws.handleSubCallbackFun(topic, message, data)
	case "sub":
		topic := data["topic"]
		logh.Info("sub: \"%s\"", topic)
	case "unsub":
		topic := data["topic"].(string)
		logh.Info("unsub: \"%s\"", topic)
		if _, found := ws.SubTopicMap[topic]; found {
			delete(ws.SubTopicMap, topic)
		}
	default:
		logh.Info("WebSocket received unknow data: %s", data)
	}
}

func (ws *WebSocketOp) handleSubCallbackFun(ch string, data []byte, jdata map[string]interface{}) {
	var mi *CallbackFunc = nil
	ch = strings.ToLower(ch)

	if _, found := ws.SubTopicMap[ch]; found {
		mi = ws.SubTopicMap[ch]
	} else if ch == "accounts" || ch == "positions" { // isolated
		data_array := jdata["data"].([]interface{})
		contract_code := data_array[0].(map[string]interface{})["contract_code"].(string)
		contract_code = strings.ToLower(contract_code)
		full_ch := fmt.Sprintf("%s.%s", ch, contract_code)
		if _, found := ws.SubTopicMap[full_ch]; found {
			mi = ws.SubTopicMap[full_ch]
		} else if _, found := ws.SubTopicMap[fmt.Sprintf("%s.*", ch)]; found {
			mi = ws.SubTopicMap[fmt.Sprintf("%s.*", ch)]
		}
	} else if strings.HasPrefix(ch, "orders.") {
		if _, found := ws.SubTopicMap["orders.*"]; found {
			mi = ws.SubTopicMap["orders.*"]
		}
	} else if strings.HasPrefix(ch, "matchorders.") {
		if _, found := ws.SubTopicMap["matchorders.*"]; found {
			mi = ws.SubTopicMap["matchorders.*"]
		}
	} else if strings.HasPrefix(ch, "trigger_order.") {
		if _, found := ws.SubTopicMap["trigger_order.*"]; found {
			mi = ws.SubTopicMap["trigger_order.*"]
		}
	} else if ch == "accounts_cross" { // isolated
		data_array := jdata["data"].([]interface{})
		margin_account := data_array[0].(map[string]interface{})["margin_account"].(string)
		margin_account = strings.ToLower(margin_account)

		full_ch := fmt.Sprintf("%s.%s", ch, margin_account)
		if _, found := ws.SubTopicMap[full_ch]; found {
			mi = ws.SubTopicMap[full_ch]
		} else if _, found := ws.SubTopicMap[fmt.Sprintf("%s.*", ch)]; found {
			mi = ws.SubTopicMap[fmt.Sprintf("%s.*", ch)]
		}
	} else if ch == "positions_cross" {
		data_array := jdata["data"].([]interface{})
		contract_code := data_array[0].(map[string]interface{})["contract_code"].(string)
		contract_code = strings.ToLower(contract_code)

		full_ch := fmt.Sprintf("%s.%s", ch, contract_code)
		if _, found := ws.SubTopicMap[full_ch]; found {
			mi = ws.SubTopicMap[full_ch]
		} else if _, found := ws.SubTopicMap[fmt.Sprintf("%s.*", ch)]; found {
			mi = ws.SubTopicMap[fmt.Sprintf("%s.*", ch)]
		}
	} else if strings.HasPrefix(ch, "orders_cross.") {
		if _, found := ws.SubTopicMap["orders_cross.*"]; found {
			mi = ws.SubTopicMap["orders_cross.*"]
		}
	} else if strings.HasPrefix(ch, "matchorders_cross.") {
		if _, found := ws.SubTopicMap["matchorders_cross.*"]; found {
			mi = ws.SubTopicMap["matchorders_cross.*"]
		}
	} else if strings.HasPrefix(ch, "trigger_order_cross.") {
		if _, found := ws.SubTopicMap["trigger_order_cross.*"]; found {
			mi = ws.SubTopicMap["trigger_order_cross.*"]
		}
	} else if strings.HasSuffix(ch, ".liquidation_orders") { // General
		if _, found := ws.SubTopicMap["public.*.liquidation_orders"]; found {
			mi = ws.SubTopicMap["public.*.liquidation_orders"]
		}
	} else if strings.HasSuffix(ch, ".funding_rate") {
		if _, found := ws.SubTopicMap["public.*.funding_rate"]; found {
			mi = ws.SubTopicMap["public.*.funding_rate"]
		}
	} else if strings.HasSuffix(ch, ".contract_info") {
		if _, found := ws.SubTopicMap["public.*.contract_info"]; found {
			mi = ws.SubTopicMap["public.*.contract_info"]
		}
	}
	if mi == nil {
		logh.Error("no callback function to handle: %s", jdata)
		return
	}
	ws.runFunction(mi, data)
}

func (ws *WebSocketOp) handleReqCallbackFun(ch string, data []byte, jdata map[string]interface{}) {
	ch = strings.ToLower(ch)
	mi, found := ws.ReqTopicMap[ch]
	if !found {
		logh.Error("no callback function to handle: %s", jdata)
		return
	}
	if mi == nil {
		logh.Error("no callback function to handle: %s", jdata)
		return
	}
	ws.runFunction(mi, data)
}

func (ws *WebSocketOp) runFunction(mi *CallbackFunc, data []byte) {
	param := reflect.New(mi.param).Interface()
	err := json.Unmarshal(data, &param)
	if err != nil {
		logh.Debug("msg unmarshal err :", err)
		return
	}
	rargs := make([]reflect.Value, 1)
	rargs[0] = reflect.ValueOf(param)
	fun := reflect.ValueOf(mi.fun)
	fun.Call(rargs)
}
