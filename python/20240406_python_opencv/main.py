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
        self.buffer = ""
        self.cache_event = self.__read_one()
        self.diff_ms = (
            int(time.time() * 1000) - self.cache_event["detectResult"]["timeMs"]
        )

    def get_events(self):
        """根据时间戳获取一组事件"""
        events = []
        detect_ms = int(time.time() * 1000) - self.diff_ms
        while (
            self.cache_event and self.cache_event["detectResult"]["timeMs"] <= detect_ms
        ):
            events.extend(self.cache_event["detectResult"]["items"])
            self.cache_event = self.__read_one()
        return events

    def __read_one(self):
        """从缓冲区或文件中读取一个事件。不计算时间戳"""
        buf_len = 4096*2
        while len(self.buffer) < buf_len:
            b = self.file.read(buf_len)
            # logging.debug(f"read {len(b)} bytes")
            if not b:
                break
            self.buffer += b
        if not self.buffer:
            return None
        # logging.debug(f"buffer: {self.buffer}")
        try:
            event, idx = self.json_decoder.raw_decode(self.buffer)
            self.buffer = self.buffer[idx:]
            return event
        except json.decoder.JSONDecodeError:
            self.buffer = ""
            return None

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

# 识别结果
rec = {}

# 循环读取视频帧
while cap.isOpened():
    # 读取视频帧
    ret, frame = cap.read()
    if not ret:
        break
    # 获取事件
    events = events_reader.get_events()
    # if len(events) != 0:
    #     logging.debug(f"events count: {len(events)}")
    for event in events:
        # logging.debug(event)
        # 获取框
        x = int(event["location"]["left"] * width)
        y = int(event["location"]["top"] * height)
        x1 = int(event["location"]["right"] * width)
        y1 = int(event["location"]["bottom"] * height)
        color = (255, 0, 0)
        # 如果是车或者人
        if event['type'] == 'person' and event['person']['id']:
            color = (0, 255, 0)
            id = event['person']['id']
            logging.debug(f"人物 {id}")
            if id not in rec:
                rec[id] = 1
            else:
                rec[id] += 1
            cv2.putText(frame, id, (x, y-10), cv2.FONT_HERSHEY_SIMPLEX, 0.5, color, 2)
            cv2.imshow('person', frame[y:y1, x:x1])
        elif event['type'] == 'car' and event['lp']['number']:
            color = (0, 0, 255)
            id = event['lp']['number']
            if id not in rec:
                rec[id] = 1
            else:
                rec[id] += 1
            logging.debug(f"车牌 {id}")
            cv2.putText(frame, id, (x, y-10), cv2.FONT_HERSHEY_SIMPLEX, 0.5, color, 2)
            cv2.imshow('car', frame[y:y1, x:x1])
        # 画框
        cv2.rectangle(frame, (x, y), (x1, y1), color, 2)
    # 显示视频帧
    cv2.imshow("transtream", frame)
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

logging.info(f"识别结果: {json.dumps(rec, indent=4)}")