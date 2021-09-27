package util

import (
	"log"
	"os"
)

func init() {
	log.SetFlags(log.Flags() | log.Lshortfile | log.Lmicroseconds)
	logger = log.New(os.Stderr, "", log.Flags())
}

const (
	LogLevelDebug = 0 + iota
	LogLevelInfo
	LogLevelWarn
	LogLevelError
	LogLevelFatal
	LogLevelOff
)

var curLevel int = LogLevelDebug

func SetLogLevel(level int) {
	if level < LogLevelDebug {
		curLevel = LogLevelDebug
	} else if curLevel > LogLevelOff {
		curLevel = LogLevelOff
	} else {
		curLevel = level
	}
}

var empty emptyLogger

var logger *log.Logger

type printfInterface interface {
	Printf(fmt string, v ...interface{})
}

type emptyLogger struct {
}

func (e *emptyLogger) Printf(fmt string, v ...interface{}) {
}

func D() printfInterface {
	if curLevel <= LogLevelDebug {
		return logger
	}
	return &empty
}

func I() printfInterface {
	if curLevel <= LogLevelInfo {
		return logger
	}
	return &empty
}

func E() printfInterface {
	if curLevel <= LogLevelError {
		return logger
	}
	return &empty
}
