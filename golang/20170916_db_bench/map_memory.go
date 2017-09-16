package main

import (
	"flag"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	count := flag.Int("n", 100000, "count")
	flag.Parse()

	rand.Seed(time.Now().Unix())

	i := 0
	// stats
	go func() {
		ticker := time.NewTicker(time.Second)
		acc := i
		for {
			select {
			case <-ticker.C:
				log.Printf("acc: %v, i: %v", i-acc, i)
				acc = i
				if i == *count {
					return
				}
			}
		}
	}()

	type node struct {
		DeviceId     string
		Hash         int
		LastSendTime int64
		IsOnline     bool
		HasMsg       bool
	}
	// m := make(map[string]node, *count)
	m := make(map[string][]byte, *count)
	start := time.Now()
	for i = 0; i < *count; i++ {
		// m[strconv.Itoa(rand.Int())] = node{DeviceId: strconv.Itoa(rand.Int())}
		m[strconv.Itoa(rand.Int())] = make([]byte, 400, 400)
	}
	elapsed := time.Since(start).Seconds()
	log.Printf("%v", float64(*count)/elapsed)

	select {}
}

/*
场景1：没有msglist，只保存一个node
PID USER      PR  NI  VIRT  RES  SHR S   %CPU %MEM    TIME+  COMMAND
 2957 root      20   0 1670m 1.6g  732 S      0 10.2   0:15.13 map_memory
1000w设备占用大约1.6G内存

场景2：保留一条消息，平均大小400B
3004 root      20   0 5411m 1.3g  748 S      0  7.8   0:11.75 map_memory
1000w设备占用大约5G内存

结论：
按照这个规格，每个logic 8C 16G，保存1000w设备，每个设备保存一条待推送消息，大约占用5G内存
*/
