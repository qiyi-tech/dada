package dada

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	config *Config
	Debug  bool
}

type Config struct {
	AppKey    string
	AppSecret string
	SourceId  string
}

type ResponseBody struct {
	Status    string       `json:"status"`    // 响应状态，成功为"success"，失败为"fail"
	Code      int64        `json:"code"`      // 响应返回码
	Msg       string       `json:"msg"`       // 响应描述
	Result    *interface{} `json:"result"`    // 响应结果，JSON 对象
	ErrorCode int64        `json:"errorCode"` // 错误编码，与 code 一致
}

// DebugURL 测试环境域名
const DebugURL = "https://newopen.qa.imdada.cn"

// OnlineURL 线上环境域名
const OnlineURL = "https://newopen.imdada.cn"

var defaultHttpClient = &http.Client{
	Timeout: 60 * time.Second,
}

func NewClient(config *Config, debug bool) *Client {
	return &Client{
		config: config,
		Debug:  debug,
	}
}

func (c *Client) Request(path string, bodyParams interface{}, result interface{}) error {
	bodyMap := make(map[string]string)

	if bodyParams == nil {
		bodyMap["body"] = ""
	} else {
		bodyBytes, err := json.Marshal(bodyParams)
		if err != nil {
			return err
		}
		bodyMap["body"] = string(bodyBytes)
	}

	bodyMap["app_key"] = c.config.AppKey
	bodyMap["format"] = "json"
	bodyMap["timestamp"] = fmt.Sprintf("%d", time.Now().Unix())
	bodyMap["source_id"] = c.config.SourceId
	bodyMap["v"] = "1.0"
	bodyMap["signature"] = Sign(bodyMap, c.config.AppSecret)

	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(bodyMap); err != nil {
		return err
	}

	var url string
	if c.Debug {
		url = DebugURL + path
	} else {
		url = OnlineURL + path
	}

	resp, err := defaultHttpClient.Post(url, "application/json; charset=utf-8", &buf)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	if resp.StatusCode != 200 {
		return fmt.Errorf("http status code: %d", resp.StatusCode)
	}

	respBody := &ResponseBody{
		Result: &result,
	}
	if err := json.NewDecoder(resp.Body).Decode(respBody); err != nil {
		return err
	}

	if respBody.Status != "success" {
		if v, ok := ErrorMap[respBody.Code]; ok {
			return fmt.Errorf("%d:%s", respBody.Code, v)
		}
		return fmt.Errorf("%d:%s", respBody.Code, respBody.Msg)
	}

	return nil
}
