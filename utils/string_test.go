package utils

import "testing"

func BenchmarkLengthOfNonRepeatingSubStr(b *testing.B) {
    s := "尔发度发大姐夫问发个发题"

    // 增加长度
    for i := 0; i < 13; i++ {
        s = s + s
    }

    b.Logf("len(s) = %d", len(s))
    ans := 6
    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        actual := LengthOfNonRepeatingSubStr(s)
        if actual != ans {
            b.Errorf("got %d for input %s; expected %d\n", actual, "", ans)
        }
    }
}