package HisQuotes

import (
	"strings"
)

// 获取指定交易所所有票的历史行情
// tName string - 交易市场标识
// tHisData string - 历史数据存储路径
// tDayData string - 天历史数据存储路径
func GetExchangeHisQuotes(tName string, tHisData string, tDayData string) map[string]*Stock {
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

	return nil
}
