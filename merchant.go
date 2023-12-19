package dada

import "ixingkong.com/qiyi-tech/dada/schema"

// GetCityCodeList 获取达达配送开通城市列表信息
func (c *Client) GetCityCodeList() (result []*schema.CityCode, err error) {
	err = c.Request("/api/cityCode/list", nil, &result)
	return result, err
}
