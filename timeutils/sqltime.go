package timeutils

import (
	"fmt"
	"time"
)

// SQLTime sql自定义时间格式
type SQLTime time.Time

// MarshalJSON sql进行scan时调用此方法
func (t SQLTime) MarshalJSON() ([]byte, error) {
	local, _ := time.LoadLocation(TZ)
	var stamp = fmt.Sprintf("\"%s\"", time.Time(t).In(local).Format(TimeFormart))
	return []byte(stamp), nil
}

// String 转为字符串
func (t SQLTime) String() string {
	local, _ := time.LoadLocation(TZ)
	return time.Time(t).In(local).Format(TimeFormart)
}

// SQLDate sql自定义日期格式
type SQLDate time.Time

// MarshalJSON sql进行scan时调用此方法
func (t SQLDate) MarshalJSON() ([]byte, error) {
	local, _ := time.LoadLocation(TZ)
	var stamp = fmt.Sprintf("\"%s\"", time.Time(t).In(local).Format(DataFormat))
	return []byte(stamp), nil
}

// String 转为字符串
func (t SQLDate) String() string {
	local, _ := time.LoadLocation(TZ)
	return time.Time(t).In(local).Format(DataFormat)
}
