package crsa

import (
    "testing"
)

const (
    ori            = "aaa"
    oriSignRsa     = "IIjjNLvi7m1HgvjkM4mJ/gOQLQNflgW5pyT15PEpce2oBvBdlqtoZpGIuKmYgfo1kgO6kIinq07sjnqpdKL/5UEQQXrx+ZnO3TjDEmkdQp0BzPr8S/U9BGJtFfty4p7032jV403riqDS5hxeLflIzfBB3bsqmn44e0nuOsH5v9g="
    oriSignRsa2    = "Xt5ygodmn3G1JOh0mFN8E2SEgBO17mcxH/x8o4IzrWl/tMX9f9XTMgIzLMvUE9IWAzDmKDFp+6RFAEQqsJv7Jc8ssFWbLgPExg/8DgsMGJxNqm3fdPs+0yzT7iSaFXQT3ZliVgzUOViPTcTts4IDZVrkssOmhaZ2w+8F2QhoBqKEzz3jMkkq0ox9q76VeuiXvIYenHlGgJ4RPyM2Ey9rCDcQNkc0s0++ErGP+2Om6YykktDVdMfYLBVpJC+JlUY39WohH7ZOwtazVcFY8yJc4NiXXYmA8Ym3JW93tLcocnncSzKygNcYvVSxs05D3l048M8FOJzQt6GZ46FuQN3TFw=="
    pemPrivateRsa  = "./rsa_private.pem"
    pemPublicRsa   = "./rsa_public.pem"
    pemPrivateRsa2 = "./rsa2_private.pem"
    pemPublicRsa2  = "./rsa2_public.pem"
)

func TestRsa(t *testing.T) {
    pri := LoadPrivateKey(pemPrivateRsa)
    if pri == nil {
        t.Fatalf("rsa private load: %v", pri)
    }

    sign := RsaSign(pri, ori)
    DisposeKey(pri)

    if sign != oriSignRsa {
        t.Fatalf("rsa sign fail")
    }

    pub := LoadPublicKey(pemPublicRsa)
    ok := RsaVerify(pub, ori, sign)
    DisposeKey(pub)

    if !ok {
        t.Fatalf("rsa verify fail")
    }
}

func BenchmarkRsaSign(b *testing.B) {
    pri := LoadPrivateKey(pemPrivateRsa)

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = RsaSign(pri, ori)
    }

    DisposeKey(pri)
}

func BenchmarkRsaVerify(b *testing.B) {
    pub := LoadPublicKey(pemPublicRsa)

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        RsaVerify(pub, ori, oriSignRsa)
    }

    DisposeKey(pub)
}

func BenchmarkRsa(b *testing.B) {
    pri := LoadPrivateKey(pemPrivateRsa)
    pub := LoadPublicKey(pemPublicRsa)

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        RsaVerify(pub, ori, RsaSign(pri, ori))
    }

    DisposeKey(pri)
    DisposeKey(pub)
}

func TestRsa2Sign(t *testing.T) {
    pri := LoadPrivateKey(pemPrivateRsa2)
    if pri == nil {
        t.Fatalf("rsa2 private load: %v", pri)
    }

    sign := Rsa2Sign(pri, ori)
    DisposeKey(pri)

    if sign != oriSignRsa2 {
        t.Fatalf("rsa2 sign fail")
    }

    pub := LoadPublicKey(pemPrivateRsa2)
    ok := Rsa2Verify(pub, ori, sign)
    DisposeKey(pub)

    if !ok {
        t.Fatalf("rsa2 verfiy fail")
    }
}

func BenchmarkRsa2Sign(b *testing.B) {
    pri := LoadPrivateKey(pemPrivateRsa2)

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        Rsa2Sign(pri, ori)
    }

    DisposeKey(pri)
}

func BenchmarkRsa2Verify(b *testing.B) {
    pub := LoadPublicKey(pemPublicRsa2)

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        Rsa2Verify(pub, ori, oriSignRsa2)
    }

    DisposeKey(pub)
}

func BenchmarkRsa2(b *testing.B) {
    pri := LoadPrivateKey(pemPrivateRsa2)
    pub := LoadPublicKey(pemPublicRsa2)

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        Rsa2Verify(pub, ori, Rsa2Sign(pri, ori))
    }

    DisposeKey(pri)
    DisposeKey(pub)
}
