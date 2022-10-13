package primCsv

import (
	"testing"
)

// 测试总入口
func Test(t *testing.T) {
	tReader, err := Open(`C:\Users\007\Downloads/dddddddddddddddddd.csv`, Options{HaveHeader: true})
	if nil != err {
		t.Logf("open err %s\n", err.Error())
		return
	}
	defer tReader.Close()
	for i := 0; i < 2; i++ {
		csvdata, err := tReader.ReadArr()
		if nil != err {
			break
		}
		t.Logf("FieldsPerRecord=%d data=%#v \n", tReader.mReader.FieldsPerRecord, csvdata)
	}

	tHeader, _ := tReader.GetHeader()
	t.Logf("FieldsPerRecord=%d tHeader=%#v \n", tReader.mReader.FieldsPerRecord, tHeader)
	for i := 2; i < 10; i++ {
		csvdata, err := tReader.ReadDict()
		if nil != err {
			break
		}
		t.Logf("FieldsPerRecord=%d data=%#v \n", tReader.mReader.FieldsPerRecord, csvdata)
	}
}
