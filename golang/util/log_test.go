package util

import (
	"log"
	"testing"
)

func TestLogDebug(t *testing.T) {
	log.Printf("test log debug")
	// default debug
	D().Printf("debug log should be displayed")
	I().Printf("info log should be displayed")
}

func TestLogWarn(t *testing.T) {
	log.Printf("test log warn")
	SetLogLevel(LogLevelInfo)
	D().Printf("!!!! debug log should not be displayed")
	I().Printf("info log should be displayed")
	SetLogLevel(LogLevelDebug)
}
