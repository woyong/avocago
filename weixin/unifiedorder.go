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
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	UNIFIEDORDER_URL = "https://api.mch.weixin.qq.com/pay/unifiedorder"
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
	TotalFee       int    `json:"total_fee,omitempty" xml:"total_fee,omitempty"`               // R. 订单总金额(分)
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
	if this.TotalFee == 0 {
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

type UnifiedOrderResp struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
	AppId      string `xml:"appid"`
	MchId      string `xml:"mch_id"`
	NonceStr   string `xml:"nonce_str"`
	Sign       string `xml:"sign"`
	ResultCode string `xml:"result_code"`
	PrepayId   string `xml:"prepay_id"`
	TradeType  string `xml:"trade_type"`
}

func (this *UnifiedOrderResp) IsSuccess() bool {
	return this.ResultCode == "SUCCESS"
}

func UnifiedOrder(payload *UnifiedOrderPayload, secretKey string) (response UnifiedOrderResp, err error) {
	if preSignErr := payload.PreSignCheck(); preSignErr != nil {
		err = preSignErr
		return
	}
	bs, _ := json.Marshal(payload)
	pm := make(map[string]interface{})
	if err1 := json.Unmarshal(bs, &pm); err1 != nil {
		err = err1
		return
	}
	sign := Sign(pm, secretKey)
	payload.Sign = sign
	XML, _ := xml.Marshal(payload)
	x := strings.Replace(string(XML), "UnifiedOrderPayload", "xml", 2)
	bytesXML := []byte(x)
	req, err2 := http.NewRequest(
		"POST",
		UNIFIEDORDER_URL,
		bytes.NewReader(bytesXML))
	if err2 != nil {
		err = err2
		return
	}
	req.Header.Set("Accept", "application/xml")
	req.Header.Set("Content-Type", "application/xml;charset=utf-8")
	c := http.Client{}
	resp, err3 := c.Do(req)
	if err3 != nil {
		err = err3
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	response = UnifiedOrderResp{}
	if err4 := xml.Unmarshal(body, &response); err4 != nil {
		err = err4
		return
	}
	if !response.IsSuccess() {
		err = errors.New(response.ReturnMsg)
		return
	}
	return
}
