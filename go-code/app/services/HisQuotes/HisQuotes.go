package HisQuotes

import (
	"errors"
	"strconv"
	"time"
)

// 检测时间格式
func _checkTimeFormat(tTime string) (*time.Time, error) {
	// 检测是否为 yyyy-mm-dd 格式
	{
		tTime, err := time.Parse("2006-01-02", tTime)
		if nil == err {
			return &tTime, nil
		}
	}

	// 检测是否为 yyyymmdd 格式
	{
		tTime, err := time.Parse("20060102", tTime)
		if nil == err {
			return &tTime, nil
		}
	}

	// 检测是否为 yyyy/mm/dd 格式
	{
		tTime, err := time.Parse("2006/01/02", tTime)
		if nil == err {
			return &tTime, nil
		}
	}

	return nil, errors.New("parsing time as yyyy-mm-dd or yyyy/mm/dd or yyyymmdd failed!")
}

// 字符串转int64
func _string2Int64(s string) int64 {
	tValue, _ := strconv.ParseInt(s, 10, 64)
	return tValue
}

// 字符串转float64
func _string2Float64(s string) float64 {
	tValue, _ := strconv.ParseFloat(s, 64)
	return tValue
}

// 字符串转float32
func _string2Float32(s string) float32 {
	tValue, _ := strconv.ParseFloat(s, 64)
	return float32(tValue)
}
