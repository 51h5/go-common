package utils

// slice中字符串检索
func InSliceString(v string, s []string) bool {
    for _, vv := range s {
        if vv == v {
            return true
        }
    }
    return false
}

// slice检索
func InSlice(v interface{}, s []interface{}) bool {
    for _, vv := range s {
        if vv == v {
            return true
        }
    }
    return false
}

// 返回slice差集
func SliceDiff(s1, s2 []interface{}) (diff []interface{}) {
    for _, v := range s1 {
        if !InSlice(v, s2) {
            diff = append(diff, v)
        }
    }
    return
}

// 返回slice交集
func SliceXor(s1, s2 []interface{}) (xor []interface{}) {
    for _, v := range s1 {
        if InSlice(v, s2) {
            xor = append(xor, v)
        }
    }
    return
}