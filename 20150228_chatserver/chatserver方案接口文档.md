chatserver方案接口文档
====

# 连接方式

- app和server之间是tcp over ssl长连接。
- ssl采用单向认证（client无需提供证书）
- keepalive的心跳参数当前取默认值。

# 登录

- app应该建立连接后立刻进行登陆
- app应该在登陆成功后才能发送消息
- app登录成功后会和server一直保持长连接，一旦断连就表示app退出。
- app登陆失败后应该立刻断连（server也会主动断连）
- TODO 服务器内部认证接口

## 请求

```
{
    "action": "login",
    "username": "13770827856",
    "password": "xxxxx"
}
```

## 响应

登录成功：

```
{
    "code": 0,
    "message": "success"
}
```

登录失败：

```
{
    "code": -1,
    "message": "invalid username or password"
}
```

# 消息

不需要认证好友关系，可以发给任意的合法to。

```
{
    "action": "message",
    "from": "13700000001",
    "to": "13700000002",
    "timestamp": "2015-02-28 15:20:32",
    "content": {
        "type": "text/plain",
        "data": "hello"
    }
}
```

- message消息无响应。服务器不保证能够/立刻发到to。
- content.type当前支持文本（text/plain）和图片（image/png），如果是二进制（例如图片），content.data需要base64编码。
- 如果to在线，server直接把这条消息发给to，格式与内容同上；
- 如果to不在线，则server缓存这条消息（增加server自己的timestamp）；
- 当to使用登陆接口成功上线时，server把消息按server自己的timestamp的先后顺序发给to，并清除缓存。
- 暂时不限制缓存消息条数。
