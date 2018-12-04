#!/bin/sh

while true; do
	sleep 2
	echo `date +'%F %T'`" reconnecting..."
	sshpass -p '1qaz2wsx!QAZ@WSX' ssh -R 0.0.0.0:80:192.168.0.200:80 -R 0.0.0.0:443:192.168.0.200:443 root@www.dilu.com top
done

