# deps

```
cd xx-wechat-remind
git pull xxx-wechat-remind
npm install #wechat wechat-api express mongodb date-format date-parser
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
npm start
# screen -dmS weixin-mp-remind node weixin24.js
```

# TODO
