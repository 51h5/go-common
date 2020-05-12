package xjson

import (
    "bytes"
    "encoding/json"
    "testing"
)

func TestEmptyMap(t *testing.T) {
    if EmptyMap == nil {
        t.Errorf("EmptyMap 为nil")
        t.FailNow()
        return
    }

    if len(EmptyMap) > 0 {
        t.Errorf("EmptyMap 非空")
        t.FailNow()
        return
    }

    j1 := []byte(`{}`)
    b1, _ := json.Marshal(EmptyMap)
    if !bytes.Equal(j1, b1) {
        t.Errorf("签名不通过: \n%s, \n%s\n", string(b1), string(j1))
        t.FailNow()
        return
    }
}

func TestEmptySlice(t *testing.T) {
    if EmptySlice == nil {
        t.Errorf("EmptySlice 为nil")
        t.FailNow()
        return
    }

    if len(EmptySlice) > 0 {
        t.Errorf("EmptySlice 非空")
        t.FailNow()
        return
    }

    j1 := []byte(`[]`)
    b1, _ := json.Marshal(EmptySlice)
    if !bytes.Equal(j1, b1) {
        t.Errorf("签名不通过: \n%s, \n%s\n", string(b1), string(j1))
        t.FailNow()
        return
    }
}
