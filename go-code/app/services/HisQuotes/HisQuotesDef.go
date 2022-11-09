package HisQuotes

import (
	"time"
)

// 行情信息
type Quote struct {
	Date     *time.Time // 日期
	AdjClose float32    // 调整收盘价格
	Close    float32    // 收盘价格
	High     float32    // 最高
	Low      float32    // 最低
	Open     float32    // 开盘价格
	Volume   int64      // 成交量
	Amount   float64    // 成交额
}

type Stock struct {
	Code string            // 代码
	Name string            // 名字
	His  map[string]*Quote // 历史行情(按天如 2022-01-01)
}
