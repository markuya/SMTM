package HisQuotes

import (
	"SMTM/lib/filesFinder"
	"SMTM/lib/primCsv"
	"io"
	"path"
	"strings"
	"time"
)

// 加载指定目录下的所有历史行情数据
func loadHisData(tHisData string, tStartTime time.Time, tEndTime time.Time) map[string]*Stock {
	tStockMap := make(map[string]*Stock)

	tStartTimeInt64 := tStartTime.Unix()
	tEndTimeInt64 := tEndTime.Unix()

	// 遍历指定目录
	filesFinder.List(tHisData, `csv`,
		func(tFile string, arg ...interface{}) {
			// 读取文件
			tReader, err1 := primCsv.Open(tFile, primCsv.Options{HaveHeader: true})
			if nil != err1 {
				return
			}
			defer tReader.Close()

			// 实例创建
			tStock := &Stock{
				His: make(map[string]*Quote),
			}

			// 提取文件名
			tFileName := path.Base(strings.Replace(tFile, `\`, `/`, -1))

			// 提取股票名
			tStock.Code = strings.Replace(strings.Replace(tFileName, `code-`, ``, -1), `.csv`, ``, -1)
			// 逐行读取数据
			for {
				// read
				tData, err := tReader.ReadDict()

				// 文件结束
				if io.EOF == err {
					break
				}

				tTime, _ := _checkTimeFormat(tData[`Date`])
				if tTime == nil {
					continue
				}

				if (tTime.Unix() < tStartTimeInt64) || (tTime.Unix() > tEndTimeInt64) {
					continue
				}

				// 日行情历史
				tHisQuote := &Quote{
					AdjClose: _string2Float32(tData[`AdjClose`]), // 调整收盘价格
					Close:    _string2Float32(tData[`Close`]),    // 收盘价格
					High:     _string2Float32(tData[`High`]),     // 最高
					Low:      _string2Float32(tData[`Low`]),      // 最低
					Open:     _string2Float32(tData[`Open`]),     // 开盘价格
					Volume:   _string2Int64(tData[`Volume`]),     // 成交量
					Amount:   _string2Float64(tData[`Amount`]),   // 成交额
				}
				tHisQuote.Date = tTime

				tStock.His[tData[`Date`]] = tHisQuote
			}

			// 按名 记录
			tStockMap[tStock.Code] = tStock
		})

	//
	//
	return tStockMap
}

// 加载指定目录下的所有天历史数据
func loadDayData(tDayData string, tStartTime time.Time, tEndTime time.Time, tStockMap map[string]*Stock) {

	tDaySecond := 86400 * time.Second
	tFileTime := tStartTime

	for tEndTime.After(tFileTime) {
		tFileDayStr := tFileTime.Format("20060102")
		tFileTime = tFileTime.Add(tDaySecond)

		tFile := tDayData + "\\quotes" + tFileDayStr + ".csv"

		// 读取文件
		tReader, err1 := primCsv.Open(tFile, primCsv.Options{HaveHeader: true})
		if nil != err1 {
			return
		}
		defer tReader.Close()

		// 逐行读取数据
		for {
			// read
			tData, err := tReader.ReadDict()

			// 文件结束
			if io.EOF == err {
				break
			}

			tCode := tData[`Code`]
			if tStock, ok := tStockMap[tCode]; ok && tStock != nil {
				// 日行情历史
				tHisQuote := &Quote{
					AdjClose: _string2Float32(tData[`AdjClose`]), // 调整收盘价格
					Close:    _string2Float32(tData[`Close`]),    // 收盘价格
					High:     _string2Float32(tData[`High`]),     // 最高
					Low:      _string2Float32(tData[`Low`]),      // 最低
					Open:     _string2Float32(tData[`Open`]),     // 开盘价格
					Volume:   _string2Int64(tData[`Volume`]),     // 成交量
					Amount:   _string2Float64(tData[`Amount`]),   // 成交额
				}
				tHisQuote.Date, _ = _checkTimeFormat(tFileDayStr)

				tStock.His[tFileDayStr] = tHisQuote
			}

		}
	}
}

// 获取China交易所所有票的历史行情
// tHisData string - 历史数据存储路径
// tDayData string - 天历史数据存储路径
// tStartTime time.Time - 开始时间
// tEndTime time.Time - 结束时间
func loadHisQuotes(tHisData string, tDayData string, tStartTime time.Time, tEndTime time.Time) map[string]*Stock {

	return nil
}
