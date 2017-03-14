package main

import (
	"encoding/json"
	"fmt"
	"github.com/woyong/avocado/weixin"
)

func main() {
	unifiedorder := &weixin.UnifiedOrder{}
	unifiedorder.AppId = "123"
	unifiedorder.MchId = "456"
	bytes, _ := json.Marshal(unifiedorder)
	fmt.Println(string(bytes))
}
