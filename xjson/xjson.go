package xjson

type Xjson interface {
    Success(data interface{}) *XjsonSuccess
    Error(code int, error string, data interface{}) *XjsonError
}
