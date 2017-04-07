# deps

```
cd xx-wechat-remind
git pull xxx-wechat-remind
npm install #wechat wechat-api express mongodb date-format date-parser
```

# db

Copy and modify `/etc/mongodb.conf`, `bind_ip` 127.0.0.1 and `logfile` xxx.

```
screen
mongod -f mongodb.conf
ctrl-a ctrl-d
```

# start

```
npm start
# screen -dmS weixin-mp-remind node weixin24.js
```

# /remind/new 引导用户进入页面

需要在沙盒中配置“网页授权获取用户基本信息”中的域名，也可以是IP：

http://121.41.103.23/remind/new
http%3a%2f%2f121.41.103.23%2fremind%2fnew
https://open.weixin.qq.com/connect/oauth2/authorize?appid=wx0288bf03ed5da89b&redirect_uri=http%3a%2f%2f121.41.103.23%2fremind%2fnew&response_type=code&scope=snsapi_base&state=#wechat_redirect

# /remind/get 引导用户进入页面

http://121.41.103.23/remind/get
http%3a%2f%2f121.41.103.23%2fremind%2fget
https://open.weixin.qq.com/connect/oauth2/authorize?appid=wx0288bf03ed5da89b&redirect_uri=http%3a%2f%2f121.41.103.23%2fremind%2fget&response_type=code&scope=snsapi_base&state=#wechat_redirect

# TODO

* 两个页面
