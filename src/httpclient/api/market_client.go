package api

import (
	"encoding/json"
	"fmt"
	"github.com/xiaomingping/huobi_contract_golang/src/config"
	"github.com/xiaomingping/huobi_contract_golang/src/httpclient/response"
	"github.com/xiaomingping/huobi_contract_golang/src/logh"
	"github.com/xiaomingping/huobi_contract_golang/src/utils/reqbuilder"
)

type MarketClient struct {
	PUrlBuilder *reqbuilder.PublicUrlBuilder
}

func (mc *MarketClient) Init(host string) *MarketClient {
	if host == "" {
		host = config.LINEAR_SWAP_DEFAULT_HOST
	}
	mc.PUrlBuilder = new(reqbuilder.PublicUrlBuilder).Init(host)
	return mc
}

func (mc *MarketClient) GetContractInfoAsync(data chan response.GetContractInfoResponse, contractCode string) {
	// location
	location := "/linear-swap-api/v1/swap_contract_info"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("contract_code=%s", contractCode)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetContractInfoResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetContractInfoAsync error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) GetIndexAsync(data chan response.GetIndexResponse, contractCode string) {
	// location
	location := "/linear-swap-api/v1/swap_index"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("contract_code=%s", contractCode)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetIndexResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetIndexResponse error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) GetPriceLimitAsync(data chan response.GetPriceLimitResponse, contractCode string) {
	// location
	location := "/linear-swap-api/v1/swap_price_limit"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("contract_code=%s", contractCode)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetPriceLimitResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetPriceLimitResponse error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) GetOpenInterestAsync(data chan response.GetOpenInterestResponse, contractCode string) {
	// location
	location := "/linear-swap-api/v1/swap_open_interest"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("contract_code=%s", contractCode)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetOpenInterestResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetOpenInterestResponse error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) GetDepthAsync(data chan response.GetDepthResponse, contractCode string, fcType string) {
	// location
	location := fmt.Sprintf("/linear-swap-ex/market/depth?contract_code=%s&type=%s", contractCode, fcType)

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetDepthResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetDepthResponse error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) GetBboAsync(data chan response.GetBboResponse, contractCode string) {
	// location
	location := "/linear-swap-ex/market/bbo"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("contract_code=%s", contractCode)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetBboResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetBboResponse error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) GetKLineAsync(data chan response.GetKLineResponse, contractCode string, period string, size int, from int, to int) {
	// location
	location := fmt.Sprintf("/linear-swap-ex/market/history/kline?contract_code=%s&period=%s", contractCode, period)

	// option
	option := ""
	if size != 0 {
		option += fmt.Sprintf("&size=%d", size)
	}
	if from != 0 {
		option += fmt.Sprintf("&from=%d", from)
	}
	if to != 0 {
		option += fmt.Sprintf("&to=%d", to)
	}
	if option != "" {
		location += option
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetKLineResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetKLineResponse error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) GetMarkPriceKLineAsync(data chan response.GetStrKLineResponse, contractCode string, period string, size int) {
	// location
	location := fmt.Sprintf("/index/market/history/linear_swap_mark_price_kline?contract_code=%s&period=%s&size=%d", contractCode, period, size)

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetStrKLineResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetStrKLineResponse error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) GetMergedAsync(data chan response.GetMergedResponse, contractCode string) {
	// location
	location := fmt.Sprintf("/linear-swap-ex/market/detail/merged?contract_code=%s", contractCode)

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetMergedResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetMergedResponse error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) GetBatchMergedAsync(data chan response.GetBatchMergedResponse, contractCode string) {
	// location
	location := fmt.Sprintf("/linear-swap-ex/market/detail/batch_merged")

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("&contract_code=%s", contractCode)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetBatchMergedResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetBatchMergedAsync error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) GetTradeAsync(data chan response.GetTradeResponse, contractCode string) {
	// location
	location := "/linear-swap-ex/market/trade"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("&contract_code=%s", contractCode)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetTradeResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetTradeResponse error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) GetHisTradeAsync(data chan response.GetHisTradeResponse, contractCode string, size int) {
	// location
	location := fmt.Sprintf("/linear-swap-ex/market/history/trade?contract_code=%s&size=%d", contractCode, size)

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetHisTradeResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetHisTradeResponse error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) GetRiskInfoAsync(data chan response.GetRiskInfoResponse, contractCode string) {
	// location
	location := "/linear-swap-api/v1/swap_risk_info"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("contract_code=%s", contractCode)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetRiskInfoResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetRiskInfoResponse error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) GetInsuranceFundAsync(data chan response.GetInsuranceFundResponse, contractCode string, pageIndex int, pageSize int) {
	// location
	location := fmt.Sprintf("/linear-swap-api/v1/swap_insurance_fund?contract_code=%s", contractCode)

	// option
	option := ""
	if pageIndex != 0 {
		option += fmt.Sprintf("&size=%d", pageIndex)
	}
	if pageSize != 0 {
		option += fmt.Sprintf("&from=%d", pageSize)
	}
	if option != "" {
		location += option
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetInsuranceFundResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetInsuranceFundResponse error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) IsolatedGetAdjustFactorFundAsync(data chan response.GetAdjustFactorFundResponse, contractCode string) {
	// location
	location := "/linear-swap-api/v1/swap_adjustfactor"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("?contract_code=%s", contractCode)
	}
	if option != "" {
		location += option
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetAdjustFactorFundResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetAdjustFactorFundResponse error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) CrossGetAdjustFactorFundAsync(data chan response.GetAdjustFactorFundResponse, contractCode string) {
	// location
	location := "/linear-swap-api/v1/swap_cross_adjustfactor"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("?contract_code=%s", contractCode)
	}
	if option != "" {
		location += option
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetAdjustFactorFundResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetAdjustFactorFundResponse error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) GetHisOpenInterestAsync(data chan response.GetHisOpenInterestResponse, contractCode string, period string, amountType int, size int) {
	// location
	location := fmt.Sprintf("/linear-swap-api/v1/swap_his_open_interest?contract_code=%s&period=%s&amount_type=%d", contractCode, period, amountType)

	// option
	option := ""
	if size != 0 {
		option += fmt.Sprintf("&size=%d", size)
	}
	if option != "" {
		location += option
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetHisOpenInterestResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetHisOpenInterestResponse error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) IsolatedGetLadderMarginAsync(data chan response.GetLadderMarginResponse, contractCode string) {
	// location
	location := "/linear-swap-api/v1/swap_ladder_margin"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("?contract_code=%s", contractCode)
	}
	if option != "" {
		location += option
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetLadderMarginResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetLadderMarginResponse error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) CrossGetLadderMarginAsync(data chan response.GetLadderMarginResponse, contractCode string) {
	// location
	location := "/linear-swap-api/v1/swap_cross_ladder_margin"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("?contract_code=%s", contractCode)
	}
	if option != "" {
		location += option
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetLadderMarginResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetLadderMarginResponse error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) GetEliteAccountRatioAsync(data chan response.GetEliteRatioResponse, contractCode string, period string) {
	// location
	location := fmt.Sprintf("/linear-swap-api/v1/swap_elite_account_ratio?contract_code=%s&period=%s", contractCode, period)

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetEliteRatioResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetElitePositionRatioResponse error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) GetElitePositionRatioAsync(data chan response.GetEliteRatioResponse, contractCode string, period string) {
	// location
	location := fmt.Sprintf("/linear-swap-api/v1/swap_elite_position_ratio?contract_code=%s&period=%s", contractCode, period)

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetEliteRatioResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetElitePositionRatioResponse error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) IsolatedGetApiStateAsync(data chan response.GetApiStateResponse, contractCode string) {
	// location
	location := "/linear-swap-api/v1/swap_api_state"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("?contract_code=%s", contractCode)
	}
	if option != "" {
		location += option
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetApiStateResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetApiStateResponse error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) CrossGetTransferStateAsync(data chan response.GetTransferStateResponse, marginAccount string) {
	// location
	location := "/linear-swap-api/v1/swap_cross_transfer_state"

	// option
	if marginAccount == "" {
		marginAccount = "USDT"
	}
	location += fmt.Sprintf("?margin_account=%s", marginAccount)

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetTransferStateResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetTransferStateResponse error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) CrossGetTradeStateAsync(data chan response.GetTradeStateResponse, contractCode string) {
	// location
	location := "/linear-swap-api/v1/swap_cross_trade_state"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("?contract_code=%s", contractCode)
	}
	if option != "" {
		location += option
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetTradeStateResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetTradeStatusResponse error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) GetFundingRateAsync(data chan response.GetFundingRateResponse, contractCode string) {
	// location
	location := fmt.Sprintf("/linear-swap-api/v1/swap_funding_rate?contract_code=%s", contractCode)

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetFundingRateResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetFundingRateResponse error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) GetBatchFundingRateAsync(data chan response.GetBatchFundingRateResponse, contractCode string) {
	// location
	location := "/linear-swap-api/v1/swap_batch_funding_rate"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("?contract_code=%s", contractCode)
	}
	if option != "" {
		location += option
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetBatchFundingRateResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetBatchFundingRateResponse error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) GetHisFundingRateAsync(data chan response.GetHisFundingRateResponse, contractCode string, pageIndex int, pageSize int) {
	// location
	location := fmt.Sprintf("/linear-swap-api/v1/swap_historical_funding_rate?contract_code=%s", contractCode)

	// option
	option := ""
	if pageIndex != 0 {
		option += fmt.Sprintf("&page_index=%d", pageIndex)
	}
	if pageSize != 0 {
		option += fmt.Sprintf("&page_size=%d", pageSize)
	}
	if option != "" {
		location += option
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetHisFundingRateResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetHisFundingRateResponse error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) GetLiquidationOrdersAsync(data chan response.GetLiquidationOrdersResponse, contractCode string, tradeType int, createDate int,
	pageIndex int, pageSize int) {
	// location
	location := fmt.Sprintf("/linear-swap-api/v1/swap_liquidation_orders?contract_code=%s&trade_type=%d&create_date=%d", contractCode, tradeType, createDate)

	// option
	option := ""
	if pageIndex != 0 {
		option += fmt.Sprintf("&page_index=%d", pageIndex)
	}
	if pageSize != 0 {
		option += fmt.Sprintf("&page_size=%d", pageSize)
	}
	if option != "" {
		location += option
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetLiquidationOrdersResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetLiquidationOrdersResponse error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) GetPremiumIndexKLineAsync(data chan response.GetStrKLineResponse, contractCode string, period string, size int) {
	// location
	location := fmt.Sprintf("/index/market/history/linear_swap_premium_index_kline?contract_code=%s&period=%s&size=%d", contractCode, period, size)

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetStrKLineResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetPriceLimitResponse error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) GetEstimatedRateKLineAsync(data chan response.GetStrKLineResponse, contractCode string, period string, size int) {
	// location
	location := fmt.Sprintf("/index/market/history/linear_swap_estimated_rate_kline?contract_code=%s&period=%s&size=%d", contractCode, period, size)

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetStrKLineResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetStrKLineResponse error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) GetBasisAsync(data chan response.GetBasisResponse, contractCode string, period string, size int, basisPriceType string) {
	// location
	location := fmt.Sprintf("/index/market/history/linear_swap_basis?contract_code=%s&period=%s&size=%d", contractCode, period, size)

	// option
	option := ""
	if basisPriceType != "" {
		option += fmt.Sprintf("&basis_price_type=%s", basisPriceType)
	}
	if option != "" {
		location += option
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetBasisResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetBasisResponse error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) GetEstimatedSettlementPriceAsync(data chan response.GetEstimatedSettlementPriceResponse, contractCode string) {
	// location
	location := "/linear-swap-api/v1/swap_estimated_settlement_price"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("?contract_code=%s", contractCode)
	}
	if option != "" {
		location += option
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.GetEstimatedSettlementPriceResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to GetEstimatedSettlementPriceResponse error: %s", getErr)
	}
	data <- result
}
