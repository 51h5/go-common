package utils

import (
	"testing"

    "github.com/stretchr/testify/assert"
)

func TestCryptoRand(t *testing.T) {
    if CryptoRand() <= 0 {
        t.Failed()
    }
}

func TestRateToSeed(t *testing.T) {
    var seed uint32
    var rate uint32
    var sum uint32

    rate = 30
    sum = 10000
    seed = RateToSeed(rate, sum)
    if seed != 3000 {
        t.Errorf("种子计算错误: rate=%d, sum=%d, seed=%d", rate, sum, seed)
        t.FailNow()
    }

    rate = 50
    sum = 10000
    seed = RateToSeed(rate, sum)
    if seed != 5000 {
        t.Errorf("种子计算错误: rate=%d, sum=%d, seed=%d", rate, sum, seed)
        t.FailNow()
    }

    rate = 90
    sum = 100000
    seed = RateToSeed(rate, sum)
    if seed != 90000 {
        t.Errorf("种子计算错误: rate=%d, sum=%d, seed=%d", rate, sum, seed)
        t.FailNow()
    }

    rate = 100
    sum = 10000
    seed = RateToSeed(rate, sum)
    if seed != sum {
        t.Errorf("种子计算错误: rate=%d, sum=%d, seed=%d", rate, sum, seed)
        t.FailNow()
    }
}

func TestDrawWithRate(t *testing.T) {
    var rate1 uint32 = 0
    var sum1 uint32 = 10000
    var flag1 uint8 = 1
    rates1 := map[uint8]uint32{
        flag1: RateToSeed(rate1, sum1),
    }
    r1 := DrawWithRate(rates1, sum1)
    if r1 == flag1 {
        t.Errorf("非预期中奖: rate=%d, sum=%d, result=%d", rate1, sum1, r1)
        t.FailNow()
    }

    var rate2 uint32 = 100
    var sum2 uint32 = 10000
    var flag2 uint8 = 1
    rates2 := map[uint8]uint32{
        flag2: RateToSeed(rate2, sum2),
    }
    r2 := DrawWithRate(rates2, sum2)
    if r2 != flag2 {
        t.Errorf("非预期未中奖: rate=%d, sum=%d, result=%d", rate2, sum2, r2)
        t.FailNow()
    }
}

func TestWeightRand(t *testing.T) {
    var weightsEmpty []int

    assert.Equal(t, -1, WeightRand(weightsEmpty), "空权重")

    assert.Equal(t, -1, WeightRand([]int{0, 0, 0, 0}), "全部0权重")

    assert.Equal(t, 0, WeightRand([]int{1000}), "仅一个权重且非0")

    assert.Equal(t, 1, WeightRand([]int{0, 10000}), "仅一个权重且仅最后一个非0")
    assert.Equal(t, 3, WeightRand([]int{0, 0, 0, 10000}), "多个权重且仅最后一个非0")

    assert.Equal(t, 0, WeightRand([]int{100, 0, 0, 0}), "多个权重且仅第一个非0")
    assert.Equal(t, 1, WeightRand([]int{0, 1000, 0, 0}), "多个权重且仅第二个非0")

    assert.NotEqual(t, -1, WeightRand([]int{100, 200, 400, 1000}), "多个权重且全部非0")
}

func BenchmarkCryptoRand(b *testing.B) {
    for i := 0; i < b.N; i++ {
        CryptoRand()
    }
}

func BenchmarkDrawWithRate(b *testing.B) {
    rates := map[uint8]uint32{
        1: RateToSeed(10, 100000),
    }

    for i := 0; i < b.N; i ++ {
        _ = DrawWithRate(rates, 100000)
    }
}

func BenchmarkWeightRand(b *testing.B) {
    weights := []int{100, 500, 1200, 3000, 6000}
    for i := 0; i < b.N; i++ {
        WeightRand(weights)
    }
}