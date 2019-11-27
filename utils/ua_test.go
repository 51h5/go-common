package utils

import "testing"

func TestIsTmall(t *testing.T) {
    ua := "Mozilla/5.0 (iPhone; CPU iPhone OS 9_3_2 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Mobile/13F69 AliApp(TM/5.19.1) T-UA=iPhone_5.19.1_1242x2208_201200 WindVane/8.0.0 TMIOS/201200@tmall_iphone_5.19.1 1242x2208"
    if !IsTmall(ua) {
        t.Fatalf("天猫UA检测失败 (标准)")
    }

    ua = "Mobile/13F69 aliapp(tm/5.19.1) T-UA=iPhone_5.19.1_1242x2208_201200"
    if !IsTmall(ua) {
        t.Fatalf("天猫UA检测失败 (小写)")
    }

    ua = "Mobile/13F69 ALIAPP(TM/5.19.1) T-UA=iPhone_5.19.1_1242x2208_201200"
    if !IsTmall(ua) {
        t.Fatalf("天猫UA检测失败 (大写)")
    }
}