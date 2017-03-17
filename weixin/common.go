package weixin

import (
	"crypto/md5"
	"fmt"
	"sort"
	"strings"
)

func SortAndConcat(pm map[string]interface{}) string {
	keys := []string{}
	for k, v := range pm {
		if v != "" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	params := []string{}
	for _, k := range keys {
		val := pm[k]
		params = append(params, fmt.Sprintf("%v=%v", k, val))
	}
	return strings.Join(params, "&")
}

func Sign(pm map[string]interface{}, sk string) string {
	str := SortAndConcat(pm)
	str += "&key=" + sk
	fmt.Println("Prepare signature:", str)
	return fmt.Sprintf("%X", md5.Sum([]byte(str)))
}
