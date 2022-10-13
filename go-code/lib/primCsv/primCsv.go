package primCsv

import (
	"encoding/csv"
	"errors"
	"os"
)

// Options is a configuration that is used to create a new CsvReader.
type Options struct {
	// CSV文件是否拥有标题
	// 如果HaveHeader=true,打开文件时读取第一行数据,并记录数据到标题表
	HaveHeader bool

	// FieldsPerRecord is the number of expected fields per record.
	// If FieldsPerRecord is positive, Read requires each record to
	// have the given number of fields. If FieldsPerRecord is 0, Read sets it to
	// the number of fields in the first record, so that future records must
	// have the same field count. If FieldsPerRecord is negative, no check is
	// made and records may have a variable number of fields.
	FieldsPerRecord int

	// If LazyQuotes is true, a quote may appear in an unquoted field and a
	// non-doubled quote may appear in a quoted field.
	LazyQuotes bool

	// If TrimLeadingSpace is true, leading white space in a field is ignored.
	// This is done even if the field delimiter, Comma, is white space.
	TrimLeadingSpace bool

	// ReuseRecord controls whether calls to Read may return a slice sharing
	// the backing array of the previous call's returned slice for performance.
	// By default, each call to Read returns newly allocated memory owned by the caller.
	ReuseRecord bool
}

// csv文件读取器
type CsvReader struct {
	mFile        string         // 文件
	mReader      *csv.Reader    // 数据读取句柄
	mFileHandle  *os.File       // 文件句柄
	mHeader      []string       // 标题列表
	mHeaderIndex map[string]int // 标题索引(每个标题对应 mHeader []string 中的索引)
	mHaveHeader  bool           // 是否有标题
}

// 关闭文件读取器
func (tReader *CsvReader) Close() {
	if nil != tReader.mFileHandle {
		tReader.mFileHandle.Close()
	}
}

// 逐行读取文件,并以数组格式返回数据
func (tReader *CsvReader) ReadArr() ([]string, error) {
	if (nil != tReader.mReader) && (nil != tReader.mFileHandle) {
		tStrList, err := tReader.mReader.Read()
		if nil != err {
			return nil, err
		}
		// 检查行数据是否合法
		return tStrList, nil
	}

	return nil, errors.New("handle is nil!")
}

// 逐行读取文件,并以字典表格式返回数据
func (tReader *CsvReader) ReadDict() (map[string]string, error) {
	if (nil != tReader.mReader) && (nil != tReader.mFileHandle) {
		// 未加载或设置标题信息，需加载标题
		if false == tReader.mHaveHeader {
			return nil, errors.New("the csv file have no header!")
		}

		// 读取一行数据
		tStrList, err := tReader.mReader.Read()
		if nil != err {
			return nil, err
		}

		// 将行数据转化为字典表
		if nil == tReader.mHeaderIndex {
			return nil, errors.New("Header indexs dict is nil!")
		}

		//
		tStrMap := make(map[string]string)
		tStrLen := len(tStrList)

		//
		for tTitle, tIndex := range tReader.mHeaderIndex {
			if tIndex >= tStrLen {
				return nil, errors.New("this data num is less then the title!")
			} else {
				tStrMap[tTitle] = tStrList[tIndex]
			}
		}
		//
		return tStrMap, nil
	}

	return nil, errors.New("handle is nil!")
}

// 读取文件标题
func (tReader *CsvReader) GetHeader() ([]string, error) {
	if (nil != tReader.mReader) && (nil != tReader.mFileHandle) {
		// 已缓存标题信息，直接返回
		if tReader.mHaveHeader {
			return tReader.mHeader, nil
		}

		return nil, errors.New("no is header!")
	}

	return nil, errors.New("handle is nil!")
}

// 读取文件标题
func (tReader *CsvReader) loadHeader() error {
	if (nil != tReader.mReader) && (nil != tReader.mFileHandle) {
		// 已缓存标题信息，直接返回
		if tReader.mHaveHeader {
			return nil
		}

		// 读取标题信息
		tHeader, err := tReader.mReader.Read()

		// 不管是否加载成功都设置已加载，防止反复加载标题
		tReader.mHaveHeader = true

		// 读取验证
		if nil != err {
			return err
		}

		// 记录文件头信息
		tReader.mHeader = tHeader // 标题列表
		tReader.mHeaderIndex = make(map[string]int)
		for i, tTitle := range tHeader {
			tReader.mHeaderIndex[tTitle] = i // 标题索引(每个标题对应 mHeader []string 中的索引)
		}

		return nil
	}

	return errors.New("handle is nil!")
}

// 创建新句柄
func Open(tFileName string, opts ...Options) (*CsvReader, error) {
	// 打开文件
	tFileHandle, err1 := os.Open(tFileName)

	// 打开文件错误检测
	if nil != err1 {
		return nil, err1
	}

	// 创建文件读取器
	tReader := &CsvReader{
		mFile:       tFileName,
		mReader:     csv.NewReader(tFileHandle),
		mFileHandle: tFileHandle,
		mHaveHeader: false,
	}

	var tOptions *Options
	// 未传入配置信息
	if len(opts) == 0 {
		tOptions = &Options{
			HaveHeader:      false,
			FieldsPerRecord: -1,
		}
	} else {
		tOptions = &opts[0]
	}

	// 判断文件是否有标题
	if true == tOptions.HaveHeader {
		tReader.loadHeader()
	}

	// 配置设置
	tReader.mReader.FieldsPerRecord = tOptions.FieldsPerRecord
	tReader.mReader.LazyQuotes = tOptions.LazyQuotes
	tReader.mReader.ReuseRecord = tOptions.ReuseRecord
	tReader.mReader.TrimLeadingSpace = tOptions.TrimLeadingSpace

	return tReader, nil
}
