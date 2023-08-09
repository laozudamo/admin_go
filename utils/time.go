package utils

import (
	"fmt"
	"time"
)

const (
	LocalDateTimeFormat string = "2006-01-02 15:04:05"
)

type LocalTime time.Time

func (l *LocalTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*l = LocalTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func (l LocalTime) MarshalText() (text []byte, err error) {
	b := make([]byte, 0, len(LocalDateTimeFormat))
	//b = append(b, '"')
	b = time.Time(l).AppendFormat(b, LocalDateTimeFormat)
	//b = append(b, '"')
	if string(b) == `0001-01-01 00:00:00` {
		b = []byte(``)
	}
	return b, nil
}
