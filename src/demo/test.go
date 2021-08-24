package main

import (
	"fmt"
	"github.com/xiaomingping/huobi_contract_golang/src/ws/api"
	"github.com/xiaomingping/huobi_contract_golang/src/ws/response"
	"time"
)

func main() {
	callback := make(chan struct{}, 1)
	client := new(api.WSNotifyClient).Init("", "", "", callback)
	client.CrossSubOrders("DOGE-USDT", Resp, "")
	select {
	case <-callback:
		fmt.Println("已关闭")
		time.Sleep(time.Second * 1)
	}

}

func Resp(response *response.SubOrdersResponse) {
	fmt.Printf("%+v\n", response)
}
