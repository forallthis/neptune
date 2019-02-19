package timeutils

import (
	"strconv"
	"time"
)

const (
	// TZ timezone
	TZ = "Asia/Shanghai"
	// TimeFormart 格式时间字符串
	TimeFormart = "2006-01-02 15:04:05"
	// DataFormat 格式日期字符串
	DataFormat = "2006-01-02"
)

// TimeUnix 时间字符串转时间戳
func TimeUnix(data string) int64 {
	local, _ := time.LoadLocation(TZ)
	t, _ := time.ParseInLocation(TimeFormart, string(data), local)

	return t.Unix()
}

// Now 当前时间字符串
func Now() string {
	local, _ := time.LoadLocation(TZ)
	return time.Now().In(local).Format(TimeFormart)
}

// Today 当前日期字符串
func Today() string {
	local, _ := time.LoadLocation(TZ)
	return time.Now().In(local).Format(DataFormat)
}

// YearMonthStr 年月字符串,月份自动补零
func YearMonthStr() (ym string) {
	month := int(time.Now().Month())
	year := strconv.Itoa(time.Now().Year())
	if month >= 10 {
		ym = year + strconv.Itoa(month)
	} else {
		ym = year + "0" + strconv.Itoa(month)
	}
	return
}
