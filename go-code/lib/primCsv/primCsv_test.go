package primCsv

import (
	"strconv"
	"testing"
)

func test_read(t *testing.T) {
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

func test_write(t *testing.T) {
	var header []string
	header = append(header, "id")
	header = append(header, "num")
	var tRecords [][]string
	for i := 0; i < 500000; i++ {
		var t []string
		t = append(t, strconv.Itoa(i))
		t = append(t, strconv.Itoa(i*2))
		tRecords = append(tRecords, t)
	}

	// 创建writer
	{
		_, err := WriteAll("./WriteAll20221027.csv", tRecords, WriterOptions{
			//Header:   header,
			MaxLines: 500000,
		})
		if err != nil {
			t.Logf("WriteAll err=%s\n", err.Error())
		}
	}
	// 创建writer
	tWriter, err := NewWriter("./test20221027.csv", WriterOptions{
		Header:   header,
		MaxLines: 500000,
	})

	if err != nil {
		t.Logf("NewWriter err=%s\n", err.Error())
	}

	if tWriter != nil {
		if tWriter.mFileHandle == nil {
			t.Logf("tWriter mFileHandle is nil\n")
		}
		{
			err := tWriter.Write(header)
			if err != nil {
				t.Logf("tWriter err=%s\n", err.Error())
			}
		}
		{ // 写入所有数据
			err := tWriter.WriteAll(tRecords)
			if err != nil {
				t.Logf("tWriter err=%s\n", err.Error())
			}
		}

		// 关闭
		tWriter.Close()
	}

}

// 测试总入口
func Test(t *testing.T) {
	test_read(t)
	test_write(t)
}
