nohup ./proxy sps -S socks -T tcp -P 127.0.0.1:1081 -t tcp -p :33082 &

33082是代理暴露的端口 ， 127.0.0.1 是上游的Trojan的信息

