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

func Success(data interface{}) *XjsonSuccess {
    return &XjsonSuccess{
        Status: 1,
        Data:   data,
    }
}

func Error(code int, error string, data interface{}) *XjsonError {
    return &XjsonError{
        Status:    0,
        ErrorCode: code,
        Error:     error,
        Data:      data,
    }
}
