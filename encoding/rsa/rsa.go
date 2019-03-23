package rsa

import (
    "crypto"
    "crypto/rand"
    "crypto/rsa"
    "crypto/x509"
    "encoding/asn1"
    "encoding/pem"
    "errors"
)

func SignPKCS1v15(src, key []byte, hash crypto.Hash) ([]byte, error) {
    k, err := ParsePKCS1PrivateKey(key)
    if err != nil {
        return nil, err
    }
    return SignPKCS1v15WithKey(src, k, hash)
}

func SignPKCS1v15WithKey(src []byte, key *rsa.PrivateKey, hash crypto.Hash) ([]byte, error) {
    var h = hash.New()
    h.Write(src)
    return rsa.SignPKCS1v15(rand.Reader, key, hash, h.Sum(nil))
}

func EncryptPKCS1v15(data, publicKey []byte) ([]byte, error) {
    key, err := ParsePKCS1PublicKey(publicKey)
    if err != nil {
        return nil, err
    }
    return EncryptPKCS1v15WithKey(data, key)
}

func EncryptPKCS1v15WithKey(data []byte, pub *rsa.PublicKey) ([]byte, error) {
    return rsa.EncryptPKCS1v15(rand.Reader, pub, data)
}

func DecryptPKCS1v15(ciphertext, privateKey []byte) ([]byte, error) {
    key, err := ParsePKCS1PrivateKey(privateKey)
    if err != nil {
        return nil, err
    }
    return DecryptPKCS1v15WithKey(ciphertext, key)
}

func DecryptPKCS1v15WithKey(ciphertext []byte, priv *rsa.PrivateKey) ([]byte, error) {
    return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

// RSA公钥私钥产生
func GenerateRsaKey(bits int) (pri []byte, priPkcs8 []byte, pub []byte, err error) {
    // 生成私钥
    privateKey, err := rsa.GenerateKey(rand.Reader, bits)
    if err != nil {
        return
    }
    derStream := x509.MarshalPKCS1PrivateKey(privateKey)
    pri = pem.EncodeToMemory(&pem.Block{
        Type:  "RSA PRIVATE KEY",
        Bytes: derStream,
    })
    if pri == nil {
        err = errors.New("pem encode fail")
        return
    }

    // 生成PKCS8私钥
    pk8Stream, err := MarshalPKCS8PrivateKey(derStream)
    if err != nil {
        return
    }
    priPkcs8 = pem.EncodeToMemory(&pem.Block{
        Type:  "PRIVATE KEY",
        Bytes: pk8Stream,
    })
    if priPkcs8 == nil {
        err = errors.New("pem encode fail")
        return
    }

    // 生成公钥
    publicKey := &privateKey.PublicKey
    derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
    if err != nil {
        return
    }
    // publicBuf := bytes.Buffer{}
    pub = pem.EncodeToMemory(&pem.Block{
        Type:  "PUBLIC KEY",
        Bytes: derPkix,
    })
    if pub == nil {
        err = errors.New("pem encode fail")
        return
    }

    return
}

func GenerateRsaKeyWithTrim(bits int) (pri []byte, priPkcs8 []byte, pub []byte, err error) {
    pri, priPkcs8, pub, err = GenerateRsaKey(bits)
    if err == nil {
        pri = []byte(trimKey(string(pri), pemRsaPrivateKeyPrefix, pemRsaPrivateKeySuffix))
        priPkcs8 = []byte(trimKey(string(priPkcs8), pemPrivateKeyPrefix, pemPrivateKeySuffix))
        pub = []byte(trimKey(string(pub), pemPublicKeyPrefix, pemPublicKeySuffix))
    }
    return
}


type pkcs8PrivateKey struct {
    Version             int
    PrivateKeyAlgorithm []asn1.ObjectIdentifier
    PrivateKey          []byte
}

// 由私钥获取PKCS8公钥 这种方式生成的PKCS8与OpenSSL转成的不一样，但是BouncyCastle里可用
func MarshalPKCS8PrivateKey(key []byte) ([]byte, error) {
    info := pkcs8PrivateKey{
        Version: 0,
        PrivateKeyAlgorithm: make([]asn1.ObjectIdentifier, 1),
        PrivateKey: key,
        // PrivateKey: x509.MarshalPKCS1PrivateKey(key), // key : key *rsa.PrivateKey
    }
    info.PrivateKeyAlgorithm[0] = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 1, 1}

    k, err := asn1.Marshal(info)
    if err != nil {
        return nil, err
    }
    return k, nil
}

// 由私钥获取PKCS8公钥
func MarshalPKCS8PrivateKey2(key *rsa.PrivateKey) ([]byte, error) {
    info := pkcs8PrivateKey{
        Version: 0,
        PrivateKeyAlgorithm: make([]asn1.ObjectIdentifier, 1),
        PrivateKey: x509.MarshalPKCS1PrivateKey(key),
    }
    info.PrivateKeyAlgorithm[0] = asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 1, 1}

    k, err := asn1.Marshal(info)
    if err != nil {
        return nil, err
    }
    return k, nil
}