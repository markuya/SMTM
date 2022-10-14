package HisQuotes

// 行情信息
type Quote struct {
}

type Stock struct {
	Code string            // 代码
	Name string            // 名字
	His  map[string]*Quote // 历史行情(按天如 2022-01-01)
}
