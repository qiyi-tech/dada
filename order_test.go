package dada

import (
	"fmt"
	"ixingkong.com/qiyi-tech/dada/schema"
	"testing"
)

func TestAddOrder(t *testing.T) {
	_, err := client.AddOrder(&schema.AddOrderRequest{
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
	if err.Error() != "2105:订单已下发,如要重发,请使用重发接口" {
		t.Errorf("err:%v", err)
	}
}

func TestReAddOrder(t *testing.T) {
	_, err := client.ReAddOrder(&schema.ReAddOrderRequest{
		AddOrderRequest: schema.AddOrderRequest{
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
		},
	})
	if err.Error() != "2061:只有已取消、已过期、投递异常的订单才能重发" {
		t.Errorf("err:%v", err)
	}
}

func TestQueryDeliverFee(t *testing.T) {
	resp, err := client.QueryDeliverFee(&schema.QueryDeliverFeeRequest{
		AddOrderRequest: schema.AddOrderRequest{
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
		}})
	fmt.Println(resp, err)
}
