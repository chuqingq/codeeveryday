[OpenSSL自签发自建CA签发SSL证书 - just~do~it - 博客园 (cnblogs.com)](https://www.cnblogs.com/will-space/p/11913744.html)



https、openssl



## 目的

让本机浏览器（包括chrome、edge、firefox等）访问自己搭建的服务，安全、可信任。



## CA服务器搭建

1，生成根CA私钥

openssl genrsa -out ca.key 2048

2，生成根CA证书

openssl req -x509 -new -key ca.key -out ca.crt -days 3650

以上CA服务器搭建完成



## 颁发证书

创建服务私钥

openssl genrsa -out http.key 2048

创建证书请求

openssl req -new -key http.key -out http.csr

CN需要是服务器Host名，例如wslserver或者172.30.111.213。应该需要在http.ext的SubjectAlternativeName中体现。



解决Chrome不能识别证书通用名称NET::ERR_CERT_COMMON_NAME_INVALID错误

```Rust
> vim http.ext
keyUsage = nonRepudiation, digitalSignature, keyEncipherment
extendedKeyUsage = serverAuth, clientAuth
subjectAltName=@SubjectAlternativeName

[SubjectAlternativeName ]
IP.1=172.30.111.213
IP.2=127.0.0.1
DNS.1=wslserver
DNS.2=localhost
```

签发证书

openssl x509 -req -in http.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out http.crt -days 3650 -sha256 -extfile http.ext





## 验证

sudo go run go_httpsserver.go

![](https://secure2.wostatic.cn/static/2YTMgdy66c16fzthVQt6Cn/image.png)

导入ca.crt。则无论IP还是域名（需要配置hosts）都能访问成功：

![](https://secure2.wostatic.cn/static/moFXNCWzaq2PRyjrkqa69G/image.png)

![](https://secure2.wostatic.cn/static/qth4QXnXukmHuDZM3kzRVG/image.png)




