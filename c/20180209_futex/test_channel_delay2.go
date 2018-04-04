package main

import (
	"log"
	"time"
)

var starttime int64

func main() {
	// log.Printf("starting...")
	ch := make(chan bool, 1)
	go func() {
		time.Sleep(1 * time.Second)
		starttime = time.Now().UnixNano()
		ch <- true
		// log.Printf("diff ns: %v", time.Now().UnixNano()-starttime)
	}()

	// time.Sleep(1 * time.Second)
	// log.Printf("%v", time.Now().UnixNano())
	// starttime = time.Now().UnixNano()
	<-ch
	stoptime := time.Now().UnixNano()
	log.Printf("elapsed %v ns", stoptime-starttime)
	//time.Sleep(2 * time.Second)
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

/*
换种写法：子协程给主协程发消息，保存stoptime后再打印。
chuqq@chuqq-hp:~/temp/codeeveryday/c/20180209_futex$ go build test_channel_delay2.go
chuqq@chuqq-hp:~/temp/codeeveryday/c/20180209_futex$ ./test_channel_delay2 
2018/04/04 09:44:53 elapsed 2789 ns
chuqq@chuqq-hp:~/temp/codeeveryday/c/20180209_futex$ ./test_channel_delay2 
2018/04/04 09:44:55 elapsed 2681 ns
chuqq@chuqq-hp:~/temp/codeeveryday/c/20180209_futex$ ./test_channel_delay2 
2018/04/04 09:44:57 elapsed 2619 ns
chuqq@chuqq-hp:~/temp/codeeveryday/c/20180209_futex$ ./test_channel_delay2 
2018/04/04 09:44:58 elapsed 2659 ns
chuqq@chuqq-hp:~/temp/codeeveryday/c/20180209_futex$ ./test_channel_delay2 
2018/04/04 09:45:00 elapsed 3282 ns

结论：这种方式比前面的方式快很多。
*/
