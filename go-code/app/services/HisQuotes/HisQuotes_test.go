package HisQuotes

import (
	"testing"
	"time"
)

// 测试总入口
func Test(t *testing.T) {

	ttt, _ := GetExchangeHisQuotes(`E:\test\SMTM\datafiles\his\China`, "2020-11-11", "2022-11-11",
		`E:\test\SMTM\datafiles\China`, "2022-11-12", "2023-02-23")

	t.Logf("%#v\n", ttt["000001"].His)
	t.Logf("%#v\n", ttt["000001"].His["20221110"])
	t.Logf("%#v\n", ttt["000001"].His["20221111"])
	t.Logf("%#v\n", ttt["000001"].His["20221112"])
	t.Logf("%#v\n", ttt["000001"].His["20221113"])
	t.Logf("%#v\n", ttt["000001"].His["20221114"])
	t.Logf("%#v\n", ttt["000001"].His["20221115"])

	// tAllNum := 0
	// t590Num := 0
	// for k, v := range ttt {
	// 	t.Logf("k=%s len(v)=%d\n", k, len(v.His))
	// 	tAllNum++
	// 	if len(v.His) == 590 {
	// 		t590Num++
	// 	}
	// }

	// t.Logf("tAllNum=%d t590Num=%d\n", tAllNum, t590Num)

	// for k, v := range ttt {
	// 	t.Logf("k=%s \n", k)
	// 	for k2, v2 := range v.His {
	// 		t.Logf("k2=%s \n", k2)
	// 		t.Logf("%s v2=%#v\n", v2.Date.String(), v2)
	// 	}
	// }
	tTime1 := time.Now()
	tTime2 := time.Now().Add(86400 * time.Second)
	t.Logf("%s\n", tTime1.Format("20060102"))
	t.Logf("%s\n", tTime2.Format("20060102"))

	if tTime2.After(tTime1) {
		t.Logf("tTime2 is big \n")
	}
}
