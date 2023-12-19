package dada

import (
	"crypto/md5"
	"encoding/hex"
	"sort"
	"strings"
)

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func Sign(p map[string]string, appSecret string) string {
	var keys []string
	for k := range p {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var strParam string
	for _, k := range keys {
		if k != "sign" {
			strParam += k + p[k]
		}
	}
	strParam = appSecret + strParam + appSecret

	return strings.ToUpper(Md5(strParam))
}
