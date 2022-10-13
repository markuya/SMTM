package filesFinder

import (
	"testing"
)

// 闭包模式:
// 推荐使用闭包模式
func test_walk_1(t *testing.T, tPath string, tFileTypes string) {
	var tList []string
	Walk(tPath, tFileTypes,
		func(tFileName string, arg ...interface{}) {
			tList = append(tList, tFileName)
		})

	for _, tFileName := range tList {
		t.Logf("%s\n", tFileName)
	}
	t.Logf("%s tFileTypes[%s]\n\n", "test_walk_1", tFileTypes)
}

// 传参模式:
// 注意，本模式下如果需要将文件名传出，【不能】使用传入【数组】将文件名写入数组的方式获取结果
func test_walk_2(t *testing.T, tPath string, tFileTypes string) {
	tList := make(map[string]string)
	Walk(tPath, tFileTypes,
		func(tFileName string, arg ...interface{}) {
			if (nil != arg) && (len(arg) > 0) {
				if tList2, ok := arg[0].(map[string]string); ok {
					tList2[tFileName] = tFileName
				}
			}
		}, tList)

	for _, tFileName := range tList {
		t.Logf("%s\n", tFileName)
	}
	t.Logf("%s tFileTypes[%s]\n\n", "test_walk_2", tFileTypes)
}

// 闭包模式:
// 推荐使用闭包模式
func test_list_1(t *testing.T, tPath string, tFileTypes string) {
	var tList []string
	List(tPath, tFileTypes,
		func(tFileName string, arg ...interface{}) {
			tList = append(tList, tFileName)
		})

	for _, tFileName := range tList {
		t.Logf("%s\n", tFileName)
	}
	t.Logf("%s tFileTypes[%s]\n\n", "test_list_1", tFileTypes)
}

// 传参模式:
// 注意，本模式下如果需要将文件名传出，【不能】使用传入【数组】将文件名写入数组的方式获取结果
func test_list_2(t *testing.T, tPath string, tFileTypes string) {
	tList := make(map[string]string)
	List(tPath, tFileTypes,
		func(tFileName string, arg ...interface{}) {
			if (nil != arg) && (len(arg) > 0) {
				if tList2, ok := arg[0].(map[string]string); ok {
					tList2[tFileName] = tFileName
				}
			}
		}, tList)

	for _, tFileName := range tList {
		t.Logf("%s\n", tFileName)
	}
	t.Logf("%s tFileTypes[%s]\n\n", "test_list_2", tFileTypes)
}

// 测试总入口
func Test(t *testing.T) {
	test_list_1(t, `./testPath`, `*`)
	test_list_1(t, `./testPath`, ``)
	test_list_1(t, `./testPath`, `Csv,txt`)
	test_list_1(t, `./testPath`, `CSV`)
	test_list_1(t, `./testPath`, `xls`)

	test_list_2(t, `./testPath`, `*`)
	test_list_2(t, `./testPath`, ``)
	test_list_2(t, `./testPath`, `Csv,txt`)
	test_list_2(t, `./testPath`, `CSV`)
	test_list_2(t, `./testPath`, `xls`)

	test_walk_1(t, `./testPath`, `*`)
	test_walk_1(t, `./testPath`, ``)
	test_walk_1(t, `./testPath`, `Csv,txt`)
	test_walk_1(t, `./testPath`, `CSV`)
	test_walk_1(t, `./testPath`, `xls`)

	test_walk_2(t, `./testPath`, `*`)
	test_walk_2(t, `./testPath`, ``)
	test_walk_2(t, `./testPath`, `Csv,txt`)
	test_walk_2(t, `./testPath`, `CSV`)
	test_walk_2(t, `./testPath`, `xls`)
}
