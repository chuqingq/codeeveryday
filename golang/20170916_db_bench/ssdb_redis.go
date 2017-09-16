package main

import (
	"flag"
	"log"
	"math/rand"
	"strconv"
	"time"

	"gopkg.in/redis.v5"
)

func main() {
	count := flag.Int("n", 100000, "count")
	flag.Parse()

	i := 0
	phase := "prepare"
	// stats
	go func() {
		ticker := time.NewTicker(time.Second)
		acc := i
		for {
			select {
			case <-ticker.C:
				log.Printf("%v acc: %v, i: %v", phase, i-acc, i)
				acc = i
				if phase == "run" && i == *count {
					return
				}
			}
		}
	}()

	// setup
	keys := make([]string, *count, *count)
	for i = 0; i < *count; i++ {
		keys[i] = strconv.Itoa(rand.Int())
	}

	// run
	log.Printf("running...")
	phase = "run"

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:8888",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	log.Printf("pong: %v, err: %v", pong, err)

	value := "123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123"
	start := time.Now()
	for i = 0; i < *count; i++ {
		err := client.Set(keys[i], value, 0).Err()
		if err != nil {
			panic(err)
		}
	}
	elapsed := time.Since(start).Seconds()
	log.Printf("%v", float64(*count)/elapsed)

	select {}

	client.Close()
}

/*
在191.0.0.22上验证，./ssdb_redis -n 1000000，结果是
2017/09/15 16:51:04 run acc: 13926, i: 872595
2017/09/15 16:51:05 run acc: 12501, i: 885096
2017/09/15 16:51:06 run acc: 12044, i: 897140
2017/09/15 16:51:07 run acc: 12690, i: 909830
2017/09/15 16:51:08 run acc: 14482, i: 924312
2017/09/15 16:51:09 run acc: 13770, i: 938082
2017/09/15 16:51:10 run acc: 14963, i: 953045
2017/09/15 16:51:11 run acc: 13658, i: 966703
2017/09/15 16:51:12 run acc: 13212, i: 979915
2017/09/15 16:51:13 run acc: 13693, i: 993608
2017/09/15 16:51:14 14222.948421239644
2017/09/15 16:51:14 run acc: 6392, i: 1000000

结果比较稳定，但是比C++的ssdb-bench慢很多，后者的速度约为：set: 3.5w, get: 6w, delete: 3.6w
。。。
finished: 980000
finished: 990000
finished: 1000000
qps: 35461, time: 28.199 s
========== get ==========
finished: 10000
。。。
finished: 970000
finished: 980000
finished: 990000
finished: 1000000
qps: 60442, time: 16.545 s
========== del ==========
finished: 10000
。。。
finished: 960000
finished: 970000
finished: 980000
finished: 990000
finished: 1000000
qps: 36715, time: 27.236 s

原因可能是ssdb-bench利用了多核特性，而ssdb_redis是单线程、单连接。
*/
