node install
修改index.html中的io(url)的url
node server.js
打开两个标签页（可以是同一局域网内的不同机器）
在一台上点击“接收投屏”，等待投屏
在另一台上点击“投屏”

## 问题
局域网内，如果两台机器不在同一子网，且通过IP访问，无法获取非匿名化的ice candidate。

解决办法（任一）：
1. chrome配置关闭mDNS
2. 发送方使用域名，本地hosts即可

