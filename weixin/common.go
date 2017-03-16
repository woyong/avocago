package weixin

import (
	"crypto/md5"
	"fmt"
	"sort"
	"strings"
)

func SortAndConcat(pm map[string]string) string {
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
		params = append(params, k+"="+val)
	}
	return strings.Join(params, "&")
}

func Sign(pm map[string]string, sk string) string {
	str := SortAndConcat(pm)
	str += "&key=" + sk
	return fmt.Sprintf("%X", md5.Sum([]byte(str)))
}
