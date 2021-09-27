package util

import "testing"

func TestDebug(t *testing.T) {
	// default debug
	D().Printf("debug log should be displayed")
	I().Printf("info log should be displayed")
}

func TestWarn(t *testing.T) {
	SetLogLevel(LogLevelInfo)
	D().Printf("!!!! debug log should not be displayed")
	I().Printf("info log should be displayed")
	SetLogLevel(LogLevelDebug)
}
