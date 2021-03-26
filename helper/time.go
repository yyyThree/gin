package helper

import (
	"strings"
	"time"
)

func FormatDateNow() string {
	return time.Now().Format("2006-01-02")
}

func FormatDateNowBySlash() string {
	return time.Now().Format("2006/01/02")
}

func FormatDateTimeNow() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func FormatDateTime(t string) time.Time {
	res, _ := time.Parse("2006-01-02 15:04:05", t)
	return res
}

func FormatDateTimeHour(t string) string {
	newTime, _ := time.Parse("2006-01-02 15:04:05", t)
	return newTime.Format("15:04:05")
}

func FormatDateTimeInterval() string {
	return time.Now().Format("15:04")
}

func FormatDateWeekday() time.Weekday {
	return time.Now().Weekday()
}

func FormatDateWeekdayWithDateTime(datetime string) time.Weekday {
	t, _ := time.Parse("2006-01-02", datetime)
	return t.Weekday()
}

// 补充 10:00 至 2006-01-02 15:04:05
func FormatDateSupply(t string) string {
	return time.Now().Format("2006-01-02") + " " + t + ":00"
}

func FormatDateTimeByTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func FormatDateTimeZero() string {
	return "0000-00-00 00:00:00"
}

// 基于给定的时间(如：10:00:00)，返回当前时间的标准time格式
func FormatByTime(t string) time.Time {
	if t == "" {
		return time.Now()
	}

	res, _ := time.Parse("2006-01-02 15:04:05", strings.Join([]string{FormatDateNow(), t}, " "))
	return res
}
