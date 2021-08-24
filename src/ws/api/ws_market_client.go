package api

import (
	"encoding/json"
	"fmt"
	"github.com/xiaomingping/huobi_contract_golang/src/config"
	"github.com/xiaomingping/huobi_contract_golang/src/ws/request"
	"github.com/xiaomingping/huobi_contract_golang/src/ws/response"
	"reflect"
)

type WSMarketClient struct {
	WebSocketOp
}

func (wsMk *WSMarketClient) Init(host string, callback closeChanCallback) *WSMarketClient {
	if host == "" {
		host = config.LINEAR_SWAP_DEFAULT_HOST
	}
	wsMk.open("/linear-swap-ws", host, "", "", callback)
	return wsMk
}

/**
【通用】订阅KLine 数据
*/
type OnSubKLineResponse func(*response.SubKLineResponse)
type OnReqKLineResponse func(*response.ReqKLineResponse)

func (wsMk *WSMarketClient) SubKLine(contractCode string, period string, callbackFun OnSubKLineResponse, id string) {
	if id == "" {
		id = config.DEFAULT_ID
	}
	ch := fmt.Sprintf("market.%s.kline.%s", contractCode, period)
	subData := request.WSSubData{Sub: ch, Id: id}
	data, _ := json.Marshal(subData)
	wsMk.sub(data, ch, callbackFun, reflect.TypeOf(response.SubKLineResponse{}))
}

/**
【通用】请求 KLine 数据
*/
func (wsMk *WSMarketClient) ReqKLine(contractCode string, period string, callbackFun OnReqKLineResponse, from int64, to int64, id string) {
	if id == "" {
		id = config.DEFAULT_ID
	}
	ch := fmt.Sprintf("market.%s.kline.%s", contractCode, period)
	subData := request.WSReqData{Req: ch, From: from, To: to, Id: id}
	data, _ := json.Marshal(subData)
	wsMk.req(data, ch, callbackFun, reflect.TypeOf(response.ReqKLineResponse{}))
}

/**
【通用】订阅 Market Depth 数据
*/
type OnSubDepthResponse func(*response.SubDepthResponse)

func (wsMk *WSMarketClient) SubDepth(contractCode string, fcType string, callbackFun OnSubDepthResponse, id string) {
	if id == "" {
		id = config.DEFAULT_ID
	}
	ch := fmt.Sprintf("market.%s.depth.%s", contractCode, fcType)
	subData := request.WSSubData{Sub: ch, Id: id}
	data, _ := json.Marshal(subData)
	wsMk.sub(data, ch, callbackFun, reflect.TypeOf(response.SubDepthResponse{}))
}

/**
【通用】订阅Market Depth增量数据
*/
func (wsMk *WSMarketClient) SubIncrementalDepth(contractCode string, size string, callbackFun OnSubDepthResponse, id string) {
	if id == "" {
		id = config.DEFAULT_ID
	}
	ch := fmt.Sprintf("market.%s.depth.size_%s.high_freq", contractCode, size)
	subData := request.WSSubData{Sub: ch, Id: id}
	data, _ := json.Marshal(subData)
	wsMk.sub(data, ch, callbackFun, reflect.TypeOf(response.SubDepthResponse{}))
}

/**
【通用】订阅 Market Detail 数据
*/
type OnSubDetailResponse func(*response.SubKLineResponse)

func (wsMk *WSMarketClient) SubDetail(contractCode string, callbackFun OnSubDetailResponse, id string) {
	if id == "" {
		id = config.DEFAULT_ID
	}
	ch := fmt.Sprintf("market.%s.detail", contractCode)
	subData := request.WSSubData{Sub: ch, Id: id}
	data, _ := json.Marshal(subData)
	wsMk.sub(data, ch, callbackFun, reflect.TypeOf(response.SubKLineResponse{}))
}

/**
【通用】订阅买一卖一逐笔行情推送
*/
type OnSubBBOResponse func(*response.SubBBOResponse)

func (wsMk *WSMarketClient) SubBBO(contractCode string, callbackFun OnSubBBOResponse, id string) {
	if id == "" {
		id = config.DEFAULT_ID
	}
	ch := fmt.Sprintf("market.%s.bbo", contractCode)
	subData := request.WSSubData{Sub: ch, Id: id}
	data, _ := json.Marshal(subData)
	wsMk.sub(data, ch, callbackFun, reflect.TypeOf(response.SubBBOResponse{}))
}

type OnSubTradeDetailResponse func(*response.SubTradeDetailResponse)
type OnReqTradeDetailResponse func(*response.ReqTradeDetailResponse)

/**
【通用】订阅 Trade Detail 数据
*/
func (wsMk *WSMarketClient) SubTradeDetail(contractCode string, callbackFun OnSubTradeDetailResponse, id string) {
	if id == "" {
		id = config.DEFAULT_ID
	}
	ch := fmt.Sprintf("market.%s.trade.detail", contractCode)
	subData := request.WSSubData{Sub: ch, Id: id}
	data, _ := json.Marshal(subData)
	wsMk.sub(data, ch, callbackFun, reflect.TypeOf(response.SubTradeDetailResponse{}))
}

/**
【通用】请求 Trade Detail 数据
*/
func (wsMk *WSMarketClient) ReqTradeDetail(contractCode string, callbackFun OnReqTradeDetailResponse, id string) {
	if id == "" {
		id = config.DEFAULT_ID
	}
	ch := fmt.Sprintf("market.%s.trade.detail", contractCode)
	subData := request.WSReqData{Req: ch, Id: id}
	data, _ := json.Marshal(subData)
	wsMk.req(data, ch, callbackFun, reflect.TypeOf(response.ReqTradeDetailResponse{}))
}
