package weixin

type WeTrader interface {
	WeUnifiedOrder() (UnifiedOrderPayload, string)
	WeRefund() (RefundPayload, string, string, string)
}

func UnifiedOrder(trade WeTrader) (resp UnifiedOrderResp, err error) {
	payload, sk := trade.WeUnifiedOrder()
	resp, err = unifiedorder(&payload, sk)
	return
}

func Refund(trade WeTrader) (resp RefundResponse, err error) {
	payload, sk, cert, key := trade.WeRefund()
	resp, err = refund(&payload, sk, cert, key)
	return
}
