package main

import (
//	"encoding/json"
//	"io/ioutil"
	"log"
	"net/http"
//	"strconv"
	"time"
)

func main() {
/*	msg := make(map[string]map[string]interface{})
	e := map[string]interface{}{
		"integer": 12345,
		"float":   1.2345,
		"string":  "12345",
	}
	for i := 0; i < 200; i++ {
		msg[strconv.Itoa(i)] = e
	}
	bytes, err := json.Marshal(msg)
	if err != nil {
		log.Fatalf("json marshal error: %s\n", err.Error())
	}
	err = ioutil.WriteFile("json", bytes, 0777)
	if err != nil {
		log.Fatalf("write file error: %s\n", err.Error())
	}*/

	fileHandler := http.FileServer(http.Dir("."))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(90 * time.Millisecond)
		fileHandler.ServeHTTP(w, r)
	})

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatalf("listen error: %s\n", err.Error())
	}
}

func handle(w http.ResponseWriter, r *http.Request) {

}
