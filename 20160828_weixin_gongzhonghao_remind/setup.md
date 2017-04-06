# 获取APPID和APPSECRET

```
APPID=xx
APPSECRET=yy
```

# 获取access-token

```
curl "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=APPID&secret=APPSECRET"
{"access_token":"wdBSFLlhO1192atrja2Kzcm-4lheVcrarTbhnGaVxnOcBrv9Qmw9tX2Fpq5oFTEnJpSc915SvUf2Kgpv6apinCRvd_0AQUX8SqVwBMhXSIuGPWsNBzdwHzdUPBcCq8WxLIFfABAWTQ","expires_in":7200}
```

# 创建菜单

```
curl -XPOST https://api.weixin.qq.com/cgi-bin/menu/create?access_token=qBAiFdsUPpI2RZgG0YiyQ2fqK9ql134nue7pdVEedwNEOFA51-ydH70B-sYwCpsC7NI8om8u13l2NPF9uWvaBpkTa7n1xagaYzmjPqiPvDw7uGgRqdxWG9vfEG31RZoNSJXcAAACTH -d '{"button":[{"type":"view","name":"设置提醒","url":"https://open.weixin.qq.com/connect/oauth2/authorize?appid=wx0288bf03ed5da89b&redirect_uri=http%3a%2f%2f121.41.103.23%2fremind%2fnew&response_type=code&scope=snsapi_base&state=#wechat_redirect"},{"type":"view","name":"查看提醒","url":"https://open.weixin.qq.com/connect/oauth2/authorize?appid=wx0288bf03ed5da89b&redirect_uri=http%3a%2f%2f121.41.103.23%2fremind%2fget&response_type=code&scope=snsapi_base&state=#wechat_redirect"}]}'
curl -XPOST https://api.weixin.qq.com/cgi-bin/menu/create?access_token=qBAiFdsUPpI2RZgG0YiyQ2fqK9ql134nue7pdVEedwNEOFA51-ydH70B-sYwCpsC7NI8om8u13l2NPF9uWvaBpkTa7n1xagaYzmjPqiPvDw7uGgRqdxWG9vfEG31RZoNSJXcAAACTH -d @menu.json
{"errcode":0,"errmsg":"ok"}
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
