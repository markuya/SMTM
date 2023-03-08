package QuotesAnalyzer

import (
	"SMTM/app/services/BaseType"
	"SMTM/app/services/HisQuotes"
	"testing"
	"time"

	"io"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func getxxxx(tStock *BaseType.Stock, tStartTime time.Time, tEndTime time.Time) (data []float32) {
	tDaySecond := 86400 * time.Second
	tFileTime := tStartTime
	for tEndTime.After(tFileTime) {
		tFileDayStr := tFileTime.Format("20060102")
		tFileTime = tFileTime.Add(tDaySecond)

		// 基础的涨跌标志
		if tBaseHisInfo, ok := tStock.His[tFileDayStr]; ok && tBaseHisInfo != nil {
			data = append(data, tBaseHisInfo.Change)
		}
	}

	return
}

func generateLineData(data []float32) []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < len(data); i++ {
		items = append(items, opts.LineData{Value: data[i]})
	}
	return items
}

func lineMulti(code1 string, data1 []float32, code2 string, data2 []float32) *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "multi lines",
		}),
		charts.WithInitializationOpts(opts.Initialization{
			Theme: "shine",
		}),
	)

	x := make([]string, 0)
	line.SetXAxis(x).
		AddSeries(code1, generateLineData(data1)).
		AddSeries(code2, generateLineData(data2))
	return line
}

func lineExamples(code1 string, data1 []float32, code2 string, data2 []float32) {
	page := components.NewPage()
	page.AddCharts(
		lineMulti(code1, data1, code2, data2),
	)
	f, err := os.Create("./line.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}

// 测试总入口
func Test(t *testing.T) {

	tStockMap, _ := HisQuotes.GetExchangeHisQuotes(`E:/test/SMTM/datafiles/his/China`, "2020-11-11", "2022-11-11",
		`E:/test/SMTM/datafiles/China`, "2022-11-12", "2023-02-23")

	// tStockMap2, _ := HisQuotes.GetExchangeHisQuotes(`E:\test\SMTM\datafiles\his\USA`, "2020-11-11", "2022-11-11",
	// 	`E:\test\SMTM\datafiles\USA`, "2022-11-12", "2023-02-23")

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

		tStrMap := bothRiseAndFall3(tStockMap, tStockMap, tTime1, tTime2)

		t.Logf("len(tStockMap)=%d len(tStrMap)=%d\n", len(tStockMap), len(tStrMap))
		var tMaxNum int = 0
		var tMaxKey1 string
		var tMaxInfo *brfData
		for k, v := range tStrMap {
			t.Logf("k=%s dayNum=%d code=%s\n", k, v.sortCodes[0].dayNum, v.sortCodes[0].code)
			if tMaxNum < v.sortCodes[0].dayNum {
				tMaxNum = v.sortCodes[0].dayNum
				tMaxKey1 = k
				tMaxInfo = v.sortCodes[0]
			}
		}

		t.Logf("@@@@@ k=%s code=%#v\n", tMaxKey1, tMaxInfo)

		tData1 := getxxxx(tStockMap[tMaxKey1], tTime1, tTime2)
		tData2 := getxxxx(tStockMap[tMaxInfo.code], tTime1.Add(-7*86400*time.Second), tTime2.Add(-7*86400*time.Second))
		lineExamples(tMaxKey1, tData1, tMaxInfo.code, tData2)
	}

}
