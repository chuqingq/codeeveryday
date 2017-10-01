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
		deviceid := rand.Int()
		requestid := rand.Int()
		keys[i] = []byte(strconv.Itoa(deviceid) + "-" + strconv.Itoa(requestid))
	}

	// run
	log.Printf("running...")
	phase = "run"
	db, err := leveldb.OpenFile("./db", nil)
	if err != nil {
		log.Fatalf("leveldb open error: %v", err)
	}
	start := time.Now()
	// content := []byte(":58671272327d0c220df0f9bfb0e36fdf224d4be1b701435c6c58e7064c72ba0fb4fae227cfe868780e1adb3315e9a8e922cffacfe3b9a5c685f0e64209c6bee240d903bed9f6be314831720efb794e5e901f019182485afa8a93198321f84eb0bfb539506fb9f4a9fcf42f075265c702ce7ecf97d9b64d4fa12eee25767534e1979cee6a1deec2c46e7592430c9a31eb862d9fb302b067831dd2447a0354c698c4928be0eb0896f077e72f858fdb22ba46ed7ce2dfb5191d7f090a7b830f81d4d5298b21c63e4f0e9ef4200cb0be2cf3f5fb9f6484f0fee3e9752e37b9c9c0e787afb4cfc5956dfb2ddbf44ddf89dc7d94df8648345e5fbc3f25ac57c23a441cf0585d9d4653cafdb7a702e0a03cad63")
	for i = 0; i < *count; i++ {
		db.Put(keys[i][:], keys[i][:], nil)
	}
	elapsed := time.Since(start).Seconds()
	log.Printf("%v", float64(*count)/elapsed)

	select {}

	db.Close()
}
