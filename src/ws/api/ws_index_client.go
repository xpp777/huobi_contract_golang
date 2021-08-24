package api

import (
	"encoding/json"
	"fmt"
	"github.com/xiaomingping/huobi_contract_golang/src/config"
	"github.com/xiaomingping/huobi_contract_golang/src/ws/request"
	"github.com/xiaomingping/huobi_contract_golang/src/ws/response"
	"reflect"
)

type WSIndexClient struct {
	WebSocketOp
}

func (wsIx *WSIndexClient) Init(host string, callback closeChanCallback) *WSIndexClient {
	if host == "" {
		host = config.LINEAR_SWAP_DEFAULT_HOST
	}
	wsIx.open("/ws_index", host, "", "", callback)
	return wsIx
}

/**
【通用】订阅溢价指数K线数据
*/
type OnSubPremiumIndexKLineResponse func(*response.SubIndexKLineResponse)
type OnReqPremiumIndexKLineResponse func(*response.ReqIndexKLineResponse)

func (wsIx *WSIndexClient) SubPremiumIndexKLine(contractCode string, period string, callbackFun OnSubPremiumIndexKLineResponse, id string) {
	if id == "" {
		id = config.DEFAULT_ID
	}
	ch := fmt.Sprintf("market.%s.premium_index.%s", contractCode, period)
	subData := request.WSSubData{Sub: ch, Id: id}
	data, _ := json.Marshal(subData)
	wsIx.sub(data, ch, callbackFun, reflect.TypeOf(response.SubIndexKLineResponse{}))
}

/**
【通用】请求溢价指数K线数据
*/
func (wsIx *WSIndexClient) ReqPremiumIndexKLine(contractCode string, period string, callbackFun OnReqPremiumIndexKLineResponse, from int64, to int64, id string) {
	if id == "" {
		id = config.DEFAULT_ID
	}
	ch := fmt.Sprintf("market.%s.premium_index.%s", contractCode, period)
	reqData := request.WSReqData{Req: ch, Id: id, From: from, To: to}
	data, _ := json.Marshal(reqData)
	wsIx.req(data, ch, callbackFun, reflect.TypeOf(response.ReqIndexKLineResponse{}))
}

/**
【通用】订阅标记价格K线数据
*/
type OnSubMarkPriceKLineResponse func(*response.SubIndexKLineResponse)
type OnReqMarkPriceKLineResponse func(*response.ReqIndexKLineResponse)

func (wsIx *WSIndexClient) SubMarkPriceKLine(contractCode string, period string, callbackFun OnSubMarkPriceKLineResponse, id string) {
	if id == "" {
		id = config.DEFAULT_ID
	}
	ch := fmt.Sprintf("market.%s.mark_price.%s", contractCode, period)
	subData := request.WSSubData{Sub: ch, Id: id}
	data, _ := json.Marshal(subData)
	wsIx.sub(data, ch, callbackFun, reflect.TypeOf(response.SubIndexKLineResponse{}))
}

/**
【通用】请求标记价格K线数据
*/
func (wsIx *WSIndexClient) ReqMarkPriceKLine(contractCode string, period string, callbackFun OnReqMarkPriceKLineResponse, from int64, to int64, id string) {
	if id == "" {
		id = config.DEFAULT_ID
	}
	ch := fmt.Sprintf("market.%s.mark_price.%s", contractCode, period)
	reqData := request.WSReqData{Req: ch, Id: id, From: from, To: to}
	data, _ := json.Marshal(reqData)
	wsIx.req(data, ch, callbackFun, reflect.TypeOf(response.ReqIndexKLineResponse{}))
}

/**
【通用】订阅预测资金费率K线数据
*/
type OnSubEstimatedRateResponse func(*response.SubIndexKLineResponse)
type OnReqEstimatedRateResponse func(*response.ReqIndexKLineResponse)

func (wsIx *WSIndexClient) SubEstimatedRateKLine(contractCode string, period string, callbackFun OnSubEstimatedRateResponse, id string) {
	if id == "" {
		id = config.DEFAULT_ID
	}
	ch := fmt.Sprintf("market.%s.estimated_rate.%s", contractCode, period)
	subData := request.WSSubData{Sub: ch, Id: id}
	data, _ := json.Marshal(subData)
	wsIx.sub(data, ch, callbackFun, reflect.TypeOf(response.SubIndexKLineResponse{}))
}

/**
【通用】请求预测资金费率K线数据
*/
func (wsIx *WSIndexClient) ReqEstimatedRateKLine(contractCode string, period string, callbackFun OnReqEstimatedRateResponse, from int64, to int64, id string) {
	if id == "" {
		id = config.DEFAULT_ID
	}
	ch := fmt.Sprintf("market.%s.estimated_rate.%s", contractCode, period)
	reqData := request.WSReqData{Req: ch, Id: id, From: from, To: to}
	data, _ := json.Marshal(reqData)
	wsIx.req(data, ch, callbackFun, reflect.TypeOf(response.ReqIndexKLineResponse{}))
}

/**
【通用】订阅基差数据
*/
type OnSubBasisResponse func(*response.SubBasiesResponse)
type OnReqBasisResponse func(*response.ReqBasisResponse)

func (wsIx *WSIndexClient) SubBasis(contractCode string, period string, callbackFun OnSubBasisResponse, basisPriceType string, id string) {
	if basisPriceType == "" {
		basisPriceType = "open"
	}
	if id == "" {
		id = config.DEFAULT_ID
	}
	ch := fmt.Sprintf("market.%s.basis.%s.%s", contractCode, period, basisPriceType)
	subData := request.WSSubData{Sub: ch, Id: id}
	data, _ := json.Marshal(subData)
	wsIx.sub(data, ch, callbackFun, reflect.TypeOf(response.SubBasiesResponse{}))
}

/**
【通用】请求基差数据
*/
func (wsIx *WSIndexClient) ReqBasis(contractCode string, period string, callbackFun OnReqBasisResponse, from int64, to int64,
	basisPriceType string, id string) {
	if basisPriceType == "" {
		basisPriceType = "open"
	}
	if id == "" {
		id = config.DEFAULT_ID
	}
	ch := fmt.Sprintf("market.%s.basis.%s.%s", contractCode, period, basisPriceType)
	reqData := request.WSReqData{Req: ch, Id: id, From: from, To: to}
	data, _ := json.Marshal(reqData)
	wsIx.req(data, ch, callbackFun, reflect.TypeOf(response.ReqBasisResponse{}))
}
