package utils

func Max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func Min(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func Diff(x, y int) int {
    v := x - y
    if v <= 0 {
        // v is negative or 0
        return 0
    }
    // v is positive or NaN
    return v
}
