package api

import (
	"encoding/json"
	"fmt"
	"github.com/xiaomingping/huobi_contract_golang/src/config"
	"github.com/xiaomingping/huobi_contract_golang/src/httpclient/response"
	"github.com/xiaomingping/huobi_contract_golang/src/logh"
	"github.com/xiaomingping/huobi_contract_golang/src/utils/reqbuilder"
)

type TransferClient struct {
	PUrlBuilder *reqbuilder.PrivateUrlBuilder
}

func (tc *TransferClient) Init(accessKey string, secretKey string, host string) *TransferClient {
	if host == "" {
		host = config.SPOT_DEFAULT_HOST
	}
	tc.PUrlBuilder = new(reqbuilder.PrivateUrlBuilder).Init(accessKey, secretKey, host)
	return tc
}

func (tc *TransferClient) TransferAsync(data chan response.TransferResponse, from string, to string, amount float32, marginAccount string, currency string) {
	if currency == "" {
		currency = "USDT"
	}

	// ulr
	url := tc.PUrlBuilder.Build(config.POST_METHOD, "/v2/account/transfer", nil)

	// content
	content := fmt.Sprintf("{ \"from\":\"%s\", \"to\":\"%s\", \"currency\":\"%s\", \"amount\":%f, \"margin-account\":\"%s\" }", from, to, currency, amount, marginAccount)

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		logh.Error("http get error: %s", getErr)
	}
	result := response.TransferResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		logh.Error("convert json to TransferResponse error: %s", getErr)
	}
	data <- result
}


