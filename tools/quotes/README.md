# SMTM/tools/quotes

行情数据爬取工具

## 依赖库

### requests

    pip install tushare

### tushare

* 方式1：

    pip install tushare

    如果安装网络超时可尝试国内pip源，如pip install tushare -i <https://pypi.tuna.tsinghua.edu.cn/simple>

* 方式2：

    访问<https://pypi.python.org/pypi/tushare/>下载安装 ，执行 python setup.py install

* 方式3：

    访问<https://github.com/waditu/tushare,将项目下载或者clone>到本地，进入到项目的目录下，

    执行： python setup.py install

## 数据文件(.csv)标题备注

| 标题 | 备注 |
| :----- | :---- |
| Code | 股票代码 |
| Name | 股票名称 |
| Date | 日期 |
| NowPrice | 最新价 |
| PctChange | 涨跌幅 |
| Change | 涨跌额 |
| Volume | 成交量（手） |
| Amount | 成交额 |
| Amplitude | 振幅 |
| TurnoverRatio | 换手率 |
| PE | 市盈率 |
| PB | 市净率 |
| QRR | 量比 |
| MC | 总市值 |
| FAMC | 流通市值 |
| Close | 收盘价 |
| High | 最高价 |
| Low | 最低价 |
| Open | 开盘价 |
| PreClose | 前收盘 |
