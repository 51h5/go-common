package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeQuery(t *testing.T) {
	var u, uq string
	var q map[string]string

	u = "https://www.zfkun.com/xxx/yyy.html?a=1&b=2"
	q = map[string]string {
		"c": "aaa",
		"d": "bbb",
	}
	uq = "https://www.zfkun.com/xxx/yyy.html?a=1&b=2&c=aaa&d=bbb"
	assert.Equal(t, uq, MergeQuery(u, q), "追加新参数")

	u = "https://www.zfkun.com/xxx/yyy.html?a=1&b=2"
	q = map[string]string {
		"a": "aaa",
		"c": "ccc",
	}
	uq = "https://www.zfkun.com/xxx/yyy.html?a=aaa&b=2&c=ccc"
	assert.Equal(t, uq, MergeQuery(u, q), "覆盖已有参数")

	u = "https://www.zfkun.com/xxx/yyy.html"
	q = map[string]string {
		"a": "aaa",
		"b": "bbb",
	}
	uq = "https://www.zfkun.com/xxx/yyy.html?a=aaa&b=bbb"
	assert.Equal(t, uq, MergeQuery(u, q), "原地址无参数追加")

	u = "https://www.zfkun.com/xxx/yyy.html"
	q = map[string]string {}
	uq = "https://www.zfkun.com/xxx/yyy.html"
	assert.Equal(t, uq, MergeQuery(u, q), "无参数追加")
}

func BenchmarkMergeQuery(b *testing.B) {
	u := "https://www.zfkun.com/xxx/yyy.html?a=1&b=2"
	q := map[string]string {
		"c": "aaa",
		"d": "bbb",
	}
	for i := 0; i < b.N; i++ {
		MergeQuery(u, q)
	}
}

func BenchmarkMergeQueryEmpty(b *testing.B) {
	u := "https://www.zfkun.com/xxx/yyy.html?a=1&b=2"
	var q map[string]string
	for i := 0; i < b.N; i++ {
		MergeQuery(u, q)
	}
}