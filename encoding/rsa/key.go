package rsa

import (
    "bytes"
    "strings"
)

const (
    pemRsaPrivateKeyPrefix = "-----BEGIN RSA PRIVATE KEY-----"
    pemRsaPrivateKeySuffix = "-----END RSA PRIVATE KEY-----"
    pemPrivateKeyPrefix    = "-----BEGIN PRIVATE KEY-----"
    pemPrivateKeySuffix    = "-----BEGIN PRIVATE KEY-----"
    pemPublicKeyPrefix     = "-----BEGIN PUBLIC KEY-----"
    pemPublicKeySuffix     = "-----END PUBLIC KEY-----"
)

func FormatPublicKey(raw string) (result []byte) {
    return formatKey(raw, pemPublicKeyPrefix, pemPublicKeySuffix)
}

func FormatRsaPrivateKey(raw string) (result []byte) {
    return formatKey(raw, pemRsaPrivateKeyPrefix, pemRsaPrivateKeySuffix)
}

func FormatPrivateKey(raw string) (result []byte) {
    return formatKey(raw, pemPrivateKeyPrefix, pemPrivateKeySuffix)
}

func formatKey(raw, prefix, suffix string) (result []byte) {
    if raw == "" {
        return nil
    }

    raw = trimKey(raw, prefix, suffix)

    var ll = 64
    var sl = len(raw)
    var c = sl / ll
    if sl%ll > 0 {
        c = c + 1
    }

    var buf bytes.Buffer
    buf.WriteString(prefix + "\n")
    for i := 0; i < c; i++ {
        var b = i * ll
        var e = b + ll
        if e > sl {
            buf.WriteString(raw[b:])
        } else {
            buf.WriteString(raw[b:e])
        }
        buf.WriteString("\n")
    }
    buf.WriteString(suffix)
    return buf.Bytes()
}

func trimKey(raw, prefix, suffix string) string {
    if raw != "" {
        raw = strings.Replace(raw, prefix, "", 1)
        raw = strings.Replace(raw, suffix, "", 1)
        raw = strings.Replace(raw, " ", "", -1)
        raw = strings.Replace(raw, "\n", "", -1)
        raw = strings.Replace(raw, "\r", "", -1)
        raw = strings.Replace(raw, "\t", "", -1)
    }
    return raw
}
