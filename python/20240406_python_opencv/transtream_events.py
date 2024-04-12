#!/usr/bin/env python
# coding: utf-8

import cv2
import time
import json
import logging

logging.basicConfig(
    level=logging.DEBUG,
    format="%(asctime)s - %(levelname)s - %(message)s",
)


class EventsReader:
    def __init__(self, filename):
        self.json_decoder = json.JSONDecoder()
        self.file = open(filename)
        self.eof = False
        self.buffer = ""
        self.cache_event = self.__read_one()
        self.diff_ms = (
            int(time.time() * 1000) - self.cache_event["result"]["timeMs"]
        )

    def get_events(self):
        """根据时间戳获取一组事件"""
        events = []
        detect_ms = int(time.time() * 1000) - self.diff_ms
        while (
            self.cache_event and self.cache_event["result"]["timeMs"] <= detect_ms
        ):
            events.extend(self.cache_event["result"]["items"])
            self.cache_event = self.__read_one()
        return events

    def __read_one(self):
        """从缓冲区或文件中读取一个事件。不计算时间戳"""
        while True:
            try:
                event, idx = self.json_decoder.raw_decode(self.buffer)
                self.buffer = self.buffer[idx:]
                return event
            except json.decoder.JSONDecodeError:
                if self.eof:
                    self.buffer = ""
                    return None
                else:
                    b = self.file.read(1024)
                    if not b:
                        self.eof = True
                    else:
                        self.buffer += b

    def close(self):
        self.file.close()


# 读取视频文件
cap = cv2.VideoCapture("1.mp4")

video_fps = cap.get(cv2.CAP_PROP_FPS)
frame_wait_ms = int(1000 / video_fps)
logging.info(f"视频帧率: {video_fps}, 帧等待时间: {frame_wait_ms}ms")

width = int(cap.get(cv2.CAP_PROP_FRAME_WIDTH))
height = int(cap.get(cv2.CAP_PROP_FRAME_HEIGHT))
logging.info(f"视频宽高: {width}x{height}")

events_reader = EventsReader("transtream_events.txt")

next_frame_ms = int(time.time() * 1000)

# 创建窗口，设置合适的大小和位置
cv2.namedWindow('transtream', cv2.WINDOW_NORMAL)
cv2.moveWindow('transtream', 20, 20) 
cv2.resizeWindow('transtream', int(1920/2), int(1080/2))

last_events = []

# 识别结果
rec = {}

# 循环读取视频帧
while cap.isOpened():
    # 读取视频帧
    ret, frame = cap.read()
    if not ret:
        break
    # 复制一个frame
    frame2 = frame.copy()
    # 获取事件
    events = events_reader.get_events()
    if not events:
        events = last_events
    else:
        last_events = events
    for event in events:
        # 获取框
        x = int(event["loc"]["x0"] * width)
        y = int(event["loc"]["y0"] * height)
        x1 = int(event["loc"]["x1"] * width)
        y1 = int(event["loc"]["y1"] * height)
        color = (255, 0, 0)
        # 如果是车或者人
        if event['type'] == 'person' and 'face' in event and event['face']['id']:
            color = (0, 255, 0)
            id = event['face']['id']
            if id not in rec:
                rec[id] = 0
            rec[id] += 1
            logging.debug(f"人物 {id}")
            cv2.putText(frame2, id, (x, y-10), cv2.FONT_HERSHEY_SIMPLEX, 0.5, color, 2)
            cv2.imshow(id, frame[y:y1, x:x1])
        elif event['type'] == 'car' and 'lp' in event and event['lp']['no']:
            color = (0, 255, 0)
            thickness = 2
            id = event['lp']['no']
            logging.debug(f"车牌 {id}")
            cv2.putText(frame2, id, (x, y-10), cv2.FONT_HERSHEY_SIMPLEX, 0.5, color, 2)
            # cv2.imshow('car', frame[y:y1, x:x1])
        elif 'type' in event:
            cv2.putText(frame2, event['type'], (x, y-10), cv2.FONT_HERSHEY_SIMPLEX, 0.5, color, 2)
        # 画框
        cv2.rectangle(frame2, (x, y), (x1, y1), color, 2)
    # 显示视频帧
    cv2.imshow("transtream", frame2)
    # 等待按键输入
    next_frame_ms += frame_wait_ms
    wait_ms = next_frame_ms-int(time.time() * 1000)
    if wait_ms <= 0:
        wait_ms = 1
    if cv2.waitKey(wait_ms) & 0xFF == ord("q"):
        break

events_reader.close()

# 释放视频文件
cap.release()

# 关闭所有窗口
cv2.destroyAllWindows()

logging.info(f"识别结果: {rec}")