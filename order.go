package dada

import "ixingkong.com/qiyi-tech/dada/schema"

// AddOrder 直接下单，将第三方订单下发到达达系统，第三方订单号必须唯一。
func (c *Client) AddOrder(req *schema.AddOrderRequest) (result *schema.AddOrderResponse, err error) {
	err = c.Request("/api/order/addOrder", req, &result)
	return result, err
}

// ReAddOrder 针对已取消或者已过期的订单，重新下发到达达系统。
func (c *Client) ReAddOrder(req *schema.ReAddOrderRequest) (result *schema.ReAddOrderResponse, err error) {
	err = c.Request("/api/order/reAddOrder", req, &result)
	return result, err
}

// QueryDeliverFee 查询运费
func (c *Client) QueryDeliverFee(req *schema.QueryDeliverFeeRequest) (result *schema.QueryDeliverFeeResponse, err error) {
	err = c.Request("/api/order/queryDeliverFee", req, &result)
	return result, err
}

// AddAfterQuery 查询运费后下单 https://newopen.imdada.cn/#/development/file/addAfterQuery
func (c *Client) AddAfterQuery(req *schema.AddAfterQueryRequest) (result *schema.AddAfterQueryResponse, err error) {
	err = c.Request("/api/order/addAfterQuery", req, &result)
	return result, err
}
