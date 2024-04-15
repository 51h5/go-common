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
	TIME_FORMAT_YM  = "%d%d"
)

func Today() string {
	return time.Now().Format(TIME_LAYOUT_YMD)
}

func TodayInt() int {
	v, _ := strconv.Atoi(Today())
	return v
}

func TodayWithTimestamp() (timestamp int64, today string) {
	now := time.Now()
	return now.UnixMilli(), now.Format(TIME_LAYOUT_YMD)
}

// 获取2个时间戳之间的间隔天数
// 类似浮点数向上取整, 计算2个时间戳分别对应的日期之间的天数差值 (计算精度到 天)
func SubDays(startTimestamp, endTimestamp int64) int64 {
	if endTimestamp <= 0 || endTimestamp < startTimestamp {
		return 0
	}

	start := time.UnixMilli(startTimestamp)
	end := time.UnixMilli(endTimestamp)

	// 转换到零点
	startDate := time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, time.Local)
	endDate := time.Date(end.Year(), end.Month(), end.Day(), 0, 0, 0, 0, time.Local)

	// 计算间隔天数
	return int64(endDate.Sub(startDate).Hours() / 24)
}

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
	return y*100 + w
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
		days[total-i-1] = t.Add(time.Duration(-1*i*24) * time.Hour).Format(format)
		// days = append(days, t.Add(time.Duration(-1 * i * 24) * time.Hour).Format(format))
	}

	return days
}
