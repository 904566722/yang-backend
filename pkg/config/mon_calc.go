package config

import (
	"time"
)

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
    nextY, nextM := GetNextMon(year, mon)
    e := time.Date(nextY, time.Month(nextM), MonEndDay, 23, 59, 59, 00, time.Local)
    return b, e
}

func GetNextMon(year, mon int) (nextY, nextM int) {
	if mon == 12 {
		return year + 1, 1
	}
	return year, mon + 1
}

func GetDayRange(year, mon, day int) (beginAt, endAt time.Time) {
	b := time.Date(year, time.Month(mon), day, 05, 00, 00, 00, time.Local)
	nextY, nextM, nextD := getNextDay(year, mon, day)
	e := time.Date(nextY, time.Month(nextM), nextD, 04, 59, 59, 00, time.Local)
	return b, e
}

var MonDay = []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func isRun(year int) bool {
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		return true
	}
	return false
}

func getNextDay(year, mon, day int) (nextY, nextM, nextD int) {
	if isRun(year) {
		MonDay[2] = 29
	}
	day = day + 1
	if day > MonDay[mon] {
		day = 1
		mon = mon + 1
		if mon > 12 {
			mon = 1
			year = year + 1
		}
	}
	return year, mon, day
}