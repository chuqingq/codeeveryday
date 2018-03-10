package main

import (
	"log"
	"time"
)

var starttime int64

func main() {
	log.Printf("starting...")
	ch := make(chan bool, 1)
	go func() {
		<-ch
		log.Printf("diff ns: %v", time.Now().UnixNano()-starttime)
	}()

	time.Sleep(2 * time.Second)
	// log.Printf("%v", time.Now().UnixNano())
	starttime = time.Now().UnixNano()
	ch <- true
	time.Sleep(2 * time.Second)
}

// macos:
// $ go run test_channel_delay.go
// 2018/02/11 19:19:19 starting...
// 2018/02/11 19:19:21 1518347961360782000
// 2018/02/11 19:19:21 1518347961360941000
// lubuntu@hp
// $ ./test_channel_delay
// 2018/02/12 19:14:52 starting...
// 2018/02/12 19:14:54 1518434094798942301
// 2018/02/12 19:14:54 1518434094799042050
// chuqq@chuqq-hp:~/temp/codeeveryday/c/20180209_futex$ ./test_channel_delay
// 2018/02/12 19:15:14 starting...
// 2018/02/12 19:15:16 1518434116993779883
// 2018/02/12 19:15:16 1518434116993938206
// lubuntu@virtualbox
//  go run test_channel_delay.go
// 2018/02/22 19:29:26 starting...
// 2018/02/22 19:29:28 diff ns: 143253
/*
chuqq@chuqq:~/work/codeeveryday/c/20180209_futex$ ./test_channel_delay
2018/03/10 12:39:30 starting...
2018/03/10 12:39:32 diff ns: 13436
chuqq@chuqq:~/work/codeeveryday/c/20180209_futex$ ./test_channel_delay
2018/03/10 12:39:35 starting...
2018/03/10 12:39:37 diff ns: 24344
chuqq@chuqq:~/work/codeeveryday/c/20180209_futex$ ./test_channel_delay
2018/03/10 12:39:40 starting...
2018/03/10 12:39:42 diff ns: 33013
chuqq@chuqq:~/work/codeeveryday/c/20180209_futex$ ./test_channel_delay
2018/03/10 12:39:45 starting...
2018/03/10 12:39:47 diff ns: 12233
chuqq@chuqq:~/work/codeeveryday/c/20180209_futex$ ./test_channel_delay
2018/03/10 12:39:50 starting...
2018/03/10 12:39:52 diff ns: 13538
*/
// 结论：大约20us左右，和pthread_cond_x以及futex差不多
