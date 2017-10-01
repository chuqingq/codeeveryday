#!/bin/bash

batch_send_res=`curl -X POST \
  'https://xchannel.powerapp.io/channel/v1/_batchsend?authorization=aksk-auth-v1%2FhiactionAK%2F2017-08-17T14%3A34%3A22Z%2F0%2Fhost%2Fbd322dede008d5c55eb2f18e7d030d8ff08b8469679222ee4c3ef54345add30c' \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -H 'postman-token: cd2c5f68-7d91-a90f-7f95-44c6a750b103' \
  -d '{
"appId":1,
"deviceIds": [
"FE7186C21A85B2CD"
],
"message":"{\"action\":\"com.huawei.hiaction.express.notice\",\"msg\":{\"title\":\"1\",\"content\":\"1\",\"index\":1},\"type\":0}"
}' 2>&1`

#echo "batch_send: $batch_send_res"

function do_mail
{
	# 提前安装mail命令。sudo apt-get install mailutils。选择internet
        res=`echo -e "$1" | mail -s "batch_send error" -t xuwentao1@huawei.com,chuqingqing@huawei.com`
        #echo "mail: $res"
}

case $batch_send_res in
        *"requestId"*) exit ;;
        #*"requestId"*) do_mail "batch_send success" ;;
        *) do_mail "$batch_send_res" ;;
esac


