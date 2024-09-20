# -*- coding: utf-8 -*-
# 这是一个示例 Python 脚本。
import os
import string

import pandas as pd


def get_money(fileName,excelName,sheetName):
    # 读取excel文件
    df = pd.read_excel(excelName, sheet_name=sheetName)
    # 过滤掉第一行
    # df = df.iloc[1:]

    # 循环每一行
    G = []
    H = []
    I = []
    J = []
    for index, row in df.iterrows():
        # 计算薪资
        GIndex = (1 - row[2]) * row[0] / row[1] - row[3] - row[4] - row[5]
        tmp = GIndex % 100
        if tmp >= 50:
            GIndex+=50
        GIndex = int(GIndex / 100) * 100
        G.append(GIndex)
        HIndex = int(row[0] / row[1])
        H.append(HIndex)
        IIndex = int(HIndex - GIndex - row[3] - row[4] - row[5])
        I.append(IIndex)
        JIndex = round(IIndex / HIndex * 100, 2)
        # 对JIndex拼接%
        J.append(JIndex)
    df['薪资'] = G
    df['税后收入'] = H
    df['最终毛利润'] = I
    df['最终毛利润率'] = J

    df.to_excel('./'+fileName+'薪资.xlsx', index=False, sheet_name='Sheet1')


def get_offer(fileName,excelName,sheetName):
    # 读取excel文件
    df = pd.read_excel(excelName, sheet_name=sheetName)

    # 过滤掉第一行

    # 循环每一行
    A = []
    H = []
    I = []
    J = []
    for index, row in df.iterrows():
        offer = int(((row[6] + row[3] + row[4] + row[5]) * row[1] / (1 - row[2])))
        tmp = offer % 100
        if tmp >= 50:
            offer+=50
        AIndex = int( offer / 100) * 100
        A.append(AIndex)
        HIndex = int(AIndex / row[1])
        H.append(HIndex)
        IIndex = int(HIndex - row[6] - row[3] - row[4] - row[5])
        I.append(IIndex)
        JIndex = round(IIndex / HIndex * 100, 2)
        # 保留一位小数
        J.append(JIndex)

    df['总报价(含税)'] = A
    df['税后收入'] = H
    df['最终毛利润'] = I
    df['最终毛利润率'] = J
    df.to_excel('./'+fileName+'报价.xlsx', index=False, sheet_name='Sheet1')


def main(fileName,excelName,sheetName):

    print("开始计算.....请稍后")
    try:
        get_money(fileName,excelName,sheetName)
    except Exception as e:
        print("薪资计算出错")
        print(e)
    try:
        get_offer(fileName,excelName,sheetName)
    except Exception as e:
        print("报价计算出错")
        print(e)
    print("结束！")


# 按间距中的绿色按钮以运行脚本。
if __name__ == '__main__':
    sheetName = 'Sheet1'

    dir_path = './'

    # 获取目录下的所有文件名
    file_names = os.listdir(dir_path)

    # 打印所有文件名
    for file_name in file_names:
        if '薪资' in file_name:
            continue
        if '报价' in file_name:
            continue
        if 'py' in file_name:
            continue
        names = file_name.split('.')
        fileName = names[0]
        excelName = file_name
        main(fileName,excelName,sheetName)

# 访问 https://www.jetbrains.com/help/pycharm/ 获取 PyCharm 帮助
