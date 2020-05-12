package errors

import (
    xjson "51h5.com/common/xjson"
    "encoding/json"
)

type Xjson struct {
    err *xjson.XjsonError
}

func (e *Xjson) Error() string {
    b, _ := json.Marshal(e.err)
    return string(b)
}

func XjsonError(e *xjson.XjsonError) error {
    return &Xjson{err: e}
}
