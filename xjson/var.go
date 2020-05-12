package xjson

// 特定场景下JSON使用的空值
var (
    EmptyMap   = make(map[int]int) // 空对象
    EmptySlice = make([]int, 0)    // 空切片
)