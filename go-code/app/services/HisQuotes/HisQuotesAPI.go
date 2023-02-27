package HisQuotes

import (
	"SMTM/app/services/BaseType"
)

// 获取指定交易所所有票的历史行情
// tName string - 交易市场标识
// tHisData string - 历史数据存储路径
// tDayData string - 天历史数据存储路径
// tStartTime string - 开始时间(格式:xxxx-xx-xx 例如:1992-03-21)
// tEndTime string - 结束时间(格式:xxxx-xx-xx 例如:1992-03-21)
func GetExchangeHisQuotes(tHisData string, tHisSTime string, tHisETime string, tDayData string, tDaySTime string, tDayETime string) (map[string]*BaseType.Stock, error) {

	var tStockMap map[string]*BaseType.Stock

	// 读取历史数据
	{
		// 时间格式验证
		tTime1, err1 := _checkTimeFormat(tHisSTime)
		if nil != err1 {
			return nil, err1
		}
		tTime2, err2 := _checkTimeFormat(tHisETime)
		if nil != err2 {
			return nil, err2
		}

		tStockMap = loadHisData(tHisData, *tTime1, *tTime2)
	}

	// 读取天历史数据
	if tStockMap != nil {
		// 时间格式验证
		tTime1, err1 := _checkTimeFormat(tDaySTime)
		if nil != err1 {
			return nil, err1
		}
		tTime2, err2 := _checkTimeFormat(tDayETime)
		if nil != err2 {
			return nil, err2
		}

		loadDayData(tDayData, *tTime1, *tTime2, tStockMap)
	}

	// 加载历史行情
	return tStockMap, nil
}
