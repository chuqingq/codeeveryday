## pip3 install tornado

import tornado.ioloop
import tornado.web

class MainHandler(tornado.web.RequestHandler):
    def get(self):
        self.set_header("Access-Control-Allow-Origin", "*")
        self.write("下面是" + self.get_argument('id') + "的年度图书报告：123123123123123")

application = tornado.web.Application([
    (r"/library_annual_report/query", MainHandler),
])

if __name__ == "__main__":
    application.listen(8080)
    tornado.ioloop.IOLoop.instance().start()
