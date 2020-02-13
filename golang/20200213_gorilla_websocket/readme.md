```
go get github.com/gorilla/websocket
go build remoteMainControl.go
go build remoteAgent.go
```

# 说明

在代码中不带任何keepalive逻辑情况下，抓包看，服务端会每15秒给客户端发一次keepalive；
如果断连网络连接（拔掉网线），服务端大约2分30秒能感知到断连；客户端感知不到。
客户端加上keepalive代码后，拔掉网线起2分30秒左右能感知到断连。

