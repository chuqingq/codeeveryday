package log

import (
	"os"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestLogLevel(t *testing.T) {
	logger := New()
	go func() {
		time.Sleep(time.Second * 5)
		logger.SetLevel(logrus.ErrorLevel)
	}()
	for {
		logger.Debugf("this is debug log")
		logger.Infof("this is info log")
		logger.Warnf("this is warn log")
		logger.Errorf("this is error log")

		time.Sleep(time.Second)
	}
}

func TestAppendOutput(t *testing.T) {
	logger := New()

	// 打开文件，如果不存在则创建
	file, err := os.OpenFile("/Users/chuqq/data/temp/projects/codeeveryday/golang/20230610-go-log/test.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	assert.Equal(t, nil, err)
	defer file.Close()

	go func() {
		time.Sleep(time.Second * 5)
		AppendOutput(logger, file)
	}()
	for {
		logger.Debugf("this is debug log")
		logger.Infof("this is info log")
		logger.Warnf("this is warn log")
		logger.Errorf("this is error log")

		time.Sleep(time.Second)
	}
}
