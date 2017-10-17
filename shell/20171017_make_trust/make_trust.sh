#!/bin/sh

# make_trust.sh 建立信任关系

# 生成公私钥对
ssh-keygen -t rsa

# 读取生成的公钥
ID_RSA_PUB=`cat $HOME/.ssh/id_rsa.pub`

# 写入到远端
ssh root@121.41.103.23 "mkdir /root/.ssh && echo $ID_RSA_PUB >> /root/.ssh/authorized_keys"

