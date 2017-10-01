package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/jsimonetti/berkeleydb"
)

func main() {
	count := flag.Int("n", 100000, "count")
	flag.Parse()

	var err error

	db, err := berkeleydb.NewDB()
	if err != nil {
		fmt.Printf("Unexpected failure of CreateDB %s\n", err)
	}

	err = db.Open("./test.db", berkeleydb.DbHash, berkeleydb.DbCreate)
	if err != nil {
		fmt.Printf("Could not open test_db.db. Error code %s", err)
		return
	}
	defer db.Close()

	// err = db.Put("key", "value")
	// if err != nil {
	// 	fmt.Printf("Expected clean Put: %s\n", err)
	// }

	// value, err := db.Get("key")
	// if err != nil {
	// 	fmt.Printf("Unexpected error in Get: %s\n", err)
	// 	return
	// }
	// fmt.Printf("value: %s\n", value)

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

	value := "123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123"
	start := time.Now()
	for i = 0; i < *count; i++ {
		// m[strconv.Itoa(rand.Int())] = node{DeviceId: strconv.Itoa(rand.Int())}
		// m[strconv.Itoa(rand.Int())] = make([]byte, 400, 400)
		err = db.Put(strconv.Itoa(rand.Int()), value)
		if err != nil {
			log.Printf("db.put error: %v", err)
		}
	}
	elapsed := time.Since(start).Seconds()
	log.Printf("%v", float64(*count)/elapsed)

	select {}
}
