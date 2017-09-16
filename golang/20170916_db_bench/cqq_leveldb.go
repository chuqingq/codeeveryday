package main

import (
	"encoding/binary"
	"flag"
	"log"
	"math/rand"
	"time"

	"github.com/syndtr/goleveldb/leveldb"
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
	log.Printf("seting up...")
	keys := make([][16]byte, *count, *count)
	for i = 0; i < *count; i++ {
		a := rand.Int63()
		/*ret := */ binary.PutVarint(keys[i][:], a)
		// log.Printf("ret: %v", ret)
		b := rand.Int63()
		binary.PutVarint(keys[i][6:], b)
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
		db.Put(keys[i][:], keys[i][:], nil)
	}
	elapsed := time.Since(start).Seconds()
	log.Printf("%v", float64(*count)/elapsed)

	select {}

	db.Close()
}
