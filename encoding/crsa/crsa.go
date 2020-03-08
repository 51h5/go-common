package crsa

// #cgo LDFLAGS: ${SRCDIR}/c/libcrsa.a -L/usr/local/Cellar/openssl@1.1/1.1.1d/lib/ -lcrypto -lssl -lstdc++
// #include "./c/rsa.h"
// #include <stdlib.h>
import "C"

import (
    "unsafe"
)

func LoadPrivateKey(f string) unsafe.Pointer {
    s := C.CString(f)
    defer C.free(unsafe.Pointer(s))

    return C.LoadPrivateKey(s)
}

func LoadPublicKey(f string) unsafe.Pointer {
    s := C.CString(f)
    defer C.free(unsafe.Pointer(s))

    return C.LoadPublicKey(s)
}

func RsaSign(pri unsafe.Pointer, ori string) string {
    s := C.CString(ori)
    defer C.free(unsafe.Pointer(s))
    // s := (*C.char)(unsafe.Pointer(&ori))

    sign := C.RsaSign(pri, s)
    if sign == nil {
        return ""
    }

    // return C.GoBytes(unsafe.Pointer(sign), 172)
    return C.GoString(sign)
}

func RsaVerify(pub unsafe.Pointer, ori, sign string) bool {
    s1, s2 := C.CString(ori), C.CString(sign)
    defer C.free(unsafe.Pointer(s1))
    defer C.free(unsafe.Pointer(s2))
    // s1, s2 := (*C.char)(unsafe.Pointer(&ori)), (*C.char)(unsafe.Pointer(&sign))

    ok := C.RsaVerify(pub, s1, s2)
    return ok == 1
}

func Rsa2Sign(pri unsafe.Pointer, ori string) string {
    s := C.CString(ori)
    defer C.free(unsafe.Pointer(s))
    // s := (*C.char)(unsafe.Pointer(&ori))

    sign := C.Rsa2Sign(pri, s)
    if sign == nil {
        return ""
    }

    // return C.GoBytes(unsafe.Pointer(sign), 344)
    return C.GoString(sign)
}

func Rsa2Verify(pub unsafe.Pointer, ori, sign string) bool {
    s1, s2 := C.CString(ori), C.CString(sign)
    defer C.free(unsafe.Pointer(s1))
    defer C.free(unsafe.Pointer(s2))
    // s1, s2 := (*C.char)(unsafe.Pointer(&ori)), (*C.char)(unsafe.Pointer(&sign))

    ok := C.Rsa2Verify(pub, s1, s2)
    return ok == 1
}

func DisposeKey(pri unsafe.Pointer) {
    if pri != nil {
        C.DisposeKey(pri)
    }
}