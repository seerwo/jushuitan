package order

import (
	"encoding/json"
	"fmt"
	"github.com/seerwo/jushuitan/erp/context"
	"github.com/seerwo/jushuitan/util"
)

const (
	SQ_ORDER_ACCEPT                  = "sq.order.accept"                  //订单接单
	SQ_ORDER_DETAIL_GET              = "sq.order.detail.get"              //获取买手指定订单
	SQ_ORDER_LIST_GET                = "sq.order.list.get"                //获取商家订单列表
	SQ_ORDER_DELIVERE                = "sq.order.deliver"                 //订单发货
	SQ_ORDERS_STATUS_GET             = "sq.orders.status.get"             //批量获取订单状态
	SQL_PAYMENT_PUSH                 = "sq.payment.push"                  //订单支付单推送
	SQ_CUSTOMS_CLEARANCE_STATUS_PUSH = "sq.customs.clearance.status.push" //推送订单清关状态
	SQ_ORDEREXTINFO_PUSH             = "sq.orderextinfo.push"             //推送订单扩展信息
	SQ_REFUNDBILL_LIST_GET           = "sq.refundbill.list.get"           //获取售后列表
	SQ_ORDER_FULLREFUND              = "sq.order.fullrefund"              //整单退款
	SQ_ORDER_MULTISEGMENTS_DELIVER   = "sq.order.multisegments.deliver"   //订单多段式发货
	SQ_ORDER_INFO2CUSTOMS_GET        = "sq.order.info2customs.get"        //获取订单海关报关所需相关额外信息
)

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
