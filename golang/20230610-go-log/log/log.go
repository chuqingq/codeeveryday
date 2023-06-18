package log

import (
	"fmt"
	"io"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Options struct {
	FileName    string // 日志文件名，不包含路径
	MaxSizeInMB int    // 日志文件大小，单位MB，>=1
	MaxBackups  int    // 日志文件最大备份数，>=1
	Formatter   logrus.Formatter
}

func New() *logrus.Logger {
	return NewWithOptions(&Options{})
}

func NewWithOptions(options *Options) *logrus.Logger {
	filename := "test.log"
	if options.FileName != "" {
		filename = options.FileName
	}

	maxsize := 1
	if options.MaxSizeInMB > 1 {
		maxsize = options.MaxSizeInMB
	}

	maxbackups := 1
	if options.MaxBackups > 1 {
		maxbackups = options.MaxBackups
	}

	var formatter logrus.Formatter = &myFormatter{}
	if options.Formatter != nil {
		formatter = options.Formatter
	}

	// lumberjack logger作为logrus的输出
	output := &lumberjack.Logger{
		Filename:   filepath.Join("/dev/shm/", filename), // in memory
		MaxSize:    maxsize,                              // megabytes
		MaxBackups: maxbackups,                           // reserve 1 backup
		// MaxAge:     28, //days
	}

	logger := &logrus.Logger{
		Out: output,
		// Formatter: &logrus.TextFormatter{},
		Formatter: formatter,
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.DebugLevel,
	}
	logger.SetReportCaller(true)

	return logger
}

// SetLogLevel 设置日志级别
// logger.SetLevel(logrus.DebugLevel)

// AppendOutput 添加日志输出
func AppendOutput(logger *logrus.Logger, output io.Writer) {
	logger.SetOutput(&logOutput{cur: logger.Out, next: output})
}

type logOutput struct {
	cur  io.Writer
	next io.Writer
}

func (o *logOutput) Write(p []byte) (n int, err error) {
	o.cur.Write(p)
	return o.next.Write(p)
}

// myFormatter 自定义日志格式
type myFormatter struct {
}

// Format 格式化日志
func (f *myFormatter) Format(e *logrus.Entry) ([]byte, error) {
	return []byte(fmt.Sprintf("%s %5.5s [%s:%v %s] %s\n",
		e.Time.Format("01/02 15:04:05.000000"),
		e.Level.String(),
		splitAndGetLast(e.Caller.File, "/"),
		e.Caller.Line, splitAndGetLast(e.Caller.Function, "."),
		e.Message)), nil
}

// splitAndGetLast 分割字符串并返回最后一个元素
func splitAndGetLast(str string, sep string) string {
	slice := strings.Split(str, sep)
	return slice[len(slice)-1]
}
