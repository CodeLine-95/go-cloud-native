package time

import (
	"time"
)

func ParseTime(timeStr string) (int64, error) {
	curTime, err := time.ParseInLocation(time.DateTime, timeStr, new(time.Location))
	if err != nil {
		return 0, err
	}
	return curTime.Unix(), nil
}

func FormatTime(timeNum int64) (string, error) {
	return FormatTimeWithLayout(timeNum, time.DateTime)
}

func FormatTimeWithLayout(timeNum int64, layout string) (string, error) {
	return time.Unix(timeNum, 0).In(new(time.Location)).Format(layout), nil
}
