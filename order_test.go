package dada

import (
	"ixingkong.com/qiyi-tech/dada/schema"
	"testing"
)

func TestAddOrder(t *testing.T) {
	resp, err := client.AddOrder(&schema.AddOrderRequest{
		ShopNo:          "9727fc8cd1cf4877",
		OriginId:        "orderNo-324923677-1702884024074",
		CityCode:        "021",
		CargoPrice:      100,
		IsPrepay:        0,
		ReceiverName:    "李四",
		ReceiverAddress: "东方渔人码头",
		ReceiverLat:     31.257801,
		ReceiverLng:     121.538842,
		Callback:        "www.example.com/example",
		CargoWeight:     1,
		ReceiverPhone:   "13011112222",
	})
	if err == nil {
		t.Log(resp)
	}
}
