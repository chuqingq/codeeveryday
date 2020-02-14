```
go get github.com/gorilla/websocket
go build remoteMainControl.go
go build remoteAgent.go
```

# 说明

go 1.13.x开始，客户端和服务端默认都会加tcp keepalive 15秒。
经过验证拔网线2分30秒左右后都能感知到断连。

