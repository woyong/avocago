package weixin

type WeTrader interface {
	WeUnifiedOrder() (UnifiedOrderPayload, string)
	WeRefund() (RefundPayload, string, string, string)
}

func UnifiedOrder(trade *WeixinTrader) (resp UnifiedOrderResponse, err error) {
	payload, sk := trade.WeUnifiedOrder()
	resp, err = unifiedorder(payload, sk)
	return
}

func Refund(trade *WeixinTrader) (resp RefundResponse, err error) {
	payload, sk, cert, key = trade.WeRefund()
	resp, err = refund(payload, sk, cert, key)
	return
}
