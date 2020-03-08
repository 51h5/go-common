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

func TestIsDingTalk(t *testing.T) {
    ua := "Mozilla/5.0 (iPhone; CPU iPhone OS 13_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 AliApp(DingTalk/4.7.22) com.laiwang.DingTalk/12410922 Channel/201200 language/zh-Hans UT4Aplus/0.0.6 WK"
    if !IsDingTalk(ua) {
        t.Fatalf("钉钉UA检测失败 (标准)")
    }

    ua = "Mobile/5.0 aliapp(dingtalk/4.7.22) com.laiwang.DingTalk/12410922"
    if !IsDingTalk(ua) {
        t.Fatalf("钉钉UA检测失败 (小写)")
    }

    ua = "Mobile/13F69 ALIAPP(DINGTALK/4.7.22) com.laiwang.DingTalk/12410922"
    if !IsDingTalk(ua) {
        t.Fatalf("钉钉UA检测失败 (大写)")
    }
}