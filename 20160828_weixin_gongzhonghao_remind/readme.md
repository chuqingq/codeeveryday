# deps

```
cd xx-wechat-remind
git pull xxx-wechat-remind
npm install wechat wechat-api express mongodb date-format date-parser
```

# db

Copy and modify `/etc/mongodb.conf`, use 127.0.0.1.

```mongod -f mongodb.conf```

# start

```node weixin23.js```

# TODO

* ~把api.sendText换成api.sendTemplate，确保不会超时~
* 接收提醒时回复text： Set remind success!\nTime: 2017-03-12 12:30:00\nContent: 提醒我中午吃饭
* 到时间发送模板消息： Remind\nTime:  2017-03-12 12:30:00\nContent: 提醒我中午吃饭
* 把date-parser和date-format换成fecha，只支持201703122359 YYYYMMDDHHmm和03122359两种格式

