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
import csv

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
    tSaveFile = tParams['Output']+'code-'+tCode+'.csv'

    # 判断文件是否存在
    if os.path.exists(tSaveFile):
        os.remove(tSaveFile)

    # 创建目录
    if not os.path.exists(tParams['Output']):
        os.makedirs(tParams['Output'])

    if tCode[0] == '6':
        tUrl = tParams['Url']+"0"+tCode+"&start=19900101&end="+time.strftime('%Y%m%d',time.localtime())+tParams['Parameter']
        #tUrl = "http://quotes.money.163.com/service/chddata.html?code=0"+tCode+"&start=19900101&end=20220921&fields=TCLOSE;HIGH;LOW;TOPEN;LCLOSE;CHG;PCHG;TURNOVER;VOTURNOVER;VATURNOVER;TCAP;MCAP"
    else:
        tUrl = tParams['Url']+"1"+tCode+"&start=19900101&end="+time.strftime('%Y%m%d',time.localtime())+tParams['Parameter']
        #tUrl = "http://quotes.money.163.com/service/chddata.html?code=1"+tCode+"&start=19900101&end=20220921&fields=TCLOSE;HIGH;LOW;TOPEN;LCLOSE;CHG;PCHG;TURNOVER;VOTURNOVER;VATURNOVER;TCAP;MCAP"

    print('tUrl:',tUrl)
    res = requests.get(tUrl)
    
    # 数据存入临时文件
    tTmpFile = './tmp-'+tCode+'.csv'
    f_tmp = open(tTmpFile, "wb")
    f_tmp.write(res.content)
    f_tmp.close()

    f2 = open(tSaveFile, "w")

    # 新文件标题
    tHeader = ['Date','Code','Name','Close','High','Low','Open','PreClose','Change','PctChange','TurnoverRatio','Volume','Amount','MC','FAMC']

    with open(tTmpFile) as f_tmp:
        f_reader = csv.reader(f_tmp)
        f_writer = csv.writer(f2)
        # 跳过原标题，写入新标题
        next(f_reader)
        f_writer.writerow(tHeader)
        for each_row in f_reader:
            f_writer.writerow(each_row)
    
    # 删除临时文件
    os.remove(tTmpFile)
    f2.close()
#end def

# 主函数
def main(argv):
	# 读取指定配置
	tCfg = getConfigs( './conf/Export.json' )

	# 拉取各地区交易所股票列表
	if ('SEList' in tCfg) and ('China' in tCfg['SEList']) and ('HIS' in tCfg) and ('China' in tCfg['HIS']):
		tStockList = getList(tCfg['SEList']['China'])
		tLen = len(tStockList)
		tNum = 0
		for tCode in tStockList:
			pullStockHisquotes(tCode,tCfg['HIS']['China'])
			tNum += 1
			print("# pull code(%s) history finish! %d/%d" % (tCode,tNum,tLen))
        #end for
	#end if

	print("# pull all history finish!")
#end def

# 默认运行
if __name__ == "__main__":
   main(sys.argv[1:])
