package rsa

import (
    "crypto"
    "crypto/rsa"
    "encoding/base64"
    "fmt"
    "io/ioutil"
    "testing"
)

func TestGenerateRsaKey(t *testing.T) {
    var pri, priPkcs8, pub, d, v []byte
    var err error

    d = []byte("hdlgame_01,n-1234567,1234567890")

    pri, priPkcs8, pub, err = GenerateRsaKey(1024)
    if err != nil {
        t.Fatalf("1024位生成失败: %s", err.Error())
    }

    v, err = EncryptPKCS1v15(d, pub)
    fmt.Println(string(pri))
    fmt.Println(string(priPkcs8))
    // fmt.Println(string(pub))
    // fmt.Println(string(v))
    fmt.Println(base64.StdEncoding.EncodeToString(v))
    // fmt.Println("=======================================")

    pri, priPkcs8, pub, err = GenerateRsaKey(2048)
    if err != nil {
        t.Errorf("2048位生成失败: %s", err.Error())
        t.FailNow()
    }
    v, err = EncryptPKCS1v15(d, pub)
    // fmt.Println(string(pri))
    // fmt.Println(string(priPkcs8))
    // fmt.Println(string(pub))
    // fmt.Println(string(v))
    fmt.Println(base64.StdEncoding.EncodeToString(v))
    // fmt.Println("=======================================")
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
    // fmt.Printf("encrypt origin: %s\n", v)
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
    // fmt.Printf("encrypt origin: %s\n", v)
    fmt.Printf("encrypt base64: %s\n", base64.StdEncoding.EncodeToString(v))
    fmt.Println("=======================================")
}

var (
    ori         = []byte("aaa")
    oriSignRsa  = []byte("IIjjNLvi7m1HgvjkM4mJ/gOQLQNflgW5pyT15PEpce2oBvBdlqtoZpGIuKmYgfo1kgO6kIinq07sjnqpdKL/5UEQQXrx+ZnO3TjDEmkdQp0BzPr8S/U9BGJtFfty4p7032jV403riqDS5hxeLflIzfBB3bsqmn44e0nuOsH5v9g=")
    oriSignRsa2 = []byte("Xt5ygodmn3G1JOh0mFN8E2SEgBO17mcxH/x8o4IzrWl/tMX9f9XTMgIzLMvUE9IWAzDmKDFp+6RFAEQqsJv7Jc8ssFWbLgPExg/8DgsMGJxNqm3fdPs+0yzT7iSaFXQT3ZliVgzUOViPTcTts4IDZVrkssOmhaZ2w+8F2QhoBqKEzz3jMkkq0ox9q76VeuiXvIYenHlGgJ4RPyM2Ey9rCDcQNkc0s0++ErGP+2Om6YykktDVdMfYLBVpJC+JlUY39WohH7ZOwtazVcFY8yJc4NiXXYmA8Ym3JW93tLcocnncSzKygNcYvVSxs05D3l048M8FOJzQt6GZ46FuQN3TFw==")
)

func TestRsa(t *testing.T) {
    pri, _ := loadPrivateKey("../crsa/rsa_private.pem")
    if pri == nil {
        t.Fatalf("rsa private load: %v", pri)
    }

    b, _ := SignPKCS1v15WithKey(ori, pri, crypto.SHA1)
    sign := base64.StdEncoding.EncodeToString(b)

    t.Logf("rsa sign: %s\n", sign)
    t.Logf("ori sign: %s\n", oriSignRsa)

    if sign != string(oriSignRsa) {
        t.Fatalf("rsa sign fail")
    }

    pub, _ := loadPublicKey("../crsa/rsa_public.pem")
    if err := VerifyPKCS1v15WithKey(ori, b, pub, crypto.SHA1); err != nil {
        t.Fatalf("rsa verify fail")
    }
}

func BenchmarkRsaSign(b *testing.B) {
    pri, _ := loadPrivateKey("../crsa/rsa_private.pem")

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        sign, _ := SignPKCS1v15WithKey(ori, pri, crypto.SHA1)
        _ = base64.StdEncoding.EncodeToString(sign)
    }
}

func BenchmarkRsaVerify(b *testing.B) {
    pub, _ := loadPublicKey("../crsa/rsa_public.pem")

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = VerifyPKCS1v15WithKey(ori, oriSignRsa, pub, crypto.SHA1)
    }
}

func BenchmarkRsa(b *testing.B) {
    pri, _ := loadPrivateKey("../crsa/rsa_private.pem")
    pub, _ := loadPublicKey("../crsa/rsa_public.pem")

    b.ResetTimer()
    var sign []byte
    for i := 0; i < b.N; i++ {
        sign, _ = SignPKCS1v15WithKey(ori, pri, crypto.SHA1)
        _ = base64.StdEncoding.EncodeToString(sign)
        _ = VerifyPKCS1v15WithKey(ori, sign, pub, crypto.SHA1)
        // fmt.Printf("签名: %s, %v\n", s, err)
    }
}

func TestRsa2(t *testing.T) {
    pri, _ := loadPrivateKey("../crsa/rsa2_private.pem")
    if pri == nil {
        t.Fatalf("rsa2 private load: %v", pri)
    }

    b, _ := SignPKCS1v15WithKey(ori, pri, crypto.SHA256)
    sign := base64.StdEncoding.EncodeToString(b)

    if sign != string(oriSignRsa2) {
        t.Fatalf("rsa2 sign fail")
    }

    pub, _ := loadPublicKey("../crsa/rsa2_public.pem")
    if err := VerifyPKCS1v15WithKey(ori, b, pub, crypto.SHA256); err != nil {
        t.Fatalf("rsa2 verify fail")
    }
}

func BenchmarkRsa2Sign(b *testing.B) {
    pri, _ := loadPrivateKey("../crsa/rsa2_private.pem")

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        sign, _ := SignPKCS1v15WithKey(ori, pri, crypto.SHA256)
        _ = base64.StdEncoding.EncodeToString(sign)
    }
}

func BenchmarkRsa2Verify(b *testing.B) {
    pub, _ := loadPublicKey("../crsa/rsa2_public.pem")

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = VerifyPKCS1v15WithKey(ori, oriSignRsa2, pub, crypto.SHA1)
    }
}

func BenchmarkRsa2(b *testing.B) {
    pri, _ := loadPrivateKey("../crsa/rsa2_private.pem")
    pub, _ := loadPublicKey("../crsa/rsa2_public.pem")

    var sign []byte
    for i := 0; i < b.N; i++ {
        sign, _ = SignPKCS1v15WithKey(ori, pri, crypto.SHA256)
        _ = base64.StdEncoding.EncodeToString(sign)
        _ = VerifyPKCS1v15WithKey(ori, sign, pub, crypto.SHA256)
        // fmt.Printf("签名: %s, %v\n", s, err)
    }
}

func loadPrivateKey(p string) (*rsa.PrivateKey, error) {
    buf, _ := ioutil.ReadFile(p)
    return ParsePKCS1PrivateKey(FormatRsaPrivateKey(string(buf)))
}

func loadPublicKey(p string) (*rsa.PublicKey, error) {
    buf, _ := ioutil.ReadFile(p)
    return ParsePKCS1PublicKey(FormatPublicKey(string(buf)))
}
