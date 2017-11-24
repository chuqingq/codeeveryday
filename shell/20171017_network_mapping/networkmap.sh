#!/bin/sh

# 需要先建立信任关系

while true
do
    echo `date` "will try reconnect in 2 seconds..."
    ssh -R 0.0.0.0:8090:192.168.54.118:8090 root@121.41.103.23
    sleep 2 
done

