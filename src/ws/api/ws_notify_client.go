package api

import (
	"encoding/json"
	"fmt"
	"github.com/xiaomingping/huobi_contract_golang/src/config"
	"github.com/xiaomingping/huobi_contract_golang/src/ws/request"
	"github.com/xiaomingping/huobi_contract_golang/src/ws/response"
	"reflect"
	"strings"
)

type WSNotifyClient struct {
	WebSocketOp
}

func (wsNf *WSNotifyClient) Init(accessKey string, secretKey string, host string, callback closeChanCallback) *WSNotifyClient {
	if host == "" {
		host = config.LINEAR_SWAP_DEFAULT_HOST
	}
	wsNf.open("/linear-swap-notification", host, accessKey, secretKey, callback)
	return wsNf
}

// -------------------------------------------------------------
/**
【逐仓】订阅订单成交数据（sub）
*/

type OnSubOrdersResponse func(*response.SubOrdersResponse)

func (wsNf *WSNotifyClient) IsolatedSubOrders(contractCode string, callbackFun OnSubOrdersResponse, cid string) {
	if cid == "" {
		cid = config.DEFAULT_CID
	}
	ch := fmt.Sprintf("orders.%s", contractCode)
	opData := request.WSOpData{Op: "sub", Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.sub(jdata, ch, callbackFun, reflect.TypeOf(response.SubOrdersResponse{}))
}

/**
【逐仓】取消订阅订单成交数据（unsub）
*/
func (wsNf *WSNotifyClient) IsolatedUnsubOrders(contractCode string, cid string) {
	if cid == "" {
		cid = config.DEFAULT_CID
	}

	ch := fmt.Sprintf("orders.%s", contractCode)
	opData := request.WSOpData{Op: "unsub", Cid: cid, Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.unsub(jdata, ch)
}

/**
【全仓】订阅订单成交数据（sub）
*/
func (wsNf *WSNotifyClient) CrossSubOrders(contractCode string, callbackFun OnSubOrdersResponse, cid string) {
	if cid == "" {
		cid = config.DEFAULT_CID
	}
	ch := fmt.Sprintf("orders_cross.%s", contractCode)
	opData := request.WSOpData{Op: "sub", Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.sub(jdata, ch, callbackFun, reflect.TypeOf(response.SubOrdersResponse{}))
}

/**
【全仓】取消订阅订单成交数据（unsub）
*/
func (wsNf *WSNotifyClient) CrossUnsubOrders(contractCode string, cid string) {
	if cid == "" {
		cid = config.DEFAULT_CID
	}

	ch := fmt.Sprintf("orders_cross.%s", contractCode)
	opData := request.WSOpData{Op: "unsub", Cid: cid, Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.unsub(jdata, ch)
}

// -------------------------------------------------------------

type OnSubAccountsResponse func(*response.SubAccountsResponse)

/**
【逐仓】资产变动数据（sub）
*/
func (wsNf *WSNotifyClient) IsolatedSubAcounts(contractCode string, callbackFun OnSubAccountsResponse, cid string) {
	if cid == "" {
		cid = config.DEFAULT_CID
	}

	ch := fmt.Sprintf("accounts.%s", contractCode)
	opData := request.WSOpData{Op: "sub", Cid: cid, Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.sub(jdata, ch, callbackFun, reflect.TypeOf(response.SubAccountsResponse{}))
}

/**
【逐仓】取消订阅资产变动数据（unsub）
*/
func (wsNf *WSNotifyClient) IsolatedUnsubAccounts(contractCode string, cid string) {
	if cid == "" {
		cid = config.DEFAULT_CID
	}

	ch := fmt.Sprintf("accounts.%s", contractCode)
	opData := request.WSOpData{Op: "unsub", Cid: cid, Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.unsub(jdata, ch)
}

/**
【全仓】资产变动数据（sub）
*/
func (wsNf *WSNotifyClient) CrossSubAcounts(contractCode string, callbackFun OnSubAccountsResponse, cid string) {
	if cid == "" {
		cid = config.DEFAULT_CID
	}

	ch := fmt.Sprintf("accounts_cross.%s", contractCode)
	opData := request.WSOpData{Op: "sub", Cid: cid, Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.sub(jdata, ch, callbackFun, reflect.TypeOf(response.SubAccountsResponse{}))
}

/**
【全仓】取消订阅资产变动数据（unsub）
*/
func (wsNf *WSNotifyClient) CrossUnsubAccounts(marginAccount string, cid string) {
	if cid == "" {
		cid = config.DEFAULT_CID
	}

	ch := fmt.Sprintf("accounts_cross.%s", marginAccount)
	opData := request.WSOpData{Op: "unsub", Cid: cid, Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.unsub(jdata, ch)
}

type OnSubPositionsResponse func(*response.SubPositionsResponse)

/**
【逐仓】持仓变动更新数据（sub）
*/
func (wsNf *WSNotifyClient) IsolatedSubPositions(contractCode string, callbackFun OnSubPositionsResponse, cid string) {
	if cid == "" {
		cid = config.DEFAULT_CID
	}

	ch := fmt.Sprintf("positions.%s", contractCode)
	opData := request.WSOpData{Op: "sub", Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.sub(jdata, ch, callbackFun, reflect.TypeOf(response.SubPositionsResponse{}))
}

/**
【逐仓】取消订阅持仓变动数据（unsub）
*/
func (wsNf *WSNotifyClient) IsolatdUnsubPositions(contractCode string, cid string) {
	if cid == "" {
		cid = config.DEFAULT_CID
	}

	ch := fmt.Sprintf("positions.%s", contractCode)
	opData := request.WSOpData{Op: "unsub", Cid: cid, Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.unsub(jdata, ch)
}

/**
【全仓】持仓变动更新数据（sub）
*/
func (wsNf *WSNotifyClient) CrossSubPositions(contractCode string, callbackFun OnSubPositionsResponse, cid string) {
	if cid == "" {
		cid = config.DEFAULT_CID
	}

	ch := fmt.Sprintf("positions_cross.%s", contractCode)
	opData := request.WSOpData{Op: "sub", Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.sub(jdata, ch, callbackFun, reflect.TypeOf(response.SubPositionsResponse{}))
}

/**
【全仓】取消订阅持仓变动数据（unsub）
*/
func (wsNf *WSNotifyClient) CrossUnsubPositions(contractCode string, cid string) {
	if cid == "" {
		cid = config.DEFAULT_CID
	}

	ch := fmt.Sprintf("positions_cross.%s", contractCode)
	opData := request.WSOpData{Op: "unsub", Cid: cid, Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.unsub(jdata, ch)
}

type OnSubMatchOrdersResponse func(*response.SubOrdersResponse)

/*
【逐仓】订阅合约订单撮合数据（sub）
*/
func (wsNf *WSNotifyClient) IsolatedSubMatchOrders(contractCode string, callbackFun OnSubMatchOrdersResponse, cid string) {
	if cid == "" {
		cid = config.DEFAULT_CID
	}

	contractCode = strings.ToLower(contractCode)
	ch := fmt.Sprintf("matchOrders.%s", contractCode)
	opData := request.WSOpData{Op: "sub", Cid: cid, Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.sub(jdata, ch, callbackFun, reflect.TypeOf(response.SubOrdersResponse{}))
}

/**
【逐仓】取消订阅合约订单撮合数据（unsub）
*/
func (wsNf *WSNotifyClient) IsolatedUnsubMathOrders(contractCode string, cid string) {
	if cid == "" {
		cid = config.DEFAULT_CID
	}

	contractCode = strings.ToLower(contractCode)
	ch := fmt.Sprintf("matchOrders.%s", contractCode)
	opData := request.WSOpData{Op: "unsub", Cid: cid, Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.unsub(jdata, ch)
}

/**
【全仓】订阅合约订单撮合数据（sub）
*/
func (wsNf *WSNotifyClient) CrossSubMatchOrders(contractCode string, callbackFun OnSubMatchOrdersResponse, cid string) {
	if cid == "" {
		cid = config.DEFAULT_CID
	}

	contractCode = strings.ToLower(contractCode)
	ch := fmt.Sprintf("matchOrders_cross.%s", contractCode)
	opData := request.WSOpData{Op: "sub", Cid: cid, Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.sub(jdata, ch, callbackFun, reflect.TypeOf(response.SubOrdersResponse{}))
}

/**
【全仓】取消订阅合约订单撮合数据（unsub）
*/
func (wsNf *WSNotifyClient) CrossUnsubMathOrders(contractCode string, cid string) {
	if cid == "" {
		cid = config.DEFAULT_CID
	}

	contractCode = strings.ToLower(contractCode)
	ch := fmt.Sprintf("matchOrders_cross.%s", contractCode)
	opData := request.WSOpData{Op: "unsub", Cid: cid, Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.unsub(jdata, ch)
}

type OnSubLiquidationOrdersResponse func(*response.SubLiquidationOrdersResponse)

/**
【通用】订阅强平订单数据(免鉴权)（sub）
*/
func (wsNf *WSNotifyClient) SubLiquidationOrders(contractCode string, callbackFun OnSubLiquidationOrdersResponse, cid string) {
	if cid == "" {
		cid = config.DEFAULT_CID
	}

	ch := fmt.Sprintf("public.%s.liquidation_orders", contractCode)
	opData := request.WSOpData{Op: "sub", Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.sub(jdata, ch, callbackFun, reflect.TypeOf(response.SubLiquidationOrdersResponse{}))
}

/**
【通用】取消订阅强平订单(免鉴权)（unsub）
*/
func (wsNf *WSNotifyClient) UnsubLiquidationOrders(contractCode string, cid string) {
	if cid == "" {
		cid = config.DEFAULT_CID
	}

	ch := fmt.Sprintf("public.%s.liquidation_orders", contractCode)
	opData := request.WSOpData{Op: "unsub", Cid: cid, Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.unsub(jdata, ch)
}

type OnSubFundingRateResponse func(*response.SubFundingRateResponse)

/**
【通用】订阅资金费率推送(免鉴权)（sub）
*/
func (wsNf *WSNotifyClient) SubFundingRate(contractCode string, callbackFun OnSubFundingRateResponse, cid string) {
	if cid == "" {
		cid = config.DEFAULT_CID
	}

	ch := fmt.Sprintf("public.%s.funding_rate", contractCode)
	opData := request.WSOpData{Op: "sub", Cid: cid, Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.sub(jdata, ch, callbackFun, reflect.TypeOf(response.SubFundingRateResponse{}))
}

/**
【通用】取消订阅资金费率(免鉴权)（unsub）
*/
func (wsNf *WSNotifyClient) UnsubFundingRate(contractCode string, cid string) {
	if cid == "" {
		cid = config.DEFAULT_CID
	}

	ch := fmt.Sprintf("public.%s.funding_rate", contractCode)
	opData := request.WSOpData{Op: "unsub", Cid: cid, Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.unsub(jdata, ch)
}

type OnSubContractInfoResponse func(*response.SubContractInfoResponse)

/**
【通用】订阅合约信息变动(免鉴权)（sub）
*/
func (wsNf *WSNotifyClient) SubContractInfo(contractCode string, callbackFun OnSubContractInfoResponse, cid string) {
	if cid == "" {
		cid = config.DEFAULT_CID
	}

	ch := fmt.Sprintf("public.%s.contract_info", contractCode)
	opData := request.WSOpData{Op: "sub", Cid: cid, Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.sub(jdata, ch, callbackFun, reflect.TypeOf(response.SubContractInfoResponse{}))
}

/**
【通用】取消订阅合约信息变动(免鉴权)（unsub）
*/
func (wsNf *WSNotifyClient) UnsubContractInfo(contractCode string, cid string) {
	if cid == "" {
		cid = config.DEFAULT_CID
	}

	ch := fmt.Sprintf("public.%s.contract_info", contractCode)
	opData := request.WSOpData{Op: "unsub", Cid: cid, Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.unsub(jdata, ch)
}

type OnSubTriggerOrderResponse func(*response.SubTriggerOrderResponse)

/**
【逐仓】订阅计划委托订单更新(sub)
*/
func (wsNf *WSNotifyClient) IsolatedSubTriggerOrder(contractCode string, callbackFun OnSubTriggerOrderResponse, cid string) {
	if cid == "" {
		cid = config.DEFAULT_CID
	}

	ch := fmt.Sprintf("trigger_order.%s", contractCode)
	opData := request.WSOpData{Op: "sub", Cid: cid, Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.sub(jdata, ch, callbackFun, reflect.TypeOf(response.SubTriggerOrderResponse{}))
}

/**
【逐仓】取消订阅计划委托订单更新（unsub）
*/
func (wsNf *WSNotifyClient) IsolatedUnsubTriggerOrder(contractCode string, cid string) {
	if cid == "" {
		cid = config.DEFAULT_CID
	}

	ch := fmt.Sprintf("trigger_order.%s", contractCode)
	opData := request.WSOpData{Op: "unsub", Cid: cid, Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.unsub(jdata, ch)
}

/**
【全仓】订阅计划委托订单更新(sub)
*/
func (wsNf *WSNotifyClient) CrossSubTriggerOrder(contractCode string, callbackFun OnSubTriggerOrderResponse, cid string) {
	if cid == "" {
		cid = config.DEFAULT_CID
	}

	ch := fmt.Sprintf("trigger_order_cross.%s", contractCode)
	opData := request.WSOpData{Op: "sub", Cid: cid, Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.sub(jdata, ch, callbackFun, reflect.TypeOf(response.SubTriggerOrderResponse{}))
}

/**
【全仓】取消订阅计划委托订单更新（unsub）
*/
func (wsNf *WSNotifyClient) CrossUnsubTriggerOrder(contractCode string, cid string) {
	if cid == "" {
		cid = config.DEFAULT_CID
	}

	ch := fmt.Sprintf("trigger_order_cross.%s", contractCode)
	opData := request.WSOpData{Op: "unsub", Cid: cid, Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.unsub(jdata, ch)
}
