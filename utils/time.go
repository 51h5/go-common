package utils

import (
    "fmt"
    "strconv"
    "time"
)

const (
    TIME_LAYOUT_YMD = "20060102"
    TIME_LAYOUT_YM  = "200601"
    TIME_LAYOUT_Y   = "2006"
    TIME_FORMAT_YM = "%d%d"
)

func YearWeek(t time.Time) string {
    return YearWeekf(t, TIME_FORMAT_YM)
}

func YearWeekf(t time.Time, format string) string {
    y, w := t.ISOWeek()

    if format == "" {
        return fmt.Sprintf(TIME_FORMAT_YM, y, w)
    }

    return fmt.Sprintf(format, y, w)
}

func YearWeekInt(t time.Time) int {
    y, w := t.ISOWeek()
    return y * 100 + w
}

// 获取指定时间所在周的每日时间(layout: "20060102")
func WeekDays(t time.Time) []string {
    return WeekDaysf(t, TIME_LAYOUT_YMD)
}

// 获取指定时间所在周的每日时间
func WeekDaysInt(t time.Time) []int {
    vs := WeekDaysf(t, TIME_LAYOUT_YMD)

    days := make([]int, len(vs))
    for i, v := range vs {
        days[i], _ = strconv.Atoi(v)
    }

    return days
}

// 获取指定时间所在周的每日时间(含当天)
func WeekDaysf(t time.Time, format string) []string {
    // 今天所在周周几(0 ~ 6)
    total := int(t.Weekday())

    // 若是周日, 则需要修正下, 返回整周的日期
    if total == 0 {
        total = 7
    }

    // 优化容量
    days := make([]string, total)

    // fmt.Printf("<WeekDays> %d, len: %d, cap: %d\n", total, len(days), cap(days))

    for i := total - 1; i >= 0; i-- {
        // fmt.Printf("<WeekDays.for> %d, len: %d, cap: %d\n", i, len(days), cap(days))
        days[total - i - 1] = t.Add(time.Duration(-1 * i * 24) * time.Hour).Format(format)
        // days = append(days, t.Add(time.Duration(-1 * i * 24) * time.Hour).Format(format))
    }

    return days
}
