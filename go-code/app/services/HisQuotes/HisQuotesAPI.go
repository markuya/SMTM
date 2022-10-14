package HisQuotes

import (
	"strings"
)

// 获取指定交易所所有票的历史行情
// tName string - 交易市场标识
// tHisData string - 历史数据存储路径
// tDayData string - 天历史数据存储路径
// tStartTime string - 开始时间(格式:xxxx-xx-xx 例如:1992-03-21)
// tEndTime string - 结束时间(格式:xxxx-xx-xx 例如:1992-03-21)
func GetExchangeHisQuotes(tName string, tHisData string, tDayData string, tStartTime string, tEndTime string) (map[string]*Stock, error) {
	tUpperName := strings.ToUpper(tName)
	switch tUpperName {
	case "CHINA":
		return loadHisQuotesChina(tHisData, tDayData)
		break
	case "USA":
		return loadHisQuotesUSA(tHisData, tDayData)
		break
	default:
		break
	}

	// 时间格式验证
	tTime1, err1 := _checkTimeFormat(tStartTime)
	if nil != err1 {
		return nil, err1
	}
	tTime2, err2 := _checkTimeFormat(tStartTime)
	if nil != err2 {
		return nil, err2
	}

	// 加载历史行情
	return loadHisQuotes(tHisData, tDayData, tTime1, tTime2)
}
