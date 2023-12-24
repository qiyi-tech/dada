package dada

import "ixingkong.com/qiyi-tech/dada/schema"

// AddOrder 直接下单，将第三方订单下发到达达系统，第三方订单号必须唯一。 https://newopen.imdada.cn/#/development/file/add
func (c *Client) AddOrder(req *schema.AddOrderRequest) (result *schema.AddOrderResponse, err error) {
	err = c.Request("/api/order/addOrder", req, &result)
	return result, err
}

// ReAddOrder 针对已取消或者已过期的订单，重新下发到达达系统。https://newopen.imdada.cn/#/development/file/reAdd
func (c *Client) ReAddOrder(req *schema.ReAddOrderRequest) (result *schema.ReAddOrderResponse, err error) {
	err = c.Request("/api/order/reAddOrder", req, &result)
	return result, err
}

// QueryDeliverFee 查询运费 https://newopen.imdada.cn/#/development/file/readyAdd
func (c *Client) QueryDeliverFee(req *schema.QueryDeliverFeeRequest) (result *schema.QueryDeliverFeeResponse, err error) {
	err = c.Request("/api/order/queryDeliverFee", req, &result)
	return result, err
}

// AddAfterQuery 查询运费后下单 https://newopen.imdada.cn/#/development/file/addAfterQuery
func (c *Client) AddAfterQuery(req *schema.AddAfterQueryRequest) (result *schema.AddAfterQueryResponse, err error) {
	err = c.Request("/api/order/addAfterQuery", req, &result)
	return result, err
}

// AddTip 加小费 https://newopen.imdada.cn/#/development/file/addTip
func (c *Client) AddTip(req *schema.AddTipRequest) (result bool, err error) {
	err = c.Request("/api/order/addTip", req, &result)
	if err != nil {
		return false, err
	}
	return true, nil
}

// QueryOrder 查询订单详情 https://newopen.imdada.cn/#/development/file/statusQuery
func (c *Client) QueryOrder(req *schema.QueryOrderRequest) (result *schema.QueryOrderResponse, err error) {
	err = c.Request("/api/order/status/query", req, &result)
	return result, err
}

// CancelOrder 取消订单 https://newopen.imdada.cn/#/development/file/formalCancel
func (c *Client) CancelOrder(req *schema.CancelOrderRequest) (result *schema.CancelOrderResponse, err error) {
	err = c.Request("/api/order/formalCancel", req, &result)
	return result, err
}

// AppointOrder 追加订单 https://newopen.imdada.cn/#/development/file/appointOrder
func (c *Client) AppointOrder(req *schema.AppointOrderRequest) (result *schema.AppointOrderResponse, err error) {
	err = c.Request("/api/order/appoint/exist", req, &result)
	return result, err
}

// CancelAppointOrder 取消追加订单 https://newopen.imdada.cn/#/development/file/appointOrderCancel
func (c *Client) CancelAppointOrder(req *schema.CancelAppointOrderRequest) (result bool, err error) {
	err = c.Request("/api/order/appoint/cancel", req, &result)
	if err != nil {
		return false, err
	}
	return true, nil
}

// QueryAppointTransporter 查询可追加骑士 https://newopen.imdada.cn/#/development/file/listTransportersToAppoint
func (c *Client) QueryAppointTransporter(req *schema.QueryAppointTransporterRequest) (result []*schema.Transporter, err error) {
	err = c.Request("/api/order/appoint/list/transporter", req, &result)
	return result, err
}

// Complaint 商家投诉达达 https://newopen.imdada.cn/#/development/file/complaintDada
func (c *Client) Complaint(req *schema.ComplaintRequest) (result bool, err error) {
	err = c.Request("/api/complaint/dada", req, &result)
	if err != nil {
		return false, err
	}
	return true, nil
}
