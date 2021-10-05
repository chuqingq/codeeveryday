package util

import (
	"log"
	"testing"
)

func TestLogCircular(t *testing.T) {
	log.Printf("test circurlar")
	// logger := newCircularLogger()
	// for i := 0; i < 15000; i++ {
	for i := 0; i < 21000; i++ {
		cirLogger.Printf("%v", i)
	}
	D().Printf("debug log should be displayed")
	I().Printf("info log should be displayed")
}
