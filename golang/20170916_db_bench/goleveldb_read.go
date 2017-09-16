package main

import (
	// "encoding/binary"
	"flag"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/syndtr/goleveldb/leveldb"
)

func main() {
	count := flag.Int("n", 100000, "count")
	flag.Parse()

	rand.Seed(time.Now().Unix())

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
	log.Printf("seting up...")
	keys := make([][]byte, *count, *count)
	for i = 0; i < *count; i++ {
		// a := rand.Int63()
		// /*ret := */ binary.PutVarint(keys[i][:], a)
		// // log.Printf("ret: %v", ret)
		// b := rand.Int63()
		// binary.PutVarint(keys[i][6:], b)
		keys[i] = []byte(strconv.Itoa(rand.Int()))
	}

	// run
	log.Printf("running...")
	phase = "run"
	db, err := leveldb.OpenFile("./db", nil)
	if err != nil {
		log.Fatalf("leveldb open error: %v", err)
	}
	start := time.Now()
	for i = 0; i < *count; i++ {
		// db.Put(keys[i][:], keys[i][:], nil)
		db.Get(keys[i][:], nil)
	}
	elapsed := time.Since(start).Seconds()
	log.Printf("%v", float64(*count)/elapsed)

	select {}

	db.Close()
}

/*
在191.0.0.22上验证：# ./goleveldb_write -n 1000000
单线程，经常在800左右：
2017/09/16 13:44:00 run acc: 875, i: 92942
2017/09/16 13:44:01 run acc: 867, i: 93809
2017/09/16 13:44:02 run acc: 874, i: 94683
2017/09/16 13:44:03 run acc: 876, i: 95559
2017/09/16 13:44:04 run acc: 878, i: 96437
2017/09/16 13:44:05 2346.684331029464
2017/09/16 13:44:05 run acc: 3563, i: 100000
*/
