package QuotesAnalyzer

import (
	"SMTM/app/services/BaseType"
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
