package QuotesAnalyzer

//
type brfData struct {
	code             string
	dayNum           int // 同涨同跌总天数
	continuumDays    int // 连续同涨同跌总天数
	continuumDaysMax int // 连续同涨同跌最大天数
}

//brf信息 slice类型
type brfDataSlice []*brfData

// slice类型 排序接口
func (tSlice brfDataSlice) Len() int      { return len(tSlice) }
func (tSlice brfDataSlice) Swap(i, j int) { tSlice[i], tSlice[j] = tSlice[j], tSlice[i] }
func (tSlice brfDataSlice) Less(i, j int) bool {
	if tSlice[i].dayNum < tSlice[j].dayNum {
		return false
	} else if tSlice[i].dayNum == tSlice[j].dayNum {
		if tSlice[i].continuumDaysMax > tSlice[j].continuumDaysMax {
			return false
		} else if tSlice[i].continuumDaysMax == tSlice[j].continuumDaysMax {
			return false
		}
	}
	return true
}

//
type stockBrfData struct {
	code      string              // 股票编码
	data      map[string]*brfData // 对齐其他股票的同涨同跌信息
	sortCodes brfDataSlice
}
