var express = require('express')
var proxy = require('http-proxy-middleware');

var app = express()

// reverse proxy
var options = {
  target: 'http://127.0.0.1:8000/notfound', // 目标主机
  // hangeOrigin: true,               // 需要虚拟主机站点
  router: {
    '/gpssignin': 'http://127.0.0.1:8091/',
    '/api2': 'http://127.0.0.1:8082'
  }
};

var exampleProxy = proxy(options);  //开启代理功能，并加载配置
app.use('/', exampleProxy);//对地址为’/‘的请求全部转发

app.listen(80)

console.log('Listening on port 80')

