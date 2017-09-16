package main

import (
	// "encoding/binary"
	// "flag"
	"log"
	// "math/rand"
	"time"

	"github.com/syndtr/goleveldb/leveldb"
)

func main() {
	// stats
	count := 0
	go func() {
		ticker := time.NewTicker(time.Second)
		for {
			select {
			case <-ticker.C:
				log.Printf("count: %v", count)
			}
		}
	}()

	db, err := leveldb.OpenFile("./db", nil)
	if err != nil {
		log.Fatalf("leveldb open error: %v", err)
	}

	iter := db.NewIterator(nil, nil)
	start := time.Now()
	for iter.Next() {
		// Remember that the contents of the returned slice should not be modified, and
		// only valid until the next call to Next.
		count++
	}
	elapsed := time.Since(start).Seconds()
	log.Printf("%v", float64(*count)/elapsed)

	iter.Release()
	err = iter.Error()
	if err != nil {
		log.Fatalf("iter error: %v", err)
	}

	select {}

	db.Close()
}
