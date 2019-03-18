package rsa

import (
    "crypto"
    "crypto/rand"
    "crypto/rsa"
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