package StocksPloter

import (
	"SMTM/app/services/BaseType"
	"time"
)

// 按指定需求进行图标绘制

func DrawKline(tStocks []*BaseType.Stock, tStartTime time.Time, tEndTime time.Time, tFile string) {
	klinePage(tStocks, tStartTime, tEndTime, tFile)
}
