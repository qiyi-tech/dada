package schema

type CityCode struct {
	CityName string `json:"cityName"`
	CityCode string `json:"cityCode"`
}

type Shop struct {
	OriginShopID   string  `json:"origin_shop_id"`  // 门店编码
	StationName    string  `json:"station_name"`    // 门店名称
	Business       int     `json:"business"`        // 支持配送的物品品类，见链接：https://newopen.imdada.cn/#/development/file/categoryList
	CityName       string  `json:"city_name"`       // 城市名称
	AreaName       string  `json:"area_name"`       // 区域名称
	StationAddress string  `json:"station_address"` // 门店地址
	Lng            float64 `json:"lng"`             // 门店经度
	Lat            float64 `json:"lat"`             // 门店纬度
	ContactName    string  `json:"contact_name"`    // 联系人姓名
	Phone          string  `json:"phone"`           // 联系人电话
	IDCard         string  `json:"id_card"`         // 联系人身份证
	Status         int     `json:"status"`          // 门店状态（1-门店激活，0-门店下线）
	SettlementType int     `json:"settlement_type"` // 门店独立结算(0-非独立结算, 1-独立结算，默认为0)，如制定门店独立结算，需给门店独立账户进行充值后才可发单，非独立结算门店通过商户总账户进行扣款
	BDPhone        string  `json:"bd_phone"`
	BDName         string  `json:"bd_name"`
	MerchantID     int64   `json:"merchant_id"`
}

type AddShopItem struct {
	StationName    string  `json:"station_name"`    // 门店名称
	Business       int     `json:"business"`        // 支持配送的物品品类
	StationAddress string  `json:"station_address"` // 门店地址
	Lng            float64 `json:"lng"`             // 门店经度
	Lat            float64 `json:"lat"`             // 门店纬度
	ContactName    string  `json:"contact_name"`    // 联系人姓名
	Phone          string  `json:"phone"`           // 联系人电话
	OriginShopID   string  `json:"origin_shop_id"`  // 门店编码,可自定义,但必须唯一;若不填写,则系统自动生成
	IDCard         string  `json:"id_card"`         // 联系人身份证
	Username       string  `json:"username"`        // 达达商家app账号(若不需要登陆app,则不用设置)
	Password       string  `json:"password"`        // 达达商家app密码(若不需要登陆app,则不用设置)
	SettlementType int     `json:"settlement_type"` // 门店独立结算(0-非独立结算, 1-独立结算，默认为0)，如制定门店独立结算，需给门店独立账户进行充值后才可发单，非独立结算门店通过商户总账户进行扣款
}

type ShopItem struct {
	Phone          string  `json:"phone"`          //联系人电话
	Business       int     `json:"business"`       //支持配送的物品品类见
	Lng            float64 `json:"lng"`            //门店经度
	Lat            float64 `json:"lat"`            //门店纬度
	StationName    string  `json:"stationName"`    //门店名称
	OriginShopId   string  `json:"originShopId"`   //门店编码
	ContactName    string  `json:"contactName"`    //联系人姓名
	StationAddress string  `json:"stationAddress"` //门店地址
	CityName       string  `json:"cityName"`       //城市名称
	AreaName       string  `json:"areaName"`       //区域名称
}

type ShopFailedItem struct {
	ShopNo   string `json:"shopNo"`   // 门店编码
	Msg      string `json:"msg"`      // 错误信息
	Code     int    `json:"code"`     // 错误码
	ShopName string `json:"shopName"` // 门店名称
}

type AddShopResp struct {
	Success     int              `json:"success"`     // 成功导入门店的数量
	SuccessList []ShopItem       `json:"successList"` // 成功导入的门店
	FailedList  []ShopFailedItem `json:"failedList"`  // 导入失败门店编号以及原因
}

type UpdateShopRequest struct {
	OriginShopID   string  `json:"origin_shop_id"`            // 门店编码
	NewShopID      string  `json:"new_shop_id"`               // 新的门店编码
	StationName    string  `json:"station_name,omitempty"`    // 门店名称
	Business       int     `json:"business,omitempty"`        // 支持配送的物品品类见链接。品类需按门店真实配送品类选择，如无法判断可咨询您的销售经理。
	StationAddress string  `json:"station_address,omitempty"` // 门店地址
	Lng            float64 `json:"lng,omitempty"`             // 门店经度
	Lat            float64 `json:"lat,omitempty"`             // 门店纬度
	ContactName    string  `json:"contact_name,omitempty"`    // 联系人姓名
	Phone          string  `json:"phone,omitempty"`           // 联系人电话
	Status         int     `json:"status,omitempty"`          // 门店状态（1-门店激活，0-门店下线）
}

type AddMerchantRequest struct {
	Mobile            string `json:"mobile"`             // 注册商户手机号,用于登陆商户后台
	CityName          string `json:"city_name"`          // 商户城市名称(如,上海)
	EnterpriseName    string `json:"enterprise_name"`    // 企业全称
	EnterpriseAddress string `json:"enterprise_address"` // 企业地址
	ContactName       string `json:"contact_name"`       // 联系人姓名
	ContactPhone      string `json:"contact_phone"`      // 联系人电话
	Email             string `json:"email"`              // 邮箱地址
}
