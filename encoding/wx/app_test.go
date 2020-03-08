package wx

import "testing"

func TestApp_Code2Session(t *testing.T) {
    a := App{
        appId:  "wx969c86a28f8b0fe7",
        secret: "4a15a41f750e46f3bb2a7fcde846ebd1",
    }

    sess, err := a.Code2Session("0a1zwYZT08TIv22RRD0U0aTZZT0zwYZw")
    if err != nil {
        t.Fatalf("获取令牌失败: %s\n", err.Error())
    }

    t.Logf("获取令牌成功\n")
    t.Logf("- openid: %s\n", sess.OpenId)
    t.Logf("- unionid: %s\n", sess.UnionId)
    t.Logf("- session: %s\n", sess.SessionKey)
}

func TestApp_DecryptBizData(t *testing.T) {
    a := App{
        appId:      "wx969c86a28f8b0fe7",
        secret:     "4a15a41f750e46f3bb2a7fcde846ebd1",
        sessionKey: "ZHIUgzsx/eQLX99JZUl6cQ==",
    }

    bizData := "rJ4rbHxfZzreLiiB6U0IJ/RFQZ77zhdtL370SaOgSL8Hyk76egOuWxyUJAn4cgno4WlBKRolNk4vRnu85a+SGiZG8WLuEeE80ftzNQAnhA4NfcNgfCJ5x/9lAFPumEiHLxdCCzvW38lZfJ+GIdUG34cvvoUyyW72KIZG94ymE429lOyAUvMgJ36v9BpnV/Y1wVW7T40lHVr8znOmpptpPyd21t0TMxpq2n8VgAcsgqgY4QifxWu4y2VqVQ9EKRL9ZbRMxpfq4s0dK2qlytTEBORsPYcQKaHAcmi4080eKjNOWgKgmNWQNKnvzChexFTty+FtUxSVebArePocO2k1TkkfytCjsJpxmaEH8CSFTa02b2dWv6SGCCKP8dClvm9Y8WG8RLPtKAapV9V1oDmqyFs1mt5iK5hnL2O4IWV0LWB6R7yS/FegbkiCbe9B3qfGU8pnnvqbmlIpmR0ThqEOrfzlkCsEwfuZw2Ti0vKAnlBv1sUBkcG/XYUgp9LeNub1TC7JmRCxoWBAsI+mCcbdtKjvcTlLOGxkkdRPPiX5Hjs="
    iv := "Y+SqtI6RrzIS5MI7rz265g=="

    user, err := a.DecryptBizData(bizData, iv)
    if err != nil {
        t.Fatalf("解密失败: %s\n", err.Error())
    }

    t.Logf("解析成功\n")
    t.Logf("- OpenId: %s\n", user.OpenId)
    t.Logf("- UnionId: %s\n", user.UnionId)
    t.Logf("- Gender: %d\n", user.Gender)
    t.Logf("- NickName: %s\n", user.NickName)
    t.Logf("- AvatarUrl: %s\n", user.AvatarUrl)
    t.Logf("- Country: %s\n", user.Country)
    t.Logf("- Province: %s\n", user.Province)
    t.Logf("- City: %s\n", user.City)
    t.Logf("- Language: %s\n", user.Language)
    t.Logf("- Watermark.AppId: %s\n", user.Watermark.AppId)
    t.Logf("- Watermark.Timestamp: %d\n", user.Watermark.Timestamp)
}

func TestApp_DecryptBizData2(t *testing.T) {
    a := App{
        appId:  "wx969c86a28f8b0fe7",
        secret: "4a15a41f750e46f3bb2a7fcde846ebd1",
    }

    sess, err := a.Code2Session("071pKPTh2N8rtD0yg3Rh2lz5Uh2pKPTk")
    if err != nil {
        t.Fatalf("获取令牌失败: %s\n", err.Error())
    }

    t.Logf("获取令牌成功\n")
    t.Logf("- openid: %s\n", sess.OpenId)
    t.Logf("- unionid: %s\n", sess.UnionId)
    t.Logf("- session: %s\n", sess.SessionKey)

    bizData := "LPBhftV2AjqNepT9ctSkN2koGUf31Ksxd5gT8JiLenS/wKuI1bzWC6d3T001Glf3fDRY34Cxf7OoSL0PwVmzmkTanksB63GPK64BaVFFx6WgTbNrC1rGzDBNxE84zXzMOjlXASBbuUIZv7i4I/Nq/saDkZC95OxU18QyuwGdAIZOpfQHT4SOCuQdcflfOULG6a4O/PkhyBE7vFsSgcvjiqmVr2yXCdfH83jrujehp1/Bt7YuS4LyKLdFCFEYc+dDwCtgS8tVEEAkKJhuZcC9yXllgUSiJ6UYnT/E7M6wgi5in5Q7Oup/QXV2ctyJAil+AIwY6M1nTwHIBG7o0ZOW3Ih/11AlxkNizrtWTIOWdEc4qSpd2uQXHDcTdM+fCgclIuGycAbfBsctDW86ZI59V/ZPAJoRTQNhOzNtEmqamV2CngvCpDkAzobMfXwIvn8qgkD9nZ06n3Z7dtIi3/XNOUHA5cwSqUNPKQ0CTuMmUZ18xnQl+5/R+Kc/dQxFw1yO3zpQdU9eq01w32ZZA5a+7XslCNN10IP80b6xjk9mi4k="
    iv := "gb520zmUEoW3phHyOx9h2A=="

    user, err := a.DecryptBizData(bizData, iv)
    if err != nil {
        t.Fatalf("解密失败: %s\n", err.Error())
    }

    t.Logf("解密成功: %v\n", user)
    t.Logf("- OpenId: %s\n", user.OpenId)
    t.Logf("- UnionId: %s\n", user.UnionId)
    t.Logf("- Gender: %d\n", user.Gender)
    t.Logf("- NickName: %s\n", user.NickName)
    t.Logf("- AvatarUrl: %s\n", user.AvatarUrl)
    t.Logf("- Country: %s\n", user.Country)
    t.Logf("- Province: %s\n", user.Province)
    t.Logf("- City: %s\n", user.City)
    t.Logf("- Language: %s\n", user.Language)
    t.Logf("- Watermark.AppId: %s\n", user.Watermark.AppId)
    t.Logf("- Watermark.Timestamp: %d\n", user.Watermark.Timestamp)
}
