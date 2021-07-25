package utils

import "net/url"

// MergeQuery Query参数追加合并
//
// https://xxx.com/?a=1&b=2, {c:"aaa", d:"bbb"}  => https://xxx.com/?a=1&b=2&c=aaa&d=bbb
func MergeQuery(u string, query map[string]string) string {
	if len(query) == 0 {
		return u
	}

	p, err := url.Parse(u)
	if err != nil {
		return u
	}

	q := p.Query()
	for k, v := range query {
		q.Set(k, v)
	}

	p.ForceQuery = true
	p.RawQuery = q.Encode()

	return p.String()
}