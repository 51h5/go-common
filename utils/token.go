package utils

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
)

const (
	tokenRetryMax = 10 // 字符集的长度
)

// 简易token生成 (仅小写)
func NewToken() string {
	token := make([]byte, 16)

	var retry int
	for {
		_, err := rand.Read(token)
		if err != nil {
			if retry > tokenRetryMax {
				return ""
			}
			retry++
			continue
		}

		return hex.EncodeToString(token)
	}
}

// 简易token生成V2 (含大小写)
func NewTokenV2() string {
	token := make([]byte, 16)

	var retry int
	for {
		_, err := rand.Read(token)
		if err != nil {
			if retry > tokenRetryMax {
				return ""
			}
			retry++
			continue
		}

		return base64.RawURLEncoding.EncodeToString(token)
	}
}
