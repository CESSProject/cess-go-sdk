package utils

import (
	"time"
)

func GetUnixTimeAtZeroToday() int64 {
	return time.Now().UTC().Truncate(24 * time.Hour).Unix()
}

func DateTimeToUnix(t string) (int64, error) {
	ti, err := time.Parse(time.DateTime, t)
	if err != nil {
		return 0, err
	}
	return ti.Unix(), nil
}
