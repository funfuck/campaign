package mytime

import "time"

const (
	YmdHis = "2006-01-02 15:04:05"
)

func GetTime(format string, date string) *time.Time {
	loc, _ := time.LoadLocation("Asia/Bangkok")
	t, _ := time.ParseInLocation(format, date, loc)
	return &t
}
