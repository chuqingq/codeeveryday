1、通过strace -f sbing/nginx跟踪发现，nginx没有dup这个listen socket。（之前看h2o是dup的）
2、使用新的openresty，前后是否配置reuseport发现，通过sudo strace -f openresty确认新版本在支持reuseport时也没有dup。
3、整个nginx的dup只有一个场景，就是关闭标准输入输出错误描述符时使用。

