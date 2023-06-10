package main

import "test_go/log"

func main() {
	for {
		log.Logger().Debugf("this is debug log")
		log.Logger().Infof("this is info log")
		log.Logger().Warnf("this is warn log")
		log.Logger().Errorf("this is error log")
	}
}
