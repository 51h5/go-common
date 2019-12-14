package rsa

import (
    "crypto/rsa"
    "crypto/x509"
    "encoding/pem"
    "errors"
)

var (
    errPublicKeyInvalid  = errors.New("public key invalid")
    errPrivateKeyInvalid = errors.New("private key invalid")
)

func ParsePKCS1PrivateKey(data []byte) (*rsa.PrivateKey, error) {
    block, _ := pem.Decode(data)
    if block == nil {
        return nil, errPrivateKeyInvalid
    }

    key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
    if err != nil {
        return nil, err
    }

    return key, nil
}

func ParsePKCS1PublicKey(data []byte) (*rsa.PublicKey, error) {
    block, _ := pem.Decode(data)
    if block == nil {
        return nil, errPublicKeyInvalid
    }

    pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
    if err != nil {
        return nil, err
    }

    key, ok := pubInterface.(*rsa.PublicKey)
    if !ok {
        return nil, errPublicKeyInvalid
    }

    return key, nil
}
