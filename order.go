package dada

import "ixingkong.com/qiyi-tech/dada/schema"

// AddOrder 直接下单，将第三方订单下发到达达系统，第三方订单号必须唯一。
func (c *Client) AddOrder(req *schema.AddOrderRequest) (result *schema.AddOrderResponse, err error) {
	err = c.Request("/api/order/addOrder", req, &result)
	return result, err
}
