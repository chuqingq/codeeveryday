from aip import AipOcr
import base64

""" 你的 APPID AK SK """
APP_ID = '10682639'
API_KEY = 'yYkzlkhdkO4CsOo7fGHZmgXx'
SECRET_KEY = 'DWIxGQsDGbuTY9v7qCC5t5VqOkDZC8c1'

client = AipOcr(APP_ID, API_KEY, SECRET_KEY)

""" 读取图片 """
def get_file_content(filePath):
    with open(filePath, 'rb') as fp:
        return fp.read()

image = get_file_content('11.jpg')
imagebase64 = base64.b64encode(image)
print(len(imagebase64))

# """ 调用通用文字识别, 图片参数为本地图片 """
# client.basicGeneral(image);

""" 如果有可选参数 """
options = {}
# options["language_type"] = "CHN_ENG"
# options["detect_direction"] = "true"
# options["detect_language"] = "true"
# options["probability"] = "true"

""" 带参数调用通用文字识别, 图片参数为本地图片 """
# client.basicGeneral(image, options)
#retbig = client.accurate(image, options)
retbig = client.general(image, options)
print(retbig)

res = ''
for w in retbig['words_result']:
	res += w['words'] + ' '
print(res)

# {'log_id': 4096811966239625715, 'direction': 0, 'words_result_num': 15, 'words_result': [{'words': '的', 'probability': {'variance': 0, 'average': 0.997502, 'min': 0.997502}}, {'words': '郎咸平全集', 'probability': {'variance': 0.052254, 'average': 0.834156, 'min': 0.394306}}, {'words': '货币战', 'probability': {'variance': 0.030865, 'average': 0.840569, 'min': 0.595434}}, {'words': '孙子兵法', 'probability': {'variance': 0.002903, 'average': 0.932002, 'min': 0.843681}}, {'words': '为', 'probability': {'variance': 0, 'average': 0.999037, 'min': 0.999037}}, {'words': '白鹿原', 'probability': {'variance': 0.016763, 'average': 0.866531, 'min': 0.690722}}, {'words': '城', 'probability': {'variance': 0, 'average': 0.999557, 'min': 0.999557}}, {'words': '纲', 'probability': {'variance': 0, 'average': 0.924304, 'min': 0.924304}}, {'words': '著', 'probability': {'variance': 0.092798, 'average': 0.655222, 'min': 0.350595}}, {'words': '听', 'probability': {'variance': 0, 'average': 0.986715, 'min': 0.986715}}, {'words': '我', 'probability': {'variance': 0, 'average': 0.997738, 'min': 0.997738}}, {'words': '的', 'probability': {'variance': 0, 'average': 0.999326, 'min': 0.999326}}, {'words': '郎咸平著', 'probability': {'variance': 0.032205, 'average': 0.839712, 'min': 0.54961}}, {'words': '计', 'probability': {'variance': 0, 'average': 0.993832, 'min': 0.993832}}, {'words': '方击', 'probability': {'variance': 0.027372, 'average': 0.606964, 'min': 0.44152}}], 'language': -1}
# 的 郎咸平全集 货币战 孙子兵法 为 白鹿原 城 纲 著 听 我 的 郎咸平著 计 方击

# 需要把每个纵列的所有文字整合到一起
res = {}
for w in retbig['words_result']:
	left = w['location']['left'] + w['location']['width']/2
	exists = False
	for l in res:
		if left < l+25 and left > l-25:
			exists = True
			res[l] += ' '+w['words']
	if not exists:
		res[left] = w['words']

keys = sorted(res)
for l in keys:
	print(res[l])
