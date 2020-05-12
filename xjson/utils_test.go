package xjson

import (
    "bytes"
    "encoding/json"
    "testing"
)

func TestSuccess(t *testing.T) {
    j1 := []byte(`{"status":1,"data":"hello"}`)
    b1, _ := json.Marshal(Success("hello"))
    if !bytes.Equal(j1, b1) {
        t.Errorf("签名不通过: \n%s, \n%s\n", string(b1), string(j1))
        t.FailNow()
        return
    }

    j2 := []byte(`{"status":1,"data":["hello","world"]}`)
    b2, _ := json.Marshal(Success([]string{"hello", "world"}))
    if !bytes.Equal(j2, b2) {
        t.Errorf("签名不通过: \n%s, \n%s\n", string(b2), string(j2))
        t.FailNow()
        return
    }
}

func TestError(t *testing.T) {
    j3 := []byte(`{"status":0,"error":"PARAMS_INVALID","error_code":1002,"data":"hello"}`)
    b3, _ := json.Marshal(Error(1002, "PARAMS_INVALID", "hello"))
    if !bytes.Equal(j3, b3) {
        t.Errorf("签名不通过: \n%s, \n%s\n", string(b3), string(j3))
        return
    }
}
