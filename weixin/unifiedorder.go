/*
	微信统一支付API
	Autor: woyong.j@gmail.com
*/

package weixin

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type UnifiedOrderPayload struct {
	AppId          string `json:"appid,omitempty" xml:"appid,omitempty"`                       // R. 应用ID
	MchId          string `json:"mch_id,omitempty" xml:"mch_id,omitempty"`                     // R. 商户号
	DeviceInfo     string `json:"device_info,omitempty" xml:"device_info,omitempty"`           // O. 设备号
	NonceStr       string `json:"nonce_str,omitempty" xml:"nonce_str,omitempty"`               // R. 随机字符串
	Sign           string `json:"sign,omitempty" xml:"sign,omitempty"`                         // R. 签名
	SignType       string `json:"sign_type,omitempty" xml:"sign_type,omitempty"`               // R. 签名类型,默认MD5
	Body           string `json:"body,omitempty" xml:"body,omitempty"`                         // R. 交易描述
	Detail         string `json:"detail,omitempty" xml:"detail,omitempty"`                     // O. 交易商品详情
	Attach         string `json:"attach,omitempty" xml:"attach,omitempty"`                     // O. 附加数据
	OutTradeNo     string `json:"out_trade_no,omitempty" xml:"out_trade_no,omitempty"`         // R. 商户交易号
	FeeType        string `json:"fee_type,omitempty" xml:"fee_type,omitempty"`                 // O. 货币类型
	TotalFee       string `json:"total_fee,omitempty" xml:"total_fee,omitempty"`               // R. 订单总金额(分)
	SPBillCreateIp string `json:"spbill_create_ip,omitempty" xml:"spbill_create_ip,omitempty"` // R. 终端IP
	TimeStart      string `json:"time_start,omitempty" xml:"time_start,omitempty"`             // O. 订单生成时间(yyyyMMddHHmmss)
	TimeExpire     string `json:"time_expire,omitempty" xml:"time_expire,omitempty"`           // O. 订单失效时间(yyyyMMddHHmmss)
	GoodsTag       string `json:"goods_tag,omitempty" xml:"goods_tag,omitempty"`               // O. 商品标记
	NotifyURL      string `json:"notify_url,omitempty" xml:"notify_url,omitempty"`             // R. 交易回调URL
	TradeType      string `json:"trade_type,omitempty" xml:"trade_type,omitempty"`             // R. 交易类型(APP/)
	LimitPay       string `json:"limit_pay,omitempty" xml:"limit_pay,omitempty"`               // O. 指定支付方式(no_credit: 不能使用信用卡支付)
	OpenID         string `json:"open_id,omitempty" xml:"open_id,omitempty"`                   // O. 用户标识(trade_type为JSAPI时，此参数必传)
}

func (this *UnifiedOrderPayload) isWAP() bool {
	return this.TradeType == "JSAPI"
}

func (this *UnifiedOrderPayload) PreSignCheck() (err error) {
	if this.AppId == "" {
		err = errors.New("Missing required parameters: appid")
		return
	}
	if this.MchId == "" {
		err = errors.New("Missing required parameters: mch_id")
		return
	}
	if this.Body == "" {
		err = errors.New("Missing required parameters: body")
		return
	}
	if this.NonceStr == "" {
		err = errors.New("Missing required parameters: nonce_str")
		return
	}
	if this.OutTradeNo == "" {
		err = errors.New("Missing required parameters: out_trade_no")
		return
	}
	if this.TotalFee == "" {
		err = errors.New("Missing required parameters: total_fee")
		return
	}
	if this.SPBillCreateIp == "" {
		err = errors.New("Missing required parameters: spbill_create_ip")
		return
	}
	if this.NotifyURL == "" {
		err = errors.New("Missing required parameters: notify_url")
		return
	}
	if this.TradeType == "" {
		err = errors.New("Missing required parameters: trade_type")
		return
	}
	if this.isWAP() && this.OpenID == "" {
		err = errors.New("Missing required paramters for WAP payment: openid")
		return
	}
	return
}

func UnifiedOrder(payload *UnifiedOrderPayload, secretKey string) string {
	bs, _ := json.Marshal(payload)
	pm := make(map[string]string)
	err := json.Unmarshal(bs, &pm)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	sign := Sign(pm, secretKey)
	payload.Sign = sign
	payload.SignType = "MD5"
	XML, _ := xml.Marshal(payload)
	x := strings.Replace(string(XML), "UnifiedOrderPayload", "xml", 2)
	bytesXML := []byte(x)
	req, err := http.NewRequest("POST", "https://api.mch.weixin.qq.com/pay/unifiedorder", bytes.NewReader(bytesXML))
	if err != nil {
		fmt.Println(err)
		return ""
	}
	req.Header.Set("Accept", "application/xml")
	req.Header.Set("Content-Type", "application/xml;charset=utf-8")
	c := http.Client{}
	resp, resp_err := c.Do(req)
	if resp_err != nil {
		fmt.Println(resp_err)
		return ""
	}
	fmt.Println(resp)
	return ""
}
