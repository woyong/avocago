package main

import (
	"github.com/woyong/avocado/weixin"
)

func main() {
	payload := weixin.UnifiedOrderPayload{}
	payload.AppId = "123"
	payload.MchId = "456"
	weixin.UnifiedOrder(&payload, "secret key")
}
