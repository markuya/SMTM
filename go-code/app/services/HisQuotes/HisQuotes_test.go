package HisQuotes

import (
	"testing"
	"time"
)

// 测试总入口
func Test(t *testing.T) {

	// ttt := loadHisData(`E:\test\SMTM\go-code\app\services\HisQuotes\testdata\his\China`, nil, nil)

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
