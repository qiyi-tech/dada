package dada

import "testing"

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
