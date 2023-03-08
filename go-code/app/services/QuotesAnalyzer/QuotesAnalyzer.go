package QuotesAnalyzer

import (
	"SMTM/app/services/BaseType"
	"sort"
	"sync"
	"time"
)

// 指定时间段的计算
func countRiseFallString(tStock *BaseType.Stock, tStartTime time.Time, tEndTime time.Time) (tStr string) {
	tDaySecond := 86400 * time.Second
	tFileTime := tStartTime
	for tEndTime.After(tFileTime) {
		tFileDayStr := tFileTime.Format("20060102")
		tFileTime = tFileTime.Add(tDaySecond)

		if tHisInfo, ok := tStock.His[tFileDayStr]; ok && tHisInfo != nil {
			if tHisInfo.Close > tHisInfo.Open {
				tStr += "+"
			} else if tHisInfo.Close < tHisInfo.Open {
				tStr += "-"
			} else {
				tStr += "="
			}
		} else {
			tStr += " "
		}
	}

	return
}

// Both rise and fall
func bothRiseAndFall(tStockMap map[string]*BaseType.Stock, tStartTime time.Time, tEndTime time.Time) map[string][]string {
	var tWaitGroup sync.WaitGroup
	var tWriteMutex sync.Mutex

	tRecordMap := make(map[string][]string)

	for tCode, tStock := range tStockMap {
		tWaitGroup.Add(1)

		go func() {
			tStr := countRiseFallString(tStock, tStartTime, tEndTime)
			tWriteMutex.Lock()
			if tArr, ok := tRecordMap[tStr]; ok && tArr != nil {
				tArr = append(tArr, tCode)
			} else {
				var tArr []string
				tArr = append(tArr, tCode)
				tRecordMap[tStr] = tArr
			}
			tWriteMutex.Unlock()
			tWaitGroup.Done()
		}()
	}

	tWaitGroup.Wait()

	return tRecordMap
}

// 指定时间段的计算
func countRiseFallString2(tBaseCode string, tBaseStock *BaseType.Stock, tStockMap map[string]*BaseType.Stock, tStartTime time.Time, tEndTime time.Time) (tStockBrfData *stockBrfData) {
	tDaySecond := 86400 * time.Second
	tFileTime := tStartTime
	tLen := len(tStockMap)

	if _, ok := tStockMap[tBaseCode]; ok {
		tLen--
	}

	tStockBrfData = &stockBrfData{
		code:      tBaseCode,
		data:      make(map[string]*brfData),
		sortCodes: make(brfDataSlice, tLen),
	}

	tLen = 0
	for tCode, _ := range tStockMap {
		if tBaseCode != tCode {
			tBrfData := &brfData{
				code: tCode,
			}
			tStockBrfData.data[tCode] = tBrfData
			tStockBrfData.sortCodes[tLen] = tBrfData
			tLen++
		}
	}

	var tBaseRfFlag int8 = 0

	for tEndTime.After(tFileTime) {
		tFileDayStr := tFileTime.Format("20060102")
		tFileTime = tFileTime.Add(tDaySecond)

		// 基础的涨跌标志
		if tBaseHisInfo, ok := tBaseStock.His[tFileDayStr]; ok && tBaseHisInfo != nil {
			if tBaseHisInfo.Close > tBaseHisInfo.Open {
				tBaseRfFlag = 3
			} else if tBaseHisInfo.Close < tBaseHisInfo.Open {
				tBaseRfFlag = 2
			} else {
				tBaseRfFlag = 1
			}
		} else {
			continue
		}

		for tCode, tStock := range tStockMap {
			if tBaseCode != tCode {
				var tRfFlag int8 = 0
				if tHisInfo, ok := tStock.His[tFileDayStr]; ok && tHisInfo != nil {
					if tHisInfo.Close > tHisInfo.Open {
						tRfFlag = 3
					} else if tHisInfo.Close < tHisInfo.Open {
						tRfFlag = 2
					} else {
						tRfFlag = 1
					}
				}
				if tRfFlag == tBaseRfFlag {
					tStockBrfData.data[tCode].dayNum++
					tStockBrfData.data[tCode].continuumDays++
					if tStockBrfData.data[tCode].continuumDays > tStockBrfData.data[tCode].continuumDaysMax {
						tStockBrfData.data[tCode].continuumDaysMax = tStockBrfData.data[tCode].continuumDays
					}
				} else {
					tStockBrfData.data[tCode].continuumDays = 0
				}
			}
		}

	}

	sort.Stable(tStockBrfData.sortCodes)
	return
}

// Both rise and fall
func bothRiseAndFall2(tStockMap1 map[string]*BaseType.Stock, tStockMap2 map[string]*BaseType.Stock, tStartTime time.Time, tEndTime time.Time) map[string]*stockBrfData {
	var tWaitGroup sync.WaitGroup
	var tWriteMutex sync.Mutex

	tRecordMap := make(map[string]*stockBrfData)

	for tCode, tStock := range tStockMap1 {
		tWaitGroup.Add(1)
		var tBaseCode string
		var tBaseStock *BaseType.Stock
		tBaseCode = tCode
		tBaseStock = tStock
		go func() {

			tStockBrfData := countRiseFallString2(tBaseCode, tBaseStock, tStockMap2, tStartTime, tEndTime)
			tWriteMutex.Lock()

			tRecordMap[tBaseCode] = tStockBrfData

			tWriteMutex.Unlock()
			tWaitGroup.Done()
		}()
	}

	tWaitGroup.Wait()

	return tRecordMap
}

// 指定时间段的计算
func countRiseFallString3(tBaseCode string, tBaseStock *BaseType.Stock, tStockMap map[string]*BaseType.Stock, tStartTime time.Time, tEndTime time.Time) (tStockBrfData *stockBrfData) {
	tDaySecond := 86400 * time.Second
	tFileTime := tStartTime
	tLen := len(tStockMap)
	if _, ok := tStockMap[tBaseCode]; ok {
		tLen--
	}

	tStockBrfData = &stockBrfData{
		code:      tBaseCode,
		data:      make(map[string]*brfData),
		sortCodes: make(brfDataSlice, tLen),
	}

	tLen = 0
	for tCode, _ := range tStockMap {
		if tBaseCode != tCode {
			tBrfData := &brfData{
				code: tCode,
			}
			tStockBrfData.data[tCode] = tBrfData
			tStockBrfData.sortCodes[tLen] = tBrfData
			tLen++
		}
	}

	var tBaseRfFlag int8 = 0

	for tEndTime.After(tFileTime) {
		tFileDayStr := tFileTime.Format("20060102")
		tFileLastDayStr := tFileTime.Add(-7 * tDaySecond).Format("20060102")
		tFileTime = tFileTime.Add(tDaySecond)

		// 基础的涨跌标志
		if tBaseHisInfo, ok := tBaseStock.His[tFileDayStr]; ok && tBaseHisInfo != nil {
			if tBaseHisInfo.Change > 0 {
				tBaseRfFlag = 3
			} else if tBaseHisInfo.Change < 0 {
				tBaseRfFlag = 2
			} else {
				tBaseRfFlag = 1
			}
		} else {
			continue
		}

		for tCode, tStock := range tStockMap {
			if tBaseCode != tCode {
				var tRfFlag int8 = 0
				if tHisInfo, ok := tStock.His[tFileLastDayStr]; ok && tHisInfo != nil {
					if tHisInfo.Change > 0 {
						tRfFlag = 3
					} else if tHisInfo.Change < 0 {
						tRfFlag = 2
					} else {
						tRfFlag = 1
					}
				}
				if tRfFlag == tBaseRfFlag {
					tStockBrfData.data[tCode].dayNum++
					tStockBrfData.data[tCode].continuumDays++
					if tStockBrfData.data[tCode].continuumDays > tStockBrfData.data[tCode].continuumDaysMax {
						tStockBrfData.data[tCode].continuumDaysMax = tStockBrfData.data[tCode].continuumDays
					}
				} else {
					tStockBrfData.data[tCode].continuumDays = 0
				}
			}
		}

	}

	sort.Stable(tStockBrfData.sortCodes)
	return
}

// Both rise and fall
func bothRiseAndFall3(tStockMap1 map[string]*BaseType.Stock, tStockMap2 map[string]*BaseType.Stock, tStartTime time.Time, tEndTime time.Time) map[string]*stockBrfData {
	var tWaitGroup sync.WaitGroup
	var tWriteMutex sync.Mutex

	tRecordMap := make(map[string]*stockBrfData)

	for tCode, tStock := range tStockMap1 {
		tWaitGroup.Add(1)
		var tBaseCode string
		var tBaseStock *BaseType.Stock
		tBaseCode = tCode
		tBaseStock = tStock
		go func() {

			tStockBrfData := countRiseFallString3(tBaseCode, tBaseStock, tStockMap2, tStartTime, tEndTime)
			tWriteMutex.Lock()

			tRecordMap[tBaseCode] = tStockBrfData

			tWriteMutex.Unlock()
			tWaitGroup.Done()
		}()
	}

	tWaitGroup.Wait()

	return tRecordMap
}
