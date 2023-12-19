package dada

import (
	"fmt"
	"ixingkong.com/qiyi-tech/dada/schema"
	"testing"
)

var client *Client

func TestMain(m *testing.M) {
	client = NewClient(&Config{
		AppKey:    "dadaa26197c51609631",
		AppSecret: "0b06548878bac30857ec5f70bcc89db5",
		SourceId:  "324923677",
	}, true)

	m.Run()
}

func TestGetCityCodeList(t *testing.T) {
	resp, err := client.GetCityCodeList()

	if err != nil {
		t.Errorf("err:%v", err)
		return
	}
	if len(resp) == 0 {
		t.Errorf("resp is empty")
	}
}

func TestGetBusinessMap(t *testing.T) {
	m := client.GetBusinessMap()
	if m[1] != "食品小吃" {
		t.Errorf("m:%v", m)
	}
}

func TestAddShop(t *testing.T) {
	resp, err := client.AddShop([]*schema.AddShopItem{
		{
			StationName:    "东方渔人码头肉夹馍",
			StationAddress: "上海市东方渔人码头国际中心",
			ContactName:    "李四",
			Business:       1,
			Lng:            121.475527,
			Lat:            31.251967,
			Phone:          "13817568632",
			OriginShopID:   "001",
		},
	})
	fmt.Println(resp, err)
}

func TestUpdateShop(t *testing.T) {
	resp, err := client.UpdateShop(&schema.UpdateShopRequest{
		OriginShopID: "9727fc8cd1cf4877aaa",
		ContactName:  "张三111",
		StationName:  "东方渔人码头肉夹馍-0111",
		Phone:        "15000000000",
	})
	if err != nil {
		t.Errorf("err:%v", err)
		return
	}
	if !resp {
		t.Errorf("resp:%v", resp)
	}
}
