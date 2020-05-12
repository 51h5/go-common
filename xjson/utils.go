package xjson

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