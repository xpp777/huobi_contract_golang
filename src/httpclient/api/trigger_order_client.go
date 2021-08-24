package api

import (
	"encoding/json"
	"fmt"
	"github.com/xiaomingping/huobi_contract_golang/src/config"
	"github.com/xiaomingping/huobi_contract_golang/src/httpclient/request"
	"github.com/xiaomingping/huobi_contract_golang/src/httpclient/response"
	"github.com/xiaomingping/huobi_contract_golang/src/logh"
	"github.com/xiaomingping/huobi_contract_golang/src/utils/reqbuilder"
)

type TriggerOrderClient struct {
	PUrlBuilder *reqbuilder.PrivateUrlBuilder
}

func (toc *TriggerOrderClient) Init(accessKey string, secretKey string, host string) *TriggerOrderClient {
	if host == "" {
		host = config.LINEAR_SWAP_DEFAULT_HOST
	}
	toc.PUrlBuilder = new(reqbuilder.PrivateUrlBuilder).Init(accessKey, secretKey, host)
	return toc
}

func (toc *TriggerOrderClient) IsolatedPlaceOrderAsync(data chan response.PlaceOrderResponse, request request.TriggerOrderPlaceOrderRequest) {
	url := toc.PUrlBuilder.Build(config.POST_METHOD, "/linear-swap-api/v1/swap_trigger_order", nil)

	content, err := json.Marshal(request)
	if err != nil {
		logh.Error("PlaceOrderRequest to json error: %v", err)
	}

	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.PlaceOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to PlaceOrderResponse error: %s", getErr)
	}
	data <- result
}

func (toc *TriggerOrderClient) CrossPlaceOrderAsync(data chan response.PlaceOrderResponse, request request.TriggerOrderPlaceOrderRequest) {
	url := toc.PUrlBuilder.Build(config.POST_METHOD, "/linear-swap-api/v1/swap_cross_trigger_order", nil)

	content, err := json.Marshal(request)
	if err != nil {
		logh.Error("PlaceOrderRequest to json error: %v", err)
	}

	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.PlaceOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to PlaceOrderResponse error: %s", getErr)
	}
	data <- result
}

func (toc *TriggerOrderClient) IsolatedCancelOrderAsync(data chan response.CancelOrderResponse,
	contractCode string, orderId string, offset string, direction string) {
	// url
	url := toc.PUrlBuilder.Build(config.POST_METHOD, "/linear-swap-api/v1/swap_trigger_cancel", nil)
	if orderId == "" {
		url = toc.PUrlBuilder.Build(config.POST_METHOD, "/linear-swap-api/v1/swap_trigger_cancelall", nil)
	}

	// content
	content := fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	if orderId != "" {
		content += fmt.Sprintf(",\"order_id\": \"%s\"", orderId)
	}
	if offset != "" {
		content += fmt.Sprintf(",\"offset\": \"%s\"", offset)
	}
	if direction != "" {
		content += fmt.Sprintf(",\"direction\": \"%s\"", direction)
	}
	if content != "" {
		content = fmt.Sprintf("{ %s }", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.CancelOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to CancelOrderResponse error: %s", getErr)
	}
	data <- result
}

func (toc *TriggerOrderClient) CrossCancelOrderAsync(data chan response.CancelOrderResponse,
	contractCode string, orderId string, offset string, direction string) {
	// url
	url := toc.PUrlBuilder.Build(config.POST_METHOD, "/linear-swap-api/v1/swap_cross_trigger_cancel", nil)
	if orderId == "" {
		url = toc.PUrlBuilder.Build(config.POST_METHOD, "/linear-swap-api/v1/swap_cross_trigger_cancelall", nil)
	}

	// content
	content := fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	if orderId != "" {
		content += fmt.Sprintf(",\"order_id\": \"%s\"", orderId)
	}
	if offset != "" {
		content += fmt.Sprintf(",\"offset\": \"%s\"", offset)
	}
	if direction != "" {
		content += fmt.Sprintf(",\"direction\": \"%s\"", direction)
	}
	if content != "" {
		content = fmt.Sprintf("{ %s }", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.CancelOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to CancelOrderResponse error: %s", getErr)
	}
	data <- result
}

func (toc *TriggerOrderClient) IsolatedGetOpenOrderAsync(data chan response.GetOpenOrderResponse,
	contractCode string, pageIndex int, pageSize int, tradeType int) {
	// url
	url := toc.PUrlBuilder.Build(config.POST_METHOD, "/linear-swap-api/v1/swap_trigger_openorders", nil)

	// content
	content := fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	if pageIndex != 0 {
		content += fmt.Sprintf(",\"page_index\": %d", pageIndex)
	}
	if pageSize != 0 {
		content += fmt.Sprintf(",\"page_size\": %d", pageSize)
	}
	if tradeType != 0 {
		content += fmt.Sprintf(",\"trade_type\": %d", tradeType)
	}
	if content != "" {
		content = fmt.Sprintf("{ %s }", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetOpenOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetOpenOrderResponse error: %s", getErr)
	}
	data <- result
}

func (toc *TriggerOrderClient) CrossGetOpenOrderAsync(data chan response.GetOpenOrderResponse,
	contractCode string, pageIndex int, pageSize int, tradeType int) {
	// url
	url := toc.PUrlBuilder.Build(config.POST_METHOD, "/linear-swap-api/v1/swap_cross_trigger_openorders", nil)

	// content
	content := fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	if pageIndex != 0 {
		content += fmt.Sprintf(",\"page_index\": %d", pageIndex)
	}
	if pageSize != 0 {
		content += fmt.Sprintf(",\"page_size\": %d", pageSize)
	}
	if tradeType != 0 {
		content += fmt.Sprintf(",\"trade_type\": %d", tradeType)
	}
	if content != "" {
		content = fmt.Sprintf("{ %s }", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetOpenOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetOpenOrderResponse error: %s", getErr)
	}
	data <- result
}

func (toc *TriggerOrderClient) IsolatedGetHisOrderAsync(data chan response.GetHisOrderResponse, contractCode string, tradeType int, status string, createDate int,
	pageIndex int, pageSize int, sortBy string) {
	// url
	url := toc.PUrlBuilder.Build(config.POST_METHOD, "/linear-swap-api/v1/swap_trigger_hisorders", nil)

	// content
	content := fmt.Sprintf(",\"contract_code\": \"%s\",\"trade_type\": %d,\"status\": \"%s\",\"create_date\": %d", contractCode, tradeType, status, createDate)
	if pageIndex != 0 {
		content += fmt.Sprintf(",\"page_index\": %d", pageIndex)
	}
	if pageSize != 0 {
		content += fmt.Sprintf(",\"page_size\": %d", pageSize)
	}
	if sortBy != "" {
		content += fmt.Sprintf(",\"sort_by\": \"%s\"", sortBy)
	}
	if content != "" {
		content = fmt.Sprintf("{ %s }", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetHisOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetHisOrderResponse error: %s", getErr)
	}
	data <- result
}

func (toc *TriggerOrderClient) CrossGetHisOrderAsync(data chan response.GetHisOrderResponse, contractCode string, tradeType int, status string, createDate int,
	pageIndex int, pageSize int, sortBy string) {
	// url
	url := toc.PUrlBuilder.Build(config.POST_METHOD, "/linear-swap-api/v1/swap_cross_trigger_hisorders", nil)

	// content
	content := fmt.Sprintf(",\"contract_code\": \"%s\",\"trade_type\": %d,\"status\": \"%s\",\"create_date\": %d", contractCode, tradeType, status, createDate)
	if pageIndex != 0 {
		content += fmt.Sprintf(",\"page_index\": %d", pageIndex)
	}
	if pageSize != 0 {
		content += fmt.Sprintf(",\"page_size\": %d", pageSize)
	}
	if sortBy != "" {
		content += fmt.Sprintf(",\"sort_by\": \"%s\"", sortBy)
	}
	if content != "" {
		content = fmt.Sprintf("{ %s }", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetHisOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetHisOrderResponse error: %s", getErr)
	}
	data <- result
}

func (oc *OrderClient) IsolatedTpslOrderAsync(data chan response.TpslOrderResponse, request request.TriggerOrderTpslOrderRequest) {
	url := oc.PUrlBuilder.Build(config.POST_METHOD, "/linear-swap-api/v1/swap_tpsl_order", nil)

	content, err := json.Marshal(request)
	if err != nil {
		logh.Error("PlaceOrderRequest to json error: %v", err)
	}
	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.TpslOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to TpslOrderResponse error: %s", getErr)
	}
	data <- result
}

func (oc *OrderClient) CrossTpslOrderAsync(data chan response.TpslOrderResponse, request request.TriggerOrderTpslOrderRequest) {
	url := oc.PUrlBuilder.Build(config.POST_METHOD, "/linear-swap-api/v1/swap_cross_tpsl_order", nil)

	content, err := json.Marshal(request)
	if err != nil {
		logh.Error("PlaceOrderRequest to json error: %v", err)
	}
	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.TpslOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to TpslOrderResponse error: %s", getErr)
	}
	data <- result
}

func (toc *TriggerOrderClient) IsolatedTpslCancelAsync(data chan response.CancelOrderResponse,
	contractCode string, orderId string, direction string) {
	// url
	url := toc.PUrlBuilder.Build(config.POST_METHOD, "/linear-swap-api/v1/swap_tpsl_cancel", nil)
	if orderId == "" {
		url = toc.PUrlBuilder.Build(config.POST_METHOD, "/linear-swap-api/v1/swap_tpsl_cancelall", nil)
	}

	// content
	content := fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	if orderId != "" {
		content += fmt.Sprintf(",\"order_id\": \"%s\"", orderId)
	}
	if direction != "" {
		content += fmt.Sprintf(",\"direction\": \"%s\"", direction)
	}
	if content != "" {
		content = fmt.Sprintf("{ %s }", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.CancelOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to CancelOrderResponse error: %s", getErr)
	}
	data <- result
}

func (toc *TriggerOrderClient) CrossTpslCancelAsync(data chan response.CancelOrderResponse,
	contractCode string, orderId string, direction string) {
	// url
	url := toc.PUrlBuilder.Build(config.POST_METHOD, "/linear-swap-api/v1/swap_cross_tpsl_cancel", nil)
	if orderId == "" {
		url = toc.PUrlBuilder.Build(config.POST_METHOD, "/linear-swap-api/v1/swap_cross_tpsl_cancelall", nil)
	}

	// content
	content := fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	if orderId != "" {
		content += fmt.Sprintf(",\"order_id\": \"%s\"", orderId)
	}
	if direction != "" {
		content += fmt.Sprintf(",\"direction\": \"%s\"", direction)
	}
	if content != "" {
		content = fmt.Sprintf("{ %s }", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.CancelOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to CancelOrderResponse error: %s", getErr)
	}
	data <- result
}

func (toc *TriggerOrderClient) IsolatedGetTpslOpenOrderAsync(data chan response.GetOpenOrderResponse,
	contractCode string, pageIndex int, pageSize int, tradeType int) {
	// url
	url := toc.PUrlBuilder.Build(config.POST_METHOD, "/linear-swap-api/v1/swap_tpsl_openorders", nil)

	// content
	content := fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	if pageIndex != 0 {
		content += fmt.Sprintf(",\"page_index\": %d", pageIndex)
	}
	if pageSize != 0 {
		content += fmt.Sprintf(",\"page_size\": %d", pageSize)
	}
	if tradeType != 0 {
		content += fmt.Sprintf(",\"trade_type\": %d", tradeType)
	}
	if content != "" {
		content = fmt.Sprintf("{ %s }", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetOpenOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetOpenOrderResponse error: %s", getErr)
	}
	data <- result
}

func (toc *TriggerOrderClient) CrossGetTpslOpenOrderAsync(data chan response.GetOpenOrderResponse,
	contractCode string, pageIndex int, pageSize int, tradeType int) {
	// url
	url := toc.PUrlBuilder.Build(config.POST_METHOD, "/linear-swap-api/v1/swap_cross_tpsl_openorders", nil)

	// content
	content := fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	if pageIndex != 0 {
		content += fmt.Sprintf(",\"page_index\": %d", pageIndex)
	}
	if pageSize != 0 {
		content += fmt.Sprintf(",\"page_size\": %d", pageSize)
	}
	if tradeType != 0 {
		content += fmt.Sprintf(",\"trade_type\": %d", tradeType)
	}
	if content != "" {
		content = fmt.Sprintf("{ %s }", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetOpenOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetOpenOrderResponse error: %s", getErr)
	}
	data <- result
}

func (toc *TriggerOrderClient) IsolatedGetTpslHisOrderAsync(data chan response.GetHisOrderResponse,
	contractCode string, status string, createDate int, pageIndex int, pageSize int, sortBy string) {
	// url
	url := toc.PUrlBuilder.Build(config.POST_METHOD, "/linear-swap-api/v1/swap_tpsl_hisorders", nil)

	// content
	content := fmt.Sprintf(",\"contract_code\":\"%s\", \"status\":\"%s\", \"create_date\":%d", contractCode, status, createDate)
	if pageIndex != 0 {
		content += fmt.Sprintf(",\"page_index\": %d", pageIndex)
	}
	if pageSize != 0 {
		content += fmt.Sprintf(",\"page_size\": %d", pageSize)
	}
	if sortBy != "" {
		content += fmt.Sprintf(",\"sort_by\":\"%s\"", sortBy)
	}
	if content != "" {
		content = fmt.Sprintf("{ %s }", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetHisOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetHisOrderResponse error: %s", getErr)
	}
	data <- result
}

func (toc *TriggerOrderClient) CrossGetTpslHisOrderAsync(data chan response.GetHisOrderResponse,
	contractCode string, status string, createDate int, pageIndex int, pageSize int, sortBy string) {
	// url
	url := toc.PUrlBuilder.Build(config.POST_METHOD, "/linear-swap-api/v1/swap_cross_tpsl_hisorders", nil)

	// content
	content := fmt.Sprintf(",\"contract_code\":\"%s\", \"status\":\"%s\", \"create_date\":%d", contractCode, status, createDate)
	if pageIndex != 0 {
		content += fmt.Sprintf(",\"page_index\": %d", pageIndex)
	}
	if pageSize != 0 {
		content += fmt.Sprintf(",\"page_size\": %d", pageSize)
	}
	if sortBy != "" {
		content += fmt.Sprintf(",\"sort_by\":\"%s\"", sortBy)
	}
	if content != "" {
		content = fmt.Sprintf("{ %s }", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetHisOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetHisOrderResponse error: %s", getErr)
	}
	data <- result
}

func (toc *TriggerOrderClient) IsolatedGetRelationTpslOrderAsync(data chan response.GetRelationTpslOrderResponse,
	contractCode string, orderId int) {
	// url
	url := toc.PUrlBuilder.Build(config.POST_METHOD, "/linear-swap-api/v1/swap_relation_tpsl_order", nil)

	// content
	content := fmt.Sprintf(",\"contract_code\":\"%s\", \"order_id\":%d", contractCode, orderId)
	if content != "" {
		content = fmt.Sprintf("{ %s }", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetRelationTpslOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetRelationTpslOrderResponse error: %s", getErr)
	}
	data <- result
}

func (toc *TriggerOrderClient) CrossGetRelationTpslOrderAsync(data chan response.GetRelationTpslOrderResponse,
	contractCode string, orderId int) {
	// url
	url := toc.PUrlBuilder.Build(config.POST_METHOD, "/linear-swap-api/v1/swap_cross_relation_tpsl_order", nil)

	// content
	content := fmt.Sprintf(",\"contract_code\":\"%s\", \"order_id\":%d", contractCode, orderId)
	if content != "" {
		content = fmt.Sprintf("{ %s }", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetRelationTpslOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetRelationTpslOrderResponse error: %s", getErr)
	}
	data <- result
}

