package main

import (
	"test_go/log"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	logger := log.New()
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
