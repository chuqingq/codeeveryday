package util

import (
	"log"
	"os"
	"sync"
)

// 循环输出方式

type circularLogger struct {
	files   [2]*os.File
	curFile int
	*log.Logger
	num    int32 // 打印的条数
	numMax int32
	sync.Mutex
}

func newCircularLogger() *circularLogger {
	file0, err := os.Create(os.Args[0] + ".log.0")
	if err != nil {
		log.Printf("create log file error: %v", err)
		return nil
	}
	file1, err := os.Create(os.Args[0] + ".log.1")
	if err != nil {
		log.Printf("create log file error: %v", err)
		return nil
	}
	return &circularLogger{
		files:  [2]*os.File{file0, file1},
		Logger: log.New(file0, "", log.Flags()|logFlags),
		numMax: 10000,
	}
}

func (c *circularLogger) Printf(format string, v ...interface{}) {
	c.Lock()
	c.num += 1
	if c.num > c.numMax {
		c.files[1-c.curFile].Seek(0, 0)
		c.files[1-c.curFile].Truncate(0)
		c.Logger.SetOutput(c.files[1-c.curFile])
		c.curFile = 1 - c.curFile
		c.num = 0
	}
	c.Unlock()
	c.Logger.Printf(format, v...)
}
