package filesFinder

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type finderFunc func(tFileName string, args ...interface{})

// 遍历指定目录下(包括子目录)的所有指定类型文件，并将文件进行指定操作
func Walk(tRoot string, tFilesType string, tFunc finderFunc, args ...interface{}) error {
	// 无返回函数
	if nil == tFunc {
		return errors.New("no deal function!")
	}

	var (
		tAllType bool            // 遍历所有文件类型标志
		tTypeMap map[string]int8 // 指定文件类型列表
	)

	// tFilesType为`*`或者空字符串表示所有类型文件
	if (`*` == tFilesType) || (`` == tFilesType) {
		tAllType = true
	} else {
		tAllType = false
		// 需查找文件类型分割
		tList := strings.Split(tFilesType, ",")
		tTypeMap = make(map[string]int8)
		for _, tKey := range tList {
			// 文件类型名统一为小写
			tTypeMap["."+strings.ToLower(tKey)] = 1
		}
	}

	// 指定目录遍历
	return filepath.Walk(tRoot,
		func(pathStr string, f os.FileInfo, err error) error {
			// 文件打开失败
			if f == nil {
				return err
			}

			//if f.IsDir() && pathStr != tRoot {
			// if f.IsDir() {
			// 忽略目录
			// 	return nil
			// }

			// 类型匹配
			if tAllType {
				// 调用处理文件
				tFunc(pathStr, args...)
			} else {
				// 文件类型判断(文件类型名统一为小写)
				if _, ok := tTypeMap[strings.ToLower(path.Ext(f.Name()))]; ok {
					// 调用处理文件
					tFunc(pathStr, args...)
				}
			}

			return nil
		})
}

// 遍历指定目录下(不包括子目录)的所有指定类型文件，并将文件进行指定操作
func List(tRoot string, tFilesType string, tFunc finderFunc, args ...interface{}) error {
	// 无返回函数
	if nil == tFunc {
		return errors.New("no deal function!")
	}

	var (
		tAllType bool            // 遍历所有文件类型标志
		tTypeMap map[string]int8 // 指定文件类型列表
	)

	// 获取指定目录下的所有文件及文件夹信息
	dir, err := ioutil.ReadDir(tRoot)
	if nil != err {
		return err
	}

	// tFilesType为`*`或者空字符串表示所有类型文件
	if (`*` == tFilesType) || (`` == tFilesType) {
		tAllType = true
	} else {
		tAllType = false
		// 需查找文件类型分割
		tList := strings.Split(tFilesType, ",")
		tTypeMap = make(map[string]int8)
		for _, tKey := range tList {
			// 文件类型名统一为小写
			tTypeMap["."+strings.ToLower(tKey)] = 1
		}
	}

	// 遍历指定文件目录
	for _, f := range dir {
		//if f.IsDir() && pathStr != tRoot {
		// if f.IsDir() {
		// 忽略目录
		// 	return nil
		// }

		// 类型匹配
		if tAllType {
			// 调用处理文件
			tFunc(tRoot+string(os.PathSeparator)+f.Name(), args...)
		} else {
			// 文件类型判断(文件类型名统一为小写)
			if _, ok := tTypeMap[strings.ToLower(path.Ext(f.Name()))]; ok {
				// 调用处理文件
				tFunc(tRoot+string(os.PathSeparator)+f.Name(), args...)
			}
		}
	}

	return nil
}
