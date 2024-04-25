package order

import (
	"encoding/json"
	"fmt"
	"github.com/seerwo/jushuitan/erp/context"
	"github.com/seerwo/jushuitan/util"
)

const (
	JUSHUITAN_ORDER_LIST_QUERY                = "jushuitan.order.list.query"                //获取商家订单列表
)

type ReqOrder struct{
	PageIndex int `json:"page_index" comment:"页码"`
	PageSize int `json:"page_size" comment:"页数，最大100"`
	StartTime string `json:"start_time" comment:"开始时间"`
	EndTime string `json:"end_time" comment:"结束时间"`
	DateType int `json:"date_type" comment:"0:修改时间modified，2:订单日期order_date，3:发货时间send_date；非必填，默认0"`
	WmsCoId int `json:"wms_co_id" comment:"发货仓编号"`
	ShopId int `json:"shop_id" comment:"店铺编号"`
	Status string `json:"status" comment:"ERP订单状态：待付款：WaitPay；发货中：Delivering；被合并：Merged；异常：Question；被拆分：Split；等供销商|外仓发货：WaitOuterSent；已付款待审核：WaitConfirm；已客审待财审：WaitFConfirm；已发货：Sent；取消：Cancelled"`
	IsPaid bool `json:"is_paid" comment:"付款状态"`
	HasInvoice bool `json:"has_invoice" comment:"是否获取专票信息"`
	SoIds string `json:"so_ids" comment:"线上单号列表，多个线上单号以逗号分开"`
	OIds int `json:"o_ids" comment:"内部订单号列表"`
	IsValidPay bool `json:"is_valid_pay" comment:"查出无效支付信息"`
	StartTs int `json:"start_ts" comment:"开始时间戳，sql server中的行版本号，该字段查询防止分页过程中漏单"`
	IsGetTotal bool `json:"is_get_total" comment:"是否查询总条数，默认true"`
	Archive bool `json:"archive" comment:"是否查归档订单"`
	OrderTypes string `json:"order_types" comment:"普通订单,供销Plus 	订单类型"`
	OrderItemFlds string `json:"order_item_flds" comment:"src_combine_sku_id,presale_date 	原组合商品编码,计划发货时间"`
}

type ResOrder struct{
	PageSize int `json:"page_size"` //page_size 	int 		1 	页大小
	PageIndex int `json:"page_index"` //page_index 	int 		25 	页索引
	DataCount int `json:"data_count"` //data_count 	int 		100 	数据重量
	PageCount int `json:"page_count"` //page_count 	int 		4 	页数
	HasNext bool `json:"has_next"` //has_next 	bool 		true 	是否有下一页
	Orders []OrderData `json:"orders"` //orders 	list 			订单列表
}

type OrderData struct{
	OrderDate string `json:"order_date" comment:"下单时间"`
	ShopStatus string `json:"shop_status" comment:"线上订单状态"`
	QuestionType string `json:"question_type" comment:"异常类型"`
	ShopId int `json:"shop_id" comment:"店铺编号"`
	QuestionDesc string `json:"question_desc" comment:"异常描述"`
	SoId string `json:"so_id" comment:"线上单号"`
	Status string `json:"status" comment:"ERP订单状态：待付款：WaitPay；发货中：Delivering；被合并：Merged；异常：Question；被拆分：Split；等供销商|外仓发货：WaitOuterSent；已付款待审核：WaitConfirm；已客审待财审：WaitFConfirm；已发货：Sent；取消：Cancelled"`
	ReceiverState string  `json:"receiver_state" comment:"省"`
	ReceiverCity string `json:"receiver_city" comment:"市"`
	ReceiverDistrict string `json:"receiver_district" comment:"区"`
	SendDate string `json:"send_date" comment:"发货时间"`
	PlanDeliveryDate string `json:"plan_delivery_date" comment:"预计发货时间"`
	CreatorName string `json:"creator_name" comment:"操作业务员"`
	BuyerTaxNo string `json:"buyer_tax_no" comment:"发票税号"`
	InvoiceType string `json:"invoice_type" comment:"发票类型"`
	PayAmount string `json:"pay_amount" comment:"应付金额"`
	Freight string `json:"freight" comment:"运费"`
	BuyerMessage string `json:"buyer_message" comment:"买家留言"`
	Remark string `json:"remark" comment:"卖家留言"`
	InvoiceTitle string `json:"invoice_title" comment:"发票抬头"`
	IsCod bool `json:"is_cod" comment:"是否货到付款"`
	Type string `json:"type" comment:"订单类型"`
	PaidAmount string `json:"paid_amount" comment:"实际支付金额"`
	PayDate string `json:"pay_date" comment:"支付日期"`
	Modified string `json:"modified" comment:"订单修改时间"`
	OrderFrom string `json:"order_from" comment:"订单来源，ERP：手工下单； COPY：复制； TAOBAO；淘宝天猫； MERGE：合并； SPLIT：拆分；拆分还原； MOBILE：手机； IMPORT:导入； drp-s：供销推送； drp-d：分销推送； KWAISHOP：快手；微商城； PINDUODUO；拼多多； TOUTIAOFXG；头条放心购 "`
	LId string `json:"l_id" comment:"快递单号"`
	ShopName string `json:"shop_name" comment:"店铺名称"`
	WmsCoId int `json:"wms_co_id" comment:"发货仓编号"`
	LogisticsCompany string `json:"logistics_company" comment:"快递公司名称"`
	FreeAmount string `json:"free_amount" comment:"优惠金额"`
	CoId int `json:"co_id" comment:"公司编号"`
	Pays []PayData `json:"pays" comment:"下单时间"`
	DropCoIdTo string `json:"drop_co_id_to" comment:"下单时间"`
	EndTime string `json:"end_time" comment:"下单时间"`
	ReferrerId string `json:"referrer_id" comment:"下单时间"`
	InvoiceData string `json:"invoice_data" comment:"下单时间"`
	DropInfo string `json:"drop_info" comment:"下单时间"`
	ShopBuyerId string `json:"shop_buyer_id" comment:"下单时间"`
	SellerFlag string `json:"seller_flag" comment:"下单时间"`
	InvoiceAmount string `json:"invoice_amount" comment:"下单时间"`
	Oaid string `json:"oaid" comment:"下单时间"`
	OpenId string `json:"open_id" comment:"下单时间"`
	Node string `json:"node" comment:"下单时间"`
	ReferrerName string `json:"referrer_name" comment:"下单时间"`
	ShopSite string `json:"shop_site" comment:"下单时间"`
	DropCoIdFrom string `json:"drop_co_id_from" comment:"下单时间"`
	UnLid string `json:"un_lid" comment:"下单时间"`
	ReceiverZip string `json:"receiver_zip" comment:"下单时间"`
	Created string `json:"created" comment:"下单时间"`
	ReceiverCountry string `json:"receiver_country" comment:"下单时间"`
	Skus string `json:"skus" comment:"下单时间"`
	Shipment string `json:"shipment" comment:"下单时间"`
	Weight string `json:"weight" comment:"下单时间"`
	SignTime string `json:"sign_time" comment:"下单时间"`
	FWeight string `json:"f_weight" comment:"下单时间"`
	IsSplit bool `json:"is_split" comment:"下单时间"`
	IsMerge bool `json:"is_merge" comment:"下单时间"`
	OId string `json:"o_id" comment:"下单时间"`
	Items []ItemData `json:"items" comment:"下单时间"`
	Labels string `json:"labels" comment:"下单时间"`
	Currency string `json:"currency" comment:"下单时间"`
	LcId string `json:"lc_id" comment:"下单时间"`
	Ts int `json:"ts" comment:"下单时间"`
	MergeSoId string `json:"merge_so_id" comment:"下单时间"`
	LinkOId string `json:"link_o_id" comment:"下单时间"`
	SellerIncomeAmount float64 `json:"seller_income_amount" comment:"下单时间"`
	BuyerPaidAmount float64 `json:"buyer_paid_amount" comment:"下单时间"`
}

type ItemData struct{
	BatchId string `json:"batch_id" comment:"批次号"`
	ProducedDate string `json:"produced_date" comment:"生产日期"`
	ReferrerId string `json:"referrer_id" comment:"主播id "`
	ItemExtData string `json:"item_ext_data" comment:"商品明细拓展字段"`
	SrcCombineSkuId string `json:"src_combine_sku_id" comment:"原组合商品编码"`
	Pic string `json:"pic" comment:"图片"`
	SkuType string `json:"sku_type" comment:"商品类型"`
	ItemPayAmount string `json:"item_pay_amount" comment:"商品应付金额"`
	Remark string `json:"remark" comment:"备注"`
	Price string `json:"price" comment:"单价"`
	OuterOiId string `json:"outer_oi_id" comment:"线上子单号"`
	IsGift string `json:"is_gift" comment:"是否赠品"`
	RefundStatus string `json:"refund_status" comment:"退款状态"`
	RefundId string `json:"refund_id" comment:"退款单号"`
	ItemStatus string `json:"item_status" comment:"商品状态"`
	IId string `json:"i_id" comment:"系统商品款号"`
	ShopIId string `json:"shop_i_id" comment:"线上商品款号"`
	RawSoId string `json:"raw_so_id" comment:"原始线上单号"`
	IsPresale bool `json:"is_presale" comment:"是否预售"`
	OiId string `json:"oi_id" comment:"系统子单号"`
	PropertiesValue string `json:"properties_value" comment:"属性"`
	Amount string `json:"amount" comment:"总金额"`
	BasePrice string `json:"base_price" comment:"基础售价"`
	Qty float64 `json:"qty" comment:"数量"`
	Name string `json:"name" comment:"名称"`
	SkuId string `json:"sku_id" comment:"系统商品编码"`
	ShopSkuId string `json:"shop_sku_id" comment:"线上商品编码"`
	BuyerPaidAmount string `json:"buyer_paid_amount" comment:"买家实付"`
	SellerIncomeAmount string `json:"seller_income_amount" comment:"卖家实收"`
}

type PayData struct{
	Status string `json:"status" comment:"支付信息状态"`
	OuterPayId string `json:"outer_pay_id" comment:"外部支付单号"`
	PayDate string `json:"pay_date" comment:"支付时间"`
	Amount string `json:"amount" comment:"支付金额"`
	Payment string `json:"payment" comment:"支付渠道"`
	BuyerAccount string `json:"buyer_account" comment:"支付帐号"`
	IsOrderPay bool `json:"is_order_pay" comment:"是否支付"`
	PayId string `json:"pay_id" comment:"支付单ID"`
}

type ReqOrderAccept struct {
	OrderIds float64 `json:"order_ids"` //order_ids 	Long 	是 		待接单的订单号
}

type ResOrderAccept struct {
	util.CommonError
	Content struct {
		AcceptSqOrderResult AcceptSqOrderResult `json:"result"`
	} `json:"content"`
}



//Order struct
type Order struct {
	*context.Context
}

//NewMenu instance
func NewOrder(context *context.Context) *Order {
	order := new(Order)
	order.Context = context
	return order
}

//GetSupplierOrder get supplier order
func (order *Order) GetOrderListGet(req ReqOrderListGet) (res ResOrderListGet, err error) {

	var accessParam string
	accessParam, err = order.GetAccessParam(SQ_ORDER_LIST_GET, req)
	if err != nil {
		return
	}
	uri := fmt.Sprintf(util.HTTP_BASE_URL, order.AppID, SQ_ORDER_LIST_GET)
	var response []byte
	response, err = util.NewHTTPPost(uri, accessParam)

	if err != nil {
		return
	}
	err = json.Unmarshal(response, &res)
	if err != nil {
		return
	}
	if res.ErrCode != "0000" {
		err = fmt.Errorf("GetOrderListGet Error , errcode=%s , errmsg=%s", res.ErrCode, res.ErrMsg)
		return
	}
	return
}
