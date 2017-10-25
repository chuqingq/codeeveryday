#!/usr/bin/python -u
# -*- coding: utf-8 -*-

## pip install uiautomator tornado

import os
import time
from uiautomator import Device
import tornado.ioloop
import tornado.web

SERIAL = '0710ad7b00f456bb'
d = None

class MainHandler(tornado.web.RequestHandler):
    def get(self):
        print "start click_card"
        d = Device(SERIAL)
        os.system('adb -s ' + SERIAL + ' shell am start io.appium.unlock/.Unlock')
        time.sleep(1)
        d.press.home()
        time.sleep(1)
        d(text=u'一键打卡').click()
	self.write('success')
        print "click_card success"

if __name__ == "__main__":
    print "starting..."
    app = tornado.web.Application([
        (r"/click_card", MainHandler),
    ])
    app.listen(8090)
    tornado.ioloop.IOLoop.current().start()

