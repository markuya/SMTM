package QuotesAnalyzer

import (
	"SMTM/app/services/HisQuotes"
	"testing"
	"time"
)

// 测试总入口
func Test(t *testing.T) {

	tStockMap, _ := HisQuotes.GetExchangeHisQuotes(`E:\test\SMTM\datafiles\his\China`, "2020-11-11", "2022-11-11",
		`E:\test\SMTM\datafiles\China`, "2022-11-12", "2023-02-23")

	t.Logf("%#v\n", tStockMap["000001"].His)
	t.Logf("%#v\n", tStockMap["000001"].His["20221110"])
	t.Logf("%#v\n", tStockMap["000001"].His["20221111"])
	t.Logf("%#v\n", tStockMap["000001"].His["20221112"])
	t.Logf("%#v\n", tStockMap["000001"].His["20221113"])
	t.Logf("%#v\n", tStockMap["000001"].His["20221114"])
	t.Logf("%#v\n", tStockMap["000001"].His["20221115"])

	// 读取历史数据
	{
		// 时间格式验证
		tTime1, err1 := time.Parse("2006-01-02", "2021-11-11")
		if nil != err1 {
			return
		}
		tTime2, err2 := time.Parse("2006-01-02", "2022-11-11")
		if nil != err2 {
			return
		}

		tStrMap := bothRiseAndFall(tStockMap, tTime1, tTime2)

		for k, v := range tStrMap {
			if len(v) > 1 {
				t.Logf("k=%s len=%d\n", k, len(v))
			}
		}
	}
}
