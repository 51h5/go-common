package rsa

import (
    "encoding/base64"
    "fmt"
    "testing"
)

func TestGenerateRsaKey(t *testing.T) {
    var pri, priPkcs8, pub, d, v []byte
    var err error

    d = []byte("hdlgame_01,n-1234567,1234567890")

    pri, priPkcs8, pub, err = GenerateRsaKey(1024)
    if err != nil {
        t.Errorf("1024位生成失败: %s", err.Error())
        t.FailNow()
    }
    v, err = EncryptPKCS1v15(d, pub)
    fmt.Println(string(pri))
    fmt.Println(string(priPkcs8))
    fmt.Println(string(pub))
    fmt.Println(string(v))
    fmt.Println(base64.StdEncoding.EncodeToString(v))
    fmt.Println("=======================================")

    pri, priPkcs8, pub, err = GenerateRsaKey(2048)
    if err != nil {
        t.Errorf("2048位生成失败: %s", err.Error())
        t.FailNow()
    }
    v, err = EncryptPKCS1v15(d, pub)
    fmt.Println(string(pri))
    fmt.Println(string(priPkcs8))
    fmt.Println(string(pub))
    fmt.Println(string(v))
    fmt.Println(base64.StdEncoding.EncodeToString(v))
    fmt.Println("=======================================")
}

func TestGenerateRsaKeyWithTrim(t *testing.T) {
    var pri, priPkcs8, pub, d, v []byte
    var err error

    d = []byte("hdlgame_01,n-1234567,1234567890")

    pri, priPkcs8, pub, err = GenerateRsaKeyWithTrim(1024)
    if err != nil {
        t.Errorf("1024位生成失败: %s", err.Error())
        t.FailNow()
    }
    fmt.Printf("privateKey: %s\n", pri)
    fmt.Printf("privatePkcs8Key: %s\n", priPkcs8)
    fmt.Printf("publicKey: %s\n", pub)
    v, err = EncryptPKCS1v15(d, FormatPublicKey(string(pub)))
    if err != nil {
        t.Errorf("encrypt fail: %s", err.Error())
        t.FailNow()
    }
    fmt.Printf("encrypt origin: %s\n", v)
    fmt.Printf("encrypt base64: %s\n", base64.StdEncoding.EncodeToString(v))
    fmt.Println("=======================================")

    pri, priPkcs8, pub, err = GenerateRsaKeyWithTrim(2048)
    if err != nil {
        t.Errorf("2048位生成失败: %s", err.Error())
        t.FailNow()
    }
    fmt.Printf("privateKey: %s\n", pri)
    fmt.Printf("privatePkcs8Key: %s\n", priPkcs8)
    fmt.Printf("publicKey: %s\n", pub)
    v, err = EncryptPKCS1v15(d, FormatPublicKey(string(pub)))
    if err != nil {
        t.Errorf("encrypt fail: %s", err.Error())
        t.FailNow()
    }
    fmt.Printf("encrypt origin: %s\n", v)
    fmt.Printf("encrypt base64: %s\n", base64.StdEncoding.EncodeToString(v))
    fmt.Println("=======================================")
}