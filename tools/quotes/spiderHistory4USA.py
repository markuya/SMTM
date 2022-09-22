#!/usr/bin/env python
#coding=utf-8
import platform
import sys

# linux系统 设置默认编码为utf-8
if(platform.system() == "Linux"):
    reload(sys)
    sys.setdefaultencoding('utf-8')

from telnetlib import DO
import requests
import yfinance as yf
import sys,getopt
import os,os.path
import json
import time
import shutil

# 读取JSON配置文件
def getConfigs( tCfgFile ):
	print ('getConfigs: tCfgFile=', tCfgFile)
	# 判断文件是否存在
	if not os.path.exists(tCfgFile):
		return
	#end if
    
	with open(tCfgFile,'r') as f:
		data = json.load(f)
	#end with
	return data
#end def

# 获取列表
def getList(tParams):
    tStockList = []
    # 参数验证
    if ('CNT' not in tParams) or ('Url' not in tParams) or ('jQueryKey' not in tParams) or ('Parameter' not in tParams):
        print("getList：failed!输入参数不完整，请检查参数录入配置文件!")
        os._exit(0)

    # 当前时间戳
    tTime = int(time.time())
    tjQueryKey = '%s%d' % (tParams['jQueryKey'],tTime)
    # 请求页码
    tNowPage = 1
    while True:
        # 请求URL拼接
        tUrl = '%s%s&pn=%d%s&_=%d' % (tParams['Url'],tjQueryKey,tNowPage,tParams['Parameter'],tTime)
        #print('tUrl:',tUrl)
        res = requests.get(tUrl)
        #result = res.text.split(tjQueryKey)[1].split("(")[1].split(");")[0]
        result2 = res.text.split(tjQueryKey)[1][1:][:-2]
        # 数据解JSON
        result_json = json.loads(result2)
        # 判断是否请求到正确数据
        if ('rc' in result_json) and ('data' in result_json) and (0 == result_json['rc']):
            for i in result_json['data']['diff']:
                tStockList.append(i['f12'])
            tNowPage += 1
        else:
            break
        #end if
    #end while

    # 判断股票ID列表
    return tStockList
#end def

# 拉取对应股票的历史数据
def pullStockHisquotes( tCode ,tParams ):
    # 参数验证
    if ('CNT' not in tParams) or ('Url' not in tParams) or ('Parameter' not in tParams) or ('Output' not in tParams):
        print("getList：failed!输入参数不完整，请检查参数录入配置文件!")
        os._exit(0)

    # 输出文件名
    tSaveFile = tParams['Output']+tCode+'.csv'

    # 判断文件是否存在
    if os.path.exists(tSaveFile):
        os.remove(tSaveFile)

    # 创建目录
    if not os.path.exists(tParams['Output']):
        os.makedirs(tParams['Output'])

    data = yf.download(tCode, start="1990-01-01", end=time.strftime('%Y%m%d',time.localtime()))
    data.to_csv(tSaveFile)

#end def

# 主函数
def main(argv):
	# 读取指定配置
	tCfg = getConfigs( './conf/Export.json' )

	# 拉取各地区交易所股票列表
	if ('SEList' in tCfg) and ('USA' in tCfg['SEList']) and ('HIS' in tCfg) and ('USA' in tCfg['HIS']):
		tStockList = getList(tCfg['SEList']['USA'])
		tLen = len(tStockList)
		tNum = 0
		for tCode in tStockList:
			pullStockHisquotes(tCode,tCfg['HIS']['USA'])
			tNum += 1
			print("# pull code(%s) history finish! %d/%d" % (tCode,tNum,tLen))
        #end for
	#end if

	print("# pull all history finish!")
#end def

# 默认运行
if __name__ == "__main__":
   main(sys.argv[1:])
