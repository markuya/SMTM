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
import sys,getopt
import os,os.path
import json
import time
import shutil

# 保存信息到CSV文件
def saveTocsv(data, files):
    # 文件不存在 创建文件头和抬头
    if not os.path.exists(files):
        with open(files, "a+") as f:
            f.write("Code,Name,NowPrice,PctChange,Change,Volume,Amount,Amplitude,TurnoverRatio,PE,QRR,High,Low,Open,PreClose,PB\n")
            f.close()
    
    # 写入文件内容
    with open(files, "a+") as f:
        for i in data['diff']:
            Code = i['f12']
            Name = i['f14']
            Close = i['f2']
            ChangePercent = i["f3"]
            Change = i['f4']
            Volume = i['f5']
            Amount = i['f6']
            Amplitude = i['f7']
            TurnoverRate = i['f8']
            PERation = i['f9']
            VolumeRate = i['f10']
            Hign = i['f15']
            Low = i['f16']
            Open = i['f17']
            PreviousClose = i['f18']
            PB = i['f22']
            #row = '{},{},{},{},{},{},{},{},{},{},{},{},{},{},{},{}'.format(
            #    Code,Name,Close,ChangePercent,Change,Volume,Amount,Amplitude,
            #    TurnoverRate,PERation,VolumeRate,Hign,Low,Open,PreviousClose,PB)
            row = '%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s' % (
                Code,Name,Close,ChangePercent,Change,Volume,Amount,Amplitude,
                TurnoverRate,PERation,VolumeRate,Hign,Low,Open,PreviousClose,PB)
            f.write(row)
            f.write('\n')
	#endif
#end def

# 读取JSON配置文件
def getConfigs( tCfgFile ):
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
    # 参数验证
    if ('CNT' not in tParams) or ('Url' not in tParams) or ('jQueryKey' not in tParams) or ('Parameter' not in tParams) or ('Output' not in tParams) or ('FileName' not in tParams) or ('RecSTime' not in tParams) or ('RecETime' not in tParams):
        print("getList：failed!输入参数不完整，请检查参数录入配置文件!")
        os._exit(0)

    # 当前时间小时数
    tHour = time.localtime().tm_hour
    if tHour < tParams['RecSTime'] or tHour > tParams['RecETime']:
        return

    # 输出文件名
    tSaveFile = tParams['Output']+tParams['FileName']+time.strftime('%Y%m%d',time.localtime())+'.csv'

    # 判断文件是否存在
    if os.path.exists(tSaveFile):
        os.remove(tSaveFile)

    # 创建目录
    if not os.path.exists(tParams['Output']):
        os.makedirs(tParams['Output'])
        
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
            # 保存数据到CSV
            saveTocsv(result_json['data'],tSaveFile)
            tNowPage += 1
        else:
            break
        #end if
    #end while
#end def

# 主函数
def main(argv):
	# 读取指定配置
	tCfg = getConfigs( './conf/Export.json' )
	
	# 拉取各地区交易所股票列表
	if ('SEList' in tCfg):
		for tInfo in tCfg['SEList'].values():
			getList(tInfo)
	#end for
#end def

# 默认运行
if __name__ == "__main__":
   main(sys.argv[1:])
