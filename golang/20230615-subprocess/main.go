package main

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/chuqingq/go-util"
)

func main() {
	// p, err := NewSubProcess("./child/child")
	p, err := NewSubProcess("echo", `{"name": 789}`)
	if err != nil {
		panic(err)
	}
	err = p.Start()
	if err != nil {
		panic(err)
	}

	go io.Copy(os.Stderr, p.Stderr)

	time.Sleep(time.Second * 4)
	m := util.NewMessage()
	m.Set("name", 123)
	log.Printf("==== parent send to child msg: %v", m)
	log.Printf("==== parent child is alive: %v", p.IsAlive())
	err = p.Send(m)
	if err != nil {
		log.Printf("==== parent child is alive: %v", p.IsAlive())
		log.Fatalf("==== parent send error: %v", err)
	}
	m2, err := p.Recv()
	if err != nil {
		log.Fatalf("====recv error: %v", err)
	}
	log.Printf("==== parent recv: %v", m2)

	time.Sleep(time.Second * 5)
	log.Printf("==== parent stop child")
	p.Stop()
	log.Printf("==== parent after stop child")
	time.Sleep(time.Second * 5)
}
