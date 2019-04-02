package xjson

type XjsonSuccess struct {
    Status uint8       `json:"status"`
    Data   interface{} `json:"data,omitempty"`
}

type XjsonError struct {
    Status    uint8       `json:"status"`
    Error     string      `json:"error"`
    ErrorCode int         `json:"error_code"`
    Data      interface{} `json:"data,omitempty"`
}

// 特定场景下JSON使用的空值
var (
    EmptyMap   = make(map[int]int) // 空对象
    EmptySlice = make([]int, 0)    // 空切片
)
