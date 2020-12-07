// package timeformat Y-m-d H:i:s 
package timef

import (
  "time"
)

// TimeZone: Asia/Shanghai时间格式化
func Datef(timePlaceholder string, t int64) (formatTime string) {
	if timePlaceholder == "" {
		return ""
	}
	tu := time.Unix(t, 0)
	tz, _ := time.LoadLocation("Asia/Shanghai")
	tu.In(tz)
	switch timePlaceholder {
	case "Y-m-d H:i:s":
		formatTime = tu.Format("2006-01-02 15:04:05")
	case "Y-m-d":
		formatTime = tu.Format("2006-01-02")
	case "H:i:s":
		formatTime = tu.Format("15:04:05")
	}
	return
}
