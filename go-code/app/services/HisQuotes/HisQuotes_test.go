package HisQuotes

import (
	"SMTM/lib/filesFinder"
	"testing"
)

// 测试总入口
func Test(t *testing.T) {
	var tList []string
	filesFinder.List(`E:/test/SMTM/datafiles/his/China`, `csv`,
		func(tFileName string, arg ...interface{}) {
			tList = append(tList, tFileName)
		})

	for _, tFileName := range tList {
		t.Logf("%s\n", tFileName)
	}
}
