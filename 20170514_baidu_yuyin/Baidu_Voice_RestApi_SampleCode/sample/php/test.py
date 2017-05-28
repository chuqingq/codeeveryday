# -*- coding: utf-8 -*-

from urllib import request
import json
import base64

filename = '2017-05-14_16_23_18.wav'

cuid = "5415587"
apiKey = "UPKOHCZhfTNMEI9B1Xnf5PRK"
secretKey = "GYHsMLE820jkw1cBGnNbIzoeVVxdDUfs"
auth_url = "https://openapi.baidu.com/oauth/2.0/token?grant_type=client_credentials&client_id="+apiKey+"&client_secret="+secretKey;
response = request.urlopen(auth_url)
response = json.loads(response.read())
token = response['access_token']
print('token: ' + token)
# 读取文件内容
content = open('./'+filename, 'rb').read() 
# base64编码
base_data = base64.b64encode(content)
# 请求体
array = {
    "format": "wav",
    "rate": 8000,
    "channel": 1,
    # "lan" => "zh",
    "token": token,
    "cuid": cuid,
    # "url" => "http://www.xxx.com/sample.pcm",
    # "callback" => "http://www.xxx.com/audio/callback",
    "len": len(content),
    "speech": bytes.decode(base_data)
}
body = json.dumps(array)
# contentLength = 'ContentLength: ' + len(body)
# 发送语音识别请求
url = "http://vop.baidu.com/server_api"
req = request.Request(url=url, data=str.encode(body), method='POST')
response = request.urlopen(req)
response = json.loads(response.read())
print('response: ', response)