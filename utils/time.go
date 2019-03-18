package utils

import (
    "fmt"
    "time"
)

const (
    TIME_LAYOUT_YMD = "20060102"
    TIME_LAYOUT_YM  = "200601"
    TIME_LAYOUT_Y   = "2006"
    TIME_FORMAT_YM = "%s%s"
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