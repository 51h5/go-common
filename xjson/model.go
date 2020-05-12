package xjson

type XjsonSuccess struct {
    Status uint8       `json:"status"`
    Data   interface{} `json:"data,omitempty"`
}

type XjsonError struct {
    Status    uint8       `json:"status"`
    Error     string      `json:"error"`
    ErrorCode int         `json:"error_code"`
    Data      interface{} `json:"data,omitempty"`
}
