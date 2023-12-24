package schema

type OrderProduct struct {
	SkuName      string  `json:"sku_name"`       // 商品名称，限制长度128
	SrcProductNo string  `json:"src_product_no"` // 商品编码，限制长度64
	Count        float64 `json:"count"`          // 商品数量，精确到小数点后两位
	Unit         string  `json:"unit"`           // 商品单位，默认：件
}

type AddOrderRequest struct {
	ShopNo                string          `json:"shop_no"`                  // 门店编号
	OriginId              string          `json:"origin_id"`                // 第三方订单ID
	CityCode              string          `json:"city_code"`                // 订单所在城市的code
	CargoPrice            float64         `json:"cargo_price"`              // 订单金额（单位：元）
	IsPrepay              int             `json:"is_prepay"`                // 是否需要垫付 1:是 0:否 (垫付订单金额，非运费)
	ReceiverName          string          `json:"receiver_name"`            // 收货人姓名
	ReceiverAddress       string          `json:"receiver_address"`         // 收货人地址
	ReceiverLat           float64         `json:"receiver_lat"`             // 收货人地址纬度
	ReceiverLng           float64         `json:"receiver_lng"`             // 收货人地址经度
	Callback              string          `json:"callback"`                 // 回调URL
	CargoWeight           float64         `json:"cargo_weight"`             // 订单重量
	ReceiverPhone         string          `json:"receiver_phone"`           // 收货人手机号
	ReceiverTel           string          `json:"receiver_tel"`             // 收货人座机号
	Tips                  float64         `json:"tips"`                     // 小费（单位：元，精确小数点后一位）
	Info                  string          `json:"info"`                     // 订单备注
	CargoType             int64           `json:"cargo_type"`               // 请选择门店真实配送品类，如无法判断可咨询您的销售经理。
	CargoNum              int64           `json:"cargo_num"`                // 订单商品数量
	InvoiceTitle          string          `json:"invoice_title"`            // 发票抬头
	OriginMark            string          `json:"origin_mark"`              // 订单来源标示（只支持字母，最大长度为10）
	OriginMarkNo          string          `json:"origin_mark_no"`           // 订单来源编号
	IsUseInsurance        int             `json:"is_use_insurance"`         // 是否使用保价费
	IsFinishCodeNeeded    int             `json:"is_finish_code_needed"`    // 收货码 0：不需要；1：需要。
	DelayPublishTime      int64           `json:"delay_publish_time"`       // 预约发单时间
	IsExpectFinishOrder   int             `json:"is_expect_finish_order"`   // 是否根据期望送达时间预约发单（0-否，即时发单；1-是，预约发单），如传1则期望送达时间必传。
	ExpectFinishTimeLimit int64           `json:"expect_finish_time_limit"` // 期望送达时间（单位秒，不早于当前时间）
	IsDirectDelivery      int             `json:"is_direct_delivery"`       // 是否选择直拿直送 0：不需要；1：需要。
	ProductList           []*OrderProduct `json:"product_list"`             // 订单商品明细
	PickUpPos             string          `json:"pick_up_pos"`              // 货架信息,该字段可在骑士APP订单备注中展示
	FetchCode             string          `json:"fetch_code"`               // 取货码，在骑士取货时展示给骑士，门店通过取货码在骑士取货阶段核实骑士
}

type AddOrderResponse struct {
	Distance     float64 `json:"distance"`     // 距离（单位：米）
	Fee          float64 `json:"fee"`          // 实际运费（单位：元），运费减去优惠券费用
	DeliverFee   float64 `json:"deliverFee"`   // 运费(单位：元)
	CouponFee    float64 `json:"couponFee"`    // 优惠券费用(单位：元)
	Tips         float64 `json:"tips"`         // 小费(单位：元)
	InsuranceFee float64 `json:"insuranceFee"` // 保价费(单位：元)
}

type ReAddOrderRequest struct {
	AddOrderRequest
}

type ReAddOrderResponse struct {
	AddOrderResponse
}

type QueryDeliverFeeRequest struct {
	AddOrderRequest
}

type QueryDeliverFeeResponse struct {
	Distance     float64 `json:"distance"`     // 距离（单位：米）
	DeliveryNo   string  `json:"deliveryNo"`   // 平台订单号
	Fee          float64 `json:"fee"`          // 实际运费（单位：元），运费减去优惠券费用
	DeliverFee   float64 `json:"deliverFee"`   // 运费(单位：元)
	CouponFee    float64 `json:"couponFee"`    // 优惠券费用(单位：元)
	Tips         float64 `json:"tips"`         // 小费(单位：元)
	InsuranceFee float64 `json:"insuranceFee"` // 保价费(单位：元)
}

type AddAfterQueryRequest struct {
	DeliveryNo  string `json:"deliveryNo"`  // 平台订单编号
	EnableReset bool   `json:"enableReset"` // 是否重置订单备注等字段。true：重置，false或缺失：不重置
	Info        string `json:"info"`        // 订单备注内容
}

type AddAfterQueryResponse struct {
}

type AddTipRequest struct {
	OrderId string  `json:"order_id"` // 第三方订单编号
	Tips    float64 `json:"tips"`     // 小费金额(精确到小数点后一位，单位：元)
	Info    string  `json:"info"`     // 备注(字段最大长度：512)
}

type QueryOrderRequest struct {
	OrderId string `json:"order_id"` // 第三方订单编号
}

type QueryOrderResponse struct {
	OrderId          string  `json:"orderId"`          // 第三方订单编号
	StatusCode       int     `json:"statusCode"`       // 订单状态(待接单＝1,待取货＝2,配送中＝3,已完成＝4,已取消＝5, 指派单=8,妥投异常之物品返回中=9, 妥投异常之物品返回完成=10, 骑士到店=100,创建达达运单失败=1000 可参考文末的状态说明）
	StatusMsg        string  `json:"statusMsg"`        // 订单状态
	TransporterName  string  `json:"transporterName"`  // 骑手姓名
	TransporterPhone string  `json:"transporterPhone"` // 骑手电话
	TransporterLng   string  `json:"transporterLng"`   // 骑手经度
	TransporterLat   string  `json:"transporterLat"`   // 骑手纬度
	DeliveryFee      float64 `json:"deliveryFee"`      // 订单总费用，包含运费、小费、保价费
	Tips             float64 `json:"tips"`             // 小费,单位为元
	CouponFee        float64 `json:"couponFee"`        // 优惠券费用,单位为元
	InsuranceFee     float64 `json:"insuranceFee"`     // 保价费,单位为元
	ActualFee        float64 `json:"actualFee"`        // 订单实际费用消耗，订单总费用减优惠券
	Distance         float64 `json:"distance"`         // 配送距离,单位为米
	CreateTime       string  `json:"createTime"`       // 发单时间
	AcceptTime       string  `json:"acceptTime"`       // 接单时间,若未接单,则为空
	FetchTime        string  `json:"fetchTime"`        // 取货时间,若未取货,则为空
	FinishTime       string  `json:"finishTime"`       // 送达时间,若未送达,则为空
	CancelTime       string  `json:"cancelTime"`       // 取消时间,若未取消,则为空
	OrderFinishCode  string  `json:"orderFinishCode"`  // 收货码
	DeductFee        float64 `json:"deductFee"`        // 违约金
	DeliveryId       int64   `json:"deliveryId"`
	ArriveTime       string  `json:"arriveTime"`
	SupplierName     string  `json:"supplierName"`
	SupplierLat      string  `json:"supplierLat"`
	SupplierLng      string  `json:"supplierLng"`
	SupplierAddress  string  `json:"supplierAddress"`
	SupplierPhone    string  `json:"supplierPhone"`
	TransporterId    int64   `json:"transporterId"`
}

type CancelOrderRequest struct {
	OrderId        string `json:"order_id"`         // 第三方订单编号
	CancelReasonId int    `json:"cancel_reason_id"` // 取消原因ID
	CancelReason   string `json:"cancel_reason"`    // 取消原因(当取消原因ID为其他时，此字段必填)
}

type CancelOrderResponse struct {
	DeductFee float64 `json:"deduct_fee"` // 扣除的违约金(单位：元)
}
