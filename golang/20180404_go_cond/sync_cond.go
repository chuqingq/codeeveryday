package main

import (
	"sync"
	"time"
	"log"
)

func main() {
	locker := new(sync.Mutex)
	cond := sync.NewCond(locker)

	var start int64
	go func() {
		time.Sleep(1 * time.Second)
		start = time.Now().UnixNano()
		cond.Signal()
	}()

	cond.L.Lock()
	cond.Wait()
	cond.L.Unlock()

	stop := time.Now().UnixNano()
	log.Printf("elapsed %v ns", stop-start)
}

/*
env: lubuntu@hp
chuqq@chuqq-hp:~/temp/codeeveryday/golang/20180404_go_cond$ ./sync_cond 
2018/04/04 09:39:11 elapsed 2811 ns
chuqq@chuqq-hp:~/temp/codeeveryday/golang/20180404_go_cond$ ./sync_cond 
2018/04/04 09:39:13 elapsed 2648 ns
chuqq@chuqq-hp:~/temp/codeeveryday/golang/20180404_go_cond$ ./sync_cond 
2018/04/04 09:39:14 elapsed 2664 ns
chuqq@chuqq-hp:~/temp/codeeveryday/golang/20180404_go_cond$ ./sync_cond 
2018/04/04 09:39:16 elapsed 2581 ns
chuqq@chuqq-hp:~/temp/codeeveryday/golang/20180404_go_cond$ ./sync_cond 
2018/04/04 09:39:17 elapsed 4150 ns
chuqq@chuqq-hp:~/temp/codeeveryday/golang/20180404_go_cond$ ./sync_cond 
2018/04/04 09:39:21 elapsed 2706 ns

结论：
pthread_cond和futex大约38us；
sync.Cond的方式和channel方式差不多，大约2.7us。

*/
