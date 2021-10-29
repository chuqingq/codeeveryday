import pandas as pd

zhuanli = pd.read_excel('1111.xls').to_dict('records')

results.sort(key = lambda e: e['序号'])
pd.DataFrame(results).to_excel('yyyy.xlsx')

