package main

import (
	//"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

type Person struct {
	K int
	V []map[string]int
}

func main() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("bench").C("bench")
	/*
	   err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
	                  &Person{"Cla", "+55 53 8402 8510"})
	   if err != nil {
	           log.Fatal(err)
	   }
	*/

	startTime := time.Now()

	totalErr := 0
	for i := 0; i < 100000; i++ {
		result := Person{}
		// err = c.Find(bson.M{"k": i}).Hint("k").One(&result)
		err = c.Find(bson.M{"k": i}).One(&result)
		if err != nil {
			log.Printf("error: %v", err)
			totalErr += 1
		}
	}

	log.Printf("duration: %v, totalErr: %v", time.Now().Sub(startTime), totalErr)
}

/*
1.7亿条数据，占用内存60+GB

客户端循环：
1w: 1.82s
10w: 18.3s

server side code execution：
大约是一半的耗时。
*/
