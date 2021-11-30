package config

import "time"

// 计算月份的规则

const (
	MonBeginDay int = 11 // 每月开始的第一天 11 号
	MonEndDay   int = 10 // 每月结束，下个月 10 号
)

// GetMon 获取当前时间的年份、月份
func GetMon(curTime time.Time) (year, mon int) {
	nilTime := time.Time{}
	if curTime == nilTime {
		curTime = time.Now()
	}
	year = curTime.Year()
	month := curTime.Month()
	day := curTime.Day()
	if day >= MonBeginDay {
		mon = int(month)
	} else {
		mon = int(month) - 1
	}
	return
}

// GetMonRange 获取月份的范围
func GetMonRange(year, mon int) (beginAt, endAt time.Time) {
	b := time.Date(year, time.Month(mon), MonBeginDay, 00, 00, 01, 00, time.Local)
    nextY, nextM := getNextMon(year, mon)
    e := time.Date(nextY, time.Month(nextM), MonEndDay, 23, 59, 59, 00, time.Local)
    return b, e
}

func getNextMon(year, mon int) (nextY, nextM int) {
	if mon == 12 {
		return year + 1, 1
	}
	return year, mon + 1
}
