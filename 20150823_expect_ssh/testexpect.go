package main

import (
	"flag"
	// "fmt"
	"github.com/jamesharr/expect"
	"log"
	"time"
)

func main() {
	flag.Parse()

	host := flag.Arg(0)
	if host == "" {
		log.Fatalf("please specify a host to connect")
	}

	ssh, err := expect.Spawn("ssh", host)
	if err != nil {
		log.Fatalf("spawn error: %v\n", err)
	}
	ssh.SetTimeout(10 * time.Second)
	match, err := ssh.Expect(`assword:`)
	if err != nil {
		log.Fatalf("expect password error: %v\n", err)
	}
	log.Printf("%s", match)

	// get password
	password := getPasswordByUserAndHost(host)
	if password == "" {
		// fmt.Println("password: ")
		// TODO
	}

	ssh.SendMasked(password)
	ssh.SendMasked("\n")

	match, err = ssh.Expect(`\$`)
	if err != nil {
		log.Fatalf("expect prompt error: %v\n", err)
	}
	log.Printf("%s", match)

	// // set prompt
	// ssh.SendLn("export PS1=" + host + " $ ")
	// _, err = ssh.Expect(`(?m)^.*\$`)
	// if err != nil {
	// 	log.Fatalf("expect export PS1 error: %v\n", err)
	// }

	// match, err = ssh.Expect(`(?m)^.*\$`)
	// if err != nil {
	// 	log.Fatalf("1 expect error: %v\n", err)
	// }
	// log.Printf("logined: %v\n", match)

	ssh.SendLn("ls")
	match, err = ssh.Expect(`\$`)
	log.Printf("output: %v\n", match)
}

func getPasswordByUserAndHost(host string) string {
	return "chuqingq"
}
