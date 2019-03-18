package rsa

import (
    "crypto/rsa"
    "crypto/x509"
    "encoding/pem"
    "errors"
)

func ParsePKCS1PrivateKey(data []byte) (*rsa.PrivateKey, error) {
    block, _ := pem.Decode(data)
    if block == nil {
        return nil, errors.New("private key invalid")
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
        return nil, errors.New("public key invalid")
    }

    pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
    if err != nil {
        return nil, err
    }

    key, ok := pubInterface.(*rsa.PublicKey)
    if !ok {
        return nil, errors.New("public key invalid")
    }

    return key, nil
}
