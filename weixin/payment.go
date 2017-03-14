package weixin

import (
	"fmt"
)

func sign(payload Payload) string {
	if err := payload.PreSignCheck(); err != nil {
		fmt.Println(err)
		return ""
	}
	return ""
}
