package utils

import (
    "testing"
    "time"
)

func TestYearWeek(t *testing.T) {
    now := time.Now()

    t1 := time.Date(2019, 3, 17, 0, 0, 0, 0, now.Location())
    v1 := "201911"
    w1 := YearWeek(t1)
    if w1 != v1 {
        t.Errorf("YearWeek 返回数值错误: %s != %s", w1, v1)
        return
    }

    t2 := time.Date(2019, 3, 19, 0, 0, 0, 0, now.Location())
    v2 := "201912"
    w2 := YearWeek(t2)
    if w2 != v2 {
        t.Errorf("YearWeek 返回数值错误: %s != %s", w2, v2)
        return
    }
}

func TestYearWeekInt(t *testing.T) {
    now := time.Now()

    t1 := time.Date(2019, 3, 17, 0, 0, 0, 0, now.Location())
    v1 := 201911
    w1 := YearWeekInt(t1, )
    if w1 != v1 {
        t.Errorf("YearWeek 返回数值错误: %d != %d", w1, v1)
        return
    }

    t2 := time.Date(2019, 3, 19, 0, 0, 0, 0, now.Location())
    v2 := 201912
    w2 := YearWeekInt(t2)
    if w2 != v2 {
        t.Errorf("YearWeek 返回数值错误: %d != %d", w2, v2)
        return
    }
}

func TestYearWeekf(t *testing.T) {
    now := time.Now()
    format := "%d(%d)"

    t1 := time.Date(2019, 3, 17, 0, 0, 0, 0, now.Location())
    v1 := "2019(11)"
    w1 := YearWeekf(t1, format)
    if w1 != v1 {
        t.Errorf("YearWeek 返回数值错误: %s != %s", w1, v1)
        return
    }

    t2 := time.Date(2019, 3, 19, 0, 0, 0, 0, now.Location())
    v2 := "2019(12)"
    w2 := YearWeekf(t2, format)
    if w2 != v2 {
        t.Errorf("YearWeek 返回数值错误: %s != %s", w2, v2)
        return
    }
}

func TestWeekDays(t *testing.T) {
    now := time.Now()

    t1 := time.Date(2019, 3, 17, 0, 0, 0, 0, now.Location())

    days := WeekDays(t1)

    if len(days) != 7 {
        t.Errorf("WeekDays 返回长度错误: %d != %d", len(days), 7)
        return
    }

    if cap(days) != 7 {
        t.Errorf("WeekDays 返回容量错误: %d != %d", cap(days), 7)
        return
    }

    ds := []string{"20190311", "20190312", "20190313", "20190314", "20190315", "20190316", "20190317"}
    for i, v := range ds {
        if v != days[i] {
            t.Errorf("WeekDays 返回数值错误: index=%d, ds=%v, day=%v", i, v, days[i])
            return
        }
    }
}

func TestWeekDays2(t *testing.T) {
    now := time.Now()

    t1 := time.Date(2019, 3, 19, 0, 0, 0, 0, now.Location())

    days := WeekDays(t1)

    if len(days) != 2 {
        t.Errorf("WeekDays 返回长度错误: %d != %d", len(days), 2)
        return
    }

    if cap(days) != 2 {
        t.Errorf("WeekDays 返回容量错误: %d != %d", cap(days), 2)
        return
    }

    ds := []string{"20190318", "20190319"}
    for i, v := range ds {
        if v != days[i] {
            t.Errorf("WeekDays 返回数值错误: index=%d, ds=%v, day=%v", i, v, days[i])
            return
        }
    }
}

func TestWeekDaysInt(t *testing.T) {
    now := time.Now()

    t1 := time.Date(2019, 3, 17, 0, 0, 0, 0, now.Location())

    days := WeekDaysInt(t1)

    if len(days) != 7 {
        t.Errorf("WeekDays 返回长度错误: %d != %d", len(days), 7)
        return
    }

    if cap(days) != 7 {
        t.Errorf("WeekDays 返回容量错误: %d != %d", cap(days), 7)
        return
    }

    ds := []int{20190311, 20190312, 20190313, 20190314, 20190315, 20190316, 20190317}
    for i, v := range ds {
        if v != days[i] {
            t.Errorf("WeekDays 返回数值错误: index=%d, ds=%v, day=%v", i, v, days[i])
            return
        }
    }
}

func TestWeekDaysInt2(t *testing.T) {
    now := time.Now()

    t1 := time.Date(2019, 3, 19, 0, 0, 0, 0, now.Location())

    days := WeekDaysInt(t1)

    if len(days) != 2 {
        t.Errorf("WeekDays 返回长度错误: %d != %d", len(days), 2)
        return
    }

    if cap(days) != 2 {
        t.Errorf("WeekDays 返回容量错误: %d != %d", cap(days), 2)
        return
    }

    ds := []int{20190318, 20190319}
    for i, v := range ds {
        if v != days[i] {
            t.Errorf("WeekDays 返回数值错误: index=%d, ds=%v, day=%v", i, v, days[i])
            return
        }
    }
}

func TestWeekDaysf(t *testing.T) {
    now := time.Now()
    format := "2006-01-02"

    t1 := time.Date(2019, 3, 17, 0, 0, 0, 0, now.Location())

    days := WeekDaysf(t1, format)

    if len(days) != 7 {
        t.Errorf("WeekDays 返回长度错误: %d != %d", len(days), 7)
        return
    }

    if cap(days) != 7 {
        t.Errorf("WeekDays 返回容量错误: %d != %d", cap(days), 7)
        return
    }

    ds := []string{"2019-03-11", "2019-03-12", "2019-03-13", "2019-03-14", "2019-03-15", "2019-03-16", "2019-03-17"}
    for i, v := range ds {
        if v != days[i] {
            t.Errorf("WeekDays 返回数值错误: index=%d, ds=%v, day=%v", i, v, days[i])
            return
        }
    }
}

func TestWeekDaysf2(t *testing.T) {
    now := time.Now()
    format := "2006-01-02"

    t1 := time.Date(2019, 3, 19, 0, 0, 0, 0, now.Location())

    days := WeekDaysf(t1, format)

    if len(days) != 2 {
        t.Errorf("WeekDays 返回长度错误: %d != %d", len(days), 2)
        return
    }

    if cap(days) != 2 {
        t.Errorf("WeekDays 返回容量错误: %d != %d", cap(days), 2)
        return
    }

    ds := []string{"2019-03-18", "2019-03-19"}
    for i, v := range ds {
        if v != days[i] {
            t.Errorf("WeekDays 返回数值错误: index=%d, ds=%v, day=%v", i, v, days[i])
            return
        }
    }
}
