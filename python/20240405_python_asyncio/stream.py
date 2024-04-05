#!/usr/bin/env python

from time import sleep
import logging

logging.basicConfig(
    level=logging.DEBUG,  # 设置日志级别为DEBUG，可以输出所有级别的日志
    format="%(asctime)s - %(levelname)s - %(message)s",  # 设置日志格式
)

while True:
    sleep(1)
    logging.info("stream: This is an info message")
