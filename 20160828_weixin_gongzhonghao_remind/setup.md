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
curl -XPOST https://api.weixin.qq.com/cgi-bin/menu/create?access_token=wdBSFLlhO1192atrja2Kzcm-4lheVcrarTbhnGaVxnOcBrv9Qmw9tX2Fpq5oFTEnJpSc915SvUf2Kgpv6apinCRvd_0AQUX8SqVwBMhXSIuGPWsNBzdwHzdUPBcCq8WxLIFfABAWTQ -d '{"button":[{"type":"view","name":"设置提醒","url":"http://121.41.103.23/remind/new"},{"type":"view","name":"查看提醒","url":"http://121.41.103.23/remind/get"}]}'
{"errcode":0,"errmsg":"ok"}
```
