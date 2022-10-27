package primCsv

import (
	"encoding/csv"
	"errors"
	"os"

	"strconv"
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

	return nil, errors.New("handle is nil! ")
}

// 逐行读取文件,并以字典表格式返回数据
func (tReader *CsvReader) ReadDict() (map[string]string, error) {
	if (nil != tReader.mReader) && (nil != tReader.mFileHandle) {
		// 未加载或设置标题信息，需加载标题
		if !tReader.mHaveHeader {
			return nil, errors.New("the csv file have no header! ")
		}

		// 读取一行数据
		tStrList, err := tReader.mReader.Read()
		if nil != err {
			return nil, err
		}

		// 将行数据转化为字典表
		if tReader.mHeaderIndex == nil {
			return nil, errors.New(" Header indexs dict is nil! ")
		}

		//
		tStrMap := make(map[string]string)
		tStrLen := len(tStrList)

		//
		for tTitle, tIndex := range tReader.mHeaderIndex {
			if tIndex >= tStrLen {
				return nil, errors.New("this data num is less then the title! ")
			} else {
				tStrMap[tTitle] = tStrList[tIndex]
			}
		}
		//
		return tStrMap, nil
	}

	return nil, errors.New("handle is nil! ")
}

// 读取文件标题
func (tReader *CsvReader) GetHeader() ([]string, error) {
	if (nil != tReader.mReader) && (nil != tReader.mFileHandle) {
		// 已缓存标题信息，直接返回
		if tReader.mHaveHeader {
			return tReader.mHeader, nil
		}

		return nil, errors.New("no is header! ")
	}

	return nil, errors.New("handle is nil! ")
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

	return errors.New("handle is nil! ")
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
	if tOptions.HaveHeader {
		tReader.loadHeader()
	}

	// 配置设置
	tReader.mReader.FieldsPerRecord = tOptions.FieldsPerRecord
	tReader.mReader.LazyQuotes = tOptions.LazyQuotes
	tReader.mReader.ReuseRecord = tOptions.ReuseRecord
	tReader.mReader.TrimLeadingSpace = tOptions.TrimLeadingSpace

	return tReader, nil
}

// csv文件写入器
type CsvWriter struct {
	mFile       string      // 文件名
	mWriter     *csv.Writer // 数据写入句柄
	mFileHandle *os.File    // 当前文件句柄
	mHeader     []string    // 标题列表
	mMaxLines   int         // 单文件最大行
	mNowLines   int         // 当前文件行数
	mFilesNum   int         // 已写入文件数量
	mFileList   []string    // 文件列表
}

// Options is a configuration that is used to create a new CsvWriter.
type WriterOptions struct {
	// CSV文件标题
	Header []string

	// 单文件最大行.
	MaxLines int

	// Field delimiter (set to ',' by NewWriter)
	Comma rune

	// True to use \r\n as the line terminator
	UseCRLF bool
}

// 打开新的写入器
func NewWriter(tFileName string, opts ...WriterOptions) (*CsvWriter, error) {
	// 打开文件
	tFileHandle, err := os.OpenFile(tFileName, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}
	tWriter := &CsvWriter{
		mFile:       tFileName,
		mFileHandle: tFileHandle,
		mNowLines:   0,
		mFilesNum:   1,
		mWriter:     csv.NewWriter(tFileHandle),
	}
	var tOptions *WriterOptions
	// 未传入配置信息
	if len(opts) == 0 {
		tOptions = &WriterOptions{
			Comma:   ',',
			UseCRLF: true,
		}
	} else {
		tOptions = &opts[0]
		if tOptions.Comma == 0 {
			tOptions.Comma = ','
			tOptions.UseCRLF = true
		}
	}

	tWriter.mMaxLines = tOptions.MaxLines
	tWriter.mHeader = tOptions.Header
	//tWriter.mWriter.Comma = tOptions.Comma
	//tWriter.mWriter.UseCRLF = tOptions.UseCRLF

	// 存储文件添加到文件列表
	tWriter.mFileList = append(tWriter.mFileList, tFileName)

	return tWriter, nil
}

// 写入记录
func (tWriter *CsvWriter) Write(tRecords []string) error {
	// 句柄检测
	if (tWriter.mFileHandle == nil) && (tWriter.mWriter == nil) {
		return errors.New("CsvWriter write failed! the file/writer handle is nil")
	}

	// 判断是否需要创建新文件
	if (tWriter.mMaxLines > 0) && (tWriter.mNowLines >= tWriter.mMaxLines) {
		// 文件写入缓存推送
		tWriter.mWriter.Flush()
		// 文件数+1
		tWriter.mFilesNum++
		// 当前写入行数重置
		tWriter.mNowLines = 0
		// 新打开文件
		tFileName := tWriter.mFile[:len(tWriter.mFile)-4] + "-" + strconv.Itoa(tWriter.mFilesNum) + tWriter.mFile[len(tWriter.mFile)-4:]
		tFileHandle, err := os.OpenFile(tFileName, os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			return err
		}
		// 关闭原文件
		tWriter.mFileHandle.Close()
		// 新句柄切换
		tWriter.mFileHandle = tFileHandle
		tWriter.mWriter = csv.NewWriter(tFileHandle)

		// 如果有文件头
		if tWriter.mHeader != nil {
			err := tWriter.mWriter.Write(tWriter.mHeader)
			if err != nil {
				tWriter.mFileHandle.Close()
				return err
			}
		}

		// 存储文件添加到文件列表
		tWriter.mFileList = append(tWriter.mFileList, tFileName)
	}

	err := tWriter.mWriter.Write(tRecords)
	if err != nil {
		return err
	}

	tWriter.mNowLines++

	return nil
}

// 多行写入
func (tWriter *CsvWriter) WriteAll(tRecords [][]string) error {
	// 逐行写入
	for _, record := range tRecords {
		err := tWriter.Write(record)
		if err != nil {
			return err
		}
	}
	// 文件写入缓存推送
	tWriter.mWriter.Flush()
	return nil
}

// 关闭文件读取器
func (tWriter *CsvWriter) Close() []string {
	// 文件写入缓存推送
	tWriter.mWriter.Flush()

	// 关闭文件句柄
	if nil != tWriter.mFileHandle {
		tWriter.mFileHandle.Close()
	}

	return tWriter.mFileList
}

// 快速保存CSV文件
func WriteAll(tFileName string, tRecords [][]string, opts ...WriterOptions) ([]string, error) {
	// 创建writer
	tWriter, err := NewWriter(tFileName, opts...)
	if err != nil {
		return nil, err
	}
	if tWriter != nil {
		// 如果有文件头
		if tWriter.mHeader != nil {
			err := tWriter.Write(tWriter.mHeader)
			if err != nil {
				tWriter.mFileHandle.Close()
				return nil, err
			}
		}
		// 写入所有数据
		err := tWriter.WriteAll(tRecords)
		if err != nil {
			return nil, err
		}
		// 关闭
		return tWriter.Close(), nil
	}
	return nil, errors.New("the writer is nil ")
}
