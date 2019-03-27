package utils

// 查询无重复字符的最长字符片段
// Given a string, find the length of the longest substring without repeating characters.
func LengthOfNonRepeatingSubStr(s string) int {
    lastOccurred := make(map[rune]int)
    start := 0
    maxLength := 0

    for i, ch := range []rune(s) {
        if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
            start = lastI + 1
        }
        if i-start+1 > maxLength {
            maxLength = i - start + 1
        }
        lastOccurred[ch] = i
    }

    return maxLength
}

// 查询无重复字符的最长字符片段v2
func LengthOfNonRepeatingSubStrV2(s string) int {
    lastOccurred := make([]int, 0xffff)

    // 赋给一个初始值
    for i := range lastOccurred {
        lastOccurred[i] = -1
    }
    start := 0
    maxLength := 0

    for i, ch := range []rune(s) {
        if lastI := lastOccurred[ch]; lastI != -1 && lastI >= start {
            start = lastI + 1
        }
        if i-start+1 > maxLength {
            maxLength = i - start + 1
        }
        lastOccurred[ch] = i
    }

    return maxLength
}
