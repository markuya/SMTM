package HisQuotes

import (
	"SMTM/lib/filesFinder"
	"SMTM/lib/primCsv"
	"io"
)

// 加载指定目录下的所有历史行情数据
func loadHisData(tHisData string) map[string]*Stock {
	tStockMap := make(map[string]*Stock)

	// 遍历指定目录
	filesFinder.List(tHisData, `csv`,
		func(tFileName string, arg ...interface{}) {
			// 读取文件
			tReader, err1 := primCsv.Open(tFileName, primCsv.Options{HaveHeader: true})
			if nil != err1 {
				return
			}
			defer tReader.Close()

			// 逐行读取数据
			for {
				// read
				tData, err := tReader.ReadArr()

				// 文件结束
				if io.EOF == err {
					break
				}

				// 实例创建
				tStock := &Stock{
					Code: tData[0],
				}
				tStockMap[tFileName] = tStock
			}
		})

	return tStockMap
}

// 获取China交易所所有票的历史行情
// tHisData string - 历史数据存储路径
// tDayData string - 天历史数据存储路径
func loadHisQuotes(tHisData string, tDayData string) map[string]*Stock {
	return nil
}
