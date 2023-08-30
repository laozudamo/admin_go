package utils

import (
	"fmt"
	"time"
)

func GetNowFormatTodayTime() string {

	now := time.Now()
	dateStr := fmt.Sprintf("%02d-%02d-%02d", now.Year(), int(now.Month()),
		now.Day())

	return dateStr
}

func ParseTime(strTime string) (*time.Time, bool) {
	layout := "2006-01-02 15:04:05"
	parsedTime, err := time.Parse(layout, strTime)
	if err != nil {
		return nil, false
	}
	return &parsedTime, true
}
