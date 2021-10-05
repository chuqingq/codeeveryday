package util

import (
	"log"
	"net"
	"os"
)

// 日志级别
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

type printfInterface interface {
	Printf(fmt string, v ...interface{})
}

func D() printfInterface {
	if curLevel <= LogLevelDebug {
		return logger
	}
	return empty
}

func I() printfInterface {
	if curLevel <= LogLevelInfo {
		return logger
	}
	return empty
}

func E() printfInterface {
	if curLevel <= LogLevelError {
		return logger
	}
	return empty
}

// 多个输出串联

var empty *myLogger
var logger *myLogger

type PrintfI interface {
	Printf(format string, v ...interface{})
}

type myLogger struct {
	loggers []PrintfI
}

func (m *myLogger) Printf(fmt string, v ...interface{}) {
	for _, l := range m.loggers {
		l.Printf(fmt, v...)
	}
}

func (m *myLogger) AppendLogger(logger ...PrintfI) {
	if m.loggers == nil {
		m.loggers = make([]PrintfI, 0, 2)
	}
	m.loggers = append(m.loggers, logger...)
}

// 输出方式

var stdLogger *log.Logger

func newStdLogger() *log.Logger {
	return log.New(os.Stdout, "", log.Flags()|logFlags)
}

var udpLogger *log.Logger

func newUdpLogger() *log.Logger {
	conn, err := net.Dial("udp", "127.0.0.1:61234")
	if err != nil {
		log.Printf("log udp dial error: %v", err)
		return nil
	}
	return log.New(conn, "", log.Flags()|logFlags)
}

const logFlags = log.Lshortfile | log.Lmicroseconds

// init

func init() {
	// 输出方式
	stdLogger = newStdLogger()
	udpLogger = newUdpLogger()
	// 级别：输出还是不输出
	empty = &myLogger{}
	empty.AppendLogger(udpLogger)
	logger = &myLogger{}
	logger.AppendLogger(stdLogger, udpLogger)
}
