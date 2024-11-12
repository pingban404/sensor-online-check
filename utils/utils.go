package utils

import (
    "time"
    "math"
    // "fmt"
)

// ConvertTimestampToLocalTime 将时间戳转换为本地时间，并返回转换后的时间和错误
func ConvertTimestampToLocalTime(timestamp string) (time.Time, error) {
    parsedTime, err := time.Parse(time.RFC3339, timestamp)
    if err != nil {
        return time.Time{}, err
    }
    return parsedTime.Local(), nil
}

// GetCurrentTime 获取当前本地时间
func GetCurrentTime() time.Time {
    currentTime := time.Now()
    // fmt.Println("当前时间为：", currentTime)
    return currentTime
}

// IsTimeDifferenceGreaterThanFiveMinutes 检查时间差是否大于 5 分钟
func IsTimeDifferenceGreaterThanFiveMinutes(diff time.Duration) bool {
    return math.Abs(float64(diff)) > float64(5*time.Minute)
}
