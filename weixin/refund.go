package weixin

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	RefundURL = "https://api.mch.weixin.qq.com/secapi/pay/refund"
)

type RefundPayload struct {
	AppID         string `json:"appid" xml:"appid"`                                         // R. APPID
	MchID         string `json:"mch_id" xml:"mch_id"`                                       // R. 商户号
	DeviceInfo    string `json:"device_info,omitempty "xml:"device_info,omitempty"`         // O. 设备号
	NonceStr      string `json:"nonce_str" xml:"nonce_str"`                                 // R. 随机字符串
	Sign          string `json:"sign,omitempty" xml:"sign,omitempty"`                       // R. 签名
	SignType      string `json:"sign_type,omitempty" xml:"sign_type,omitempty"`             // O. 签名类型
	OutTradeNo    string `json:"out_trade_no,omitempty "xml:"out_trade_no,omitempty"`       // R. 商户订单号
	TransactionID string `json:"transaction_id,omitempty" xml:"transaction_id,omitempty"`   // C. 微信订单号
	OutRefundNo   string `json:"out_refund_no" xml:"out_refund_no"`                         // C. 商户退款号
	TotalFee      int    `json:"total_fee" xml:"total_fee"`                                 // R. 订单金额
	RefundFee     int    `json:"refund_fee" xml:"refund_fee"`                               // R. 退款金额
	OpUserID      string `json:"op_user_id" xml:"op_user_id"`                               // R. 操作员账号
	RefundAccount string `json:"refund_account,omitempty" xml:"refund_account,omitempty"`   // O. 退款资金来源
	RfundFeeType  string `json:"refund_fee_type,omitempty" xml:"refund_fee_type,omitempty"` // O. 货币类型
}

type RefundResponse struct {
	ReturnCode          string `xml:"return_code"`
	ReturnMsg           string `xml:"return_msg"`
	ResultCode          string `xml:"result_code"`
	ErrCode             string `xml:"err_code"`
	ErrCodeDes          string `xml:"err_code_des"`
	AppID               string `xml:"appid"`
	MchID               string `xml:"mch_id"`
	DeviceInfo          string `xml:"device_info"`
	NonceStr            string `xml:"nonce_str"`
	Sign                string `xml:"sign"`
	TransactionID       string `xml:"transaction_id"`
	OutTradeNo          string `xml:"out_trade_no"`
	OutRefundNo         string `xml:"out_refund_no"`
	RefundId            string `xml:"refund_id"`
	RefundChannel       string `xml:"refund_channel"`
	RefundFee           int    `xml:"refund_fee"`
	TotalFee            int    `xml:"total_fee"`
	SettlementRefundFee int    `xml:"settlement_refund_fee"`
	FeeType             string `xml:"fee_type"`
	CashFee             int    `xml:"cash_fee"`
}

func (this *RefundResponse) IsSuccess() bool {
	return this.ReturnCode == "SUCCESS" && this.ResultCode == "SUCCESS"
}

func (this *RefundPayload) PreSignCheck() (err error) {
	if this.AppID == "" {
		err = errors.New("Missing required parameters: appid")
		return
	}
	return
}

func refund(payload *RefundPayload, secretKey string, cert string, key string) (response RefundResponse, err error) {
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
	req, err2 := http.NewRequest(
		"POST",
		RefundURL,
		bytes.NewReader(XML))
	if err2 != nil {
		err = err2
		return
	}
	req.Header.Set("Accept", "application/xml")
	req.Header.Set("Content-Type", "application/xml;charset=utf-8")
	c := http.Client{}
	tlsConfig, err := NewTLSConfig(cert, key)
	if err != nil {
		return
	}
	c.Transport = &http.Transport{TLSClientConfig: tlsConfig}
	resp, err3 := c.Do(req)
	if err3 != nil {
		err = err3
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	response = RefundResponse{}
	if err4 := xml.Unmarshal(body, &response); err4 != nil {
		err = err4
		return
	}
	if !response.IsSuccess() {
		err = errors.New(response.ErrCodeDes)
		return
	}
	return
}
