/*
	微信统一支付API
	Autor: woyong.j@gmail.com
*/

package weixin

import (
	"fmt"
)

type UnifiedOrder struct {
	AppId          string `json:"appid" xml:"appid"`                       // R. 应用ID
	MchId          string `json:"mch_id" xml:"mch_id"`                     // R. 商户号
	DeviceInfo     string `json:"device_info" xml:"device_info"`           // O. 设备号
	NonceStr       string `json:"nonce_str" xml:"nonce_str"`               // R. 随机字符串
	Sign           string `json:"sign" xml:"sign"`                         // R. 签名
	SignType       string `json:"sign_type" xml:"sign_type"`               // R. 签名类型,默认MD5
	Body           string `json:"body" xml:"body"`                         // R. 交易描述
	Detail         string `json:"detail" xml:"detail"`                     // O. 交易商品详情
	Attach         string `json:"attach" xml:"attach"`                     // O. 附加数据
	OutTradeNo     string `json:"out_trade_no" xml:"out_trade_no"`         // R. 商户交易号
	FeeType        string `json:"fee_type" xml:"fee_type"`                 // O. 货币类型
	TotalFee       string `json:"total_fee" xml:"total_fee"`               // R. 订单总金额(分)
	SPBillCreateIp string `json:"spbill_create_ip" xml:"spbill_create_ip"` // R. 终端IP
	TimeStart      string `json:"time_start" xml:"time_start"`             // O. 订单生成时间(yyyyMMddHHmmss)
	TimeExpire     string `json:"time_expire" xml:"time_expire"`           // O. 订单失效时间(yyyyMMddHHmmss)
	GoodsTag       string `json:"goods_tag" xml:"goods_tag"`               // O. 商品标记
	NotifyURL      string `json:"notify_url" xml:"notify_url"`             // R. 交易回调URL
	TradeType      string `json:"trade_type" xml:"trade_type"`             // R. 交易类型(APP/)
	LimitPay       string `json:"limit_pay" xml:"limit_pay"`               // O. 指定支付方式(no_credit: 不能使用信用卡支付)
	OpenID         string `json:"open_id" xml:"open_id"`                   // O. 用户标识(trade_type为JSAPI时，此参数必传)
}
