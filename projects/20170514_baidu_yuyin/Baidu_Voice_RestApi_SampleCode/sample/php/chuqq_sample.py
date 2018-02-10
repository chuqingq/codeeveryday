# -*- coding: utf-8 -*-
from urllib import request
import json
import base64

from pyaudio import PyAudio, paInt16 
import numpy as np 
from datetime import datetime 
import wave 

FILENAME = 'temp.wav'

# 将data中的数据保存到名为filename的WAV文件中
def save_wave_file(filename, data):
    wf = wave.open(filename, 'wb')
    wf.setnchannels(1)
    wf.setsampwidth(2)
    wf.setframerate(SAMPLING_RATE)
    wf.writeframes(b"".join(data))
    wf.close()

# 把音频文件发到百度语音
def baidu_yuyin(filename):
    cuid = "5415587"
    apiKey = "UPKOHCZhfTNMEI9B1Xnf5PRK"
    secretKey = "GYHsMLE820jkw1cBGnNbIzoeVVxdDUfs"
    auth_url = "https://openapi.baidu.com/oauth/2.0/token?grant_type=client_credentials&client_id="+apiKey+"&client_secret="+secretKey;
    response = request.urlopen(auth_url)
    response = json.loads(response.read())
    token = response['access_token']
    # print('token: ' + token)
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
    print(response['result'])

NUM_SAMPLES = 2000      # pyAudio内部缓存的块的大小
SAMPLING_RATE = 8000    # 取样频率
LEVEL = 1500            # 声音保存的阈值
COUNT_NUM = 20          # NUM_SAMPLES个取样之内出现COUNT_NUM个大于LEVEL的取样则记录声音
SAVE_LENGTH = 8         # 声音记录的最小长度：SAVE_LENGTH * NUM_SAMPLES 个取样

# 开启声音输入
pa = PyAudio() 
stream = pa.open(format=paInt16, channels=1, rate=SAMPLING_RATE, input=True, frames_per_buffer=NUM_SAMPLES) 

save_count = 0 
save_buffer = [] 

# 开启声音输入
pa = PyAudio() 
stream = pa.open(format=paInt16, channels=1, rate=SAMPLING_RATE, input=True, 
                frames_per_buffer=NUM_SAMPLES) 
# stream.stop_stream()

save_count = 0
save_buffer = []
print('recording...')

while True:
    # stream.start_stream()
    # 读入NUM_SAMPLES个取样
    string_audio_data = stream.read(NUM_SAMPLES)
    # stream.stop_stream()
    # 将读入的数据转换为数组
    audio_data = np.fromstring(string_audio_data, dtype=np.short)
    print(audio_data, type(audio_data))
    # 计算大于LEVEL的取样的个数
    large_sample_count = np.sum( audio_data > LEVEL )
    # print(np.max(audio_data))
    # 如果个数大于COUNT_NUM，则至少保存SAVE_LENGTH个块
    if large_sample_count > COUNT_NUM: 
        save_count = SAVE_LENGTH 
    else:
        save_count -= 1

    if save_count < 0:
        save_count = 0

    if save_count > 0:
        # 将要保存的数据存放到save_buffer中
        save_buffer.append( string_audio_data ) 
    else: 
        # 将save_buffer中的数据写入WAV文件，WAV文件的文件名是保存的时刻
        if len(save_buffer) > 0:
            print('recorded:')
            stream.stop_stream()
            # filename = datetime.now().strftime("%Y-%m-%d_%H_%M_%S") + ".wav" 
            filename = FILENAME
            save_wave_file(filename, save_buffer) 
            save_buffer = []
            # 语音识别
            baidu_yuyin(filename)
            print('recording...')
            stream.start_stream()
    save_count = 0

# response:  {'corpus_no': '6419959540086987308', 'err_msg': 'success.', 'err_no': 0, 'result': ['娄底好，'], 'sn': '54834668321494763311'}
# Traceback (most recent call last):
#   File "php/chuqq_sample.py", line 73, in <module>
#     string_audio_data = stream.read(NUM_SAMPLES)
#   File "/usr/local/lib/python3.6/site-packages/pyaudio.py", line 608, in read
#     return pa.read_stream(self._stream, num_frames, exception_on_overflow)
# OSError: [Errno -9981] Input overflowed
