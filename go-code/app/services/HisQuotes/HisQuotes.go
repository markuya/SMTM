package HisQuotes

import (
	"errors"
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
