package utils

import (
    crand "crypto/rand"
    "encoding/binary"
    "math/rand"
)

func CryptoRand() uint32 {
    var n uint32
    binary.Read(crand.Reader, binary.LittleEndian, &n)
    return n
}

// RateToSeed 概率转换
//
// 百分比 => 概率种子
// {rate: 概率} / 100 = {x: 概率种子} / {sum: 种子基数}
func RateToSeed(rate, sum uint32) uint32 {
    rate *= sum
    if rate < 100 {
        return 0
    }
    return rate / 100
}

// DrawWithRate 概率计算
//
// rates => map[奖品标示]中奖概率
func DrawWithRate(rates map[uint8]uint32, sum uint32) uint8 {
    if len(rates) == 0 || sum == 0 {
        return 0
    }

    var pool []uint8

    var used uint32

    // 填充奖项
    for flag, rate := range rates {
        var i uint32
        for ; i < rate; i++ {
            pool = append(pool, flag)
            // pool[used + i] = flag
        }
        used += rate
    }

    // 填充无奖项
    for i := sum - used; i > 0; i-- {
        // pool = append(pool, make([]uint8, sum-used)...)
        pool = append(pool, 0)
    }

    for i := len(pool) - 1; i > 0; i-- {
        j := rand.Intn(i + 1)
        pool[i], pool[j] = pool[j], pool[i]
    }

    // flagSuccess := 0
    // flagFail := 0
    // for _, v := range pool {
    //     if v == 0 {
    //         flagFail++
    //     } else {
    //         flagSuccess++
    //     }
    // }
    // fmt.Printf("抽奖序列: len=%d, flagSuccess=%d, flagFail=%d\n", len(pool), flagSuccess, flagFail)
    // fmt.Printf("%v\n", pool)

    return pool[0]
}

// WeightRand 权重随机抽奖
//
// [ 权重1, 权重2, 权重3, ... ] =>  命中索引
func WeightRand(weights []int) int {
    var weightSum int
    for _, w := range weights {
        weightSum += w
    }

    if weightSum <= 0 {
        return -1
    }

    weight := rand.Intn(weightSum)

    var start, end int
    for i, w := range weights {
        if w < 0 {
            continue
        }

        end += w

        if w > 0 && start <= weight && weight < end {
            return i
        }

        start = end
    }

    return -1
}