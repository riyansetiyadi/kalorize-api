package utils

import "time"

func GetExpiredTime() time.Time {
	now := time.Now()
	expiredTime := now.Add(3 * 24 * time.Hour)
	return expiredTime
}
