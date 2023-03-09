package StocksPloter

import (
	"SMTM/app/services/BaseType"
	"path"
	"time"

	"io"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

// 按指定需求进行图标绘制
func klineDataZoomBoth(tStock *BaseType.Stock, tStartTime time.Time, tEndTime time.Time) *charts.Kline {

	tDaySecond := 86400 * time.Second
	tFileTime := tStartTime

	kline := charts.NewKLine()

	x := make([]string, 0)
	y := make([]opts.KlineData, 0)
	for tEndTime.After(tFileTime) {
		tFileDayStr := tFileTime.Format("20060102")

		// 基础的涨跌标志
		if tBaseHisInfo, ok := tStock.His[tFileDayStr]; ok && tBaseHisInfo != nil {
			x = append(x, tFileTime.Format("2006/01/02"))
			y = append(y, opts.KlineData{Value: [4]float32{tBaseHisInfo.Open, tBaseHisInfo.Close, tBaseHisInfo.High, tBaseHisInfo.Low}})
		}

		tFileTime = tFileTime.Add(tDaySecond)
	}

	kline.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: tStock.Code,
		}),
		charts.WithXAxisOpts(opts.XAxis{
			SplitNumber: 20,
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Scale: true,
		}),
		charts.WithDataZoomOpts(opts.DataZoom{
			Type:       "inside",
			Start:      50,
			End:        100,
			XAxisIndex: []int{0},
		}),
		charts.WithDataZoomOpts(opts.DataZoom{
			Type:       "slider",
			Start:      50,
			End:        100,
			XAxisIndex: []int{0},
		}),
	)

	kline.SetXAxis(x).AddSeries(tStock.Code, y)
	return kline
}

func klinePage(tStocks []*BaseType.Stock, tStartTime time.Time, tEndTime time.Time, tFile string) {
	page := components.NewPage()
	page.PageTitle = "SMTM"

	for _, tStock := range tStocks {
		page.AddCharts(
			klineDataZoomBoth(tStock, tStartTime, tEndTime),
		)
	}

	tPath := path.Dir(tFile)
	if _, err := os.Stat(tPath); os.IsNotExist(err) {
		err := os.MkdirAll(tPath, 0755)
		if err != nil {
			panic(err)
		}
	}

	f, err := os.Create(tFile)
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
