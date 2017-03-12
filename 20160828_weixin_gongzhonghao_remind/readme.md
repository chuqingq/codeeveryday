# deps

```
cd xx-wechat-remind
git pull xxx-wechat-remind
npm install wechat wechat-api express mongodb date-format date-parser
```

# db

Copy and modify `/etc/mongodb.conf`, bind_ip 127.0.0.1 and logfile xxx.

```
screen
mongod -f mongodb.conf
ctrl-a ctrl-d
```

# start

```
screen
nodejs weixin24.js 2>&1 1>>log &
ctrl-a ctrl-d
```

# TODO
