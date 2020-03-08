# golang common package
公共基础包

## crsa
rsa 编解码的 cgo 版本

> 压测 `cd encoding/crsa && go test -v -bench=. -run=none -benchmem -memprofile memprofile.out -cpuprofile profile.out`


## rsa
rsa 编解码 golang 原生版

> 压测 `cd encoding/rsa && go test -v -bench=. -run=none -benchmem -memprofile memprofile.out -cpuprofile profile.out`