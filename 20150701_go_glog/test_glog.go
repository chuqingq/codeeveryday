package main

import (
	"flag"
	// _ "fmt"
	"github.com/golang/glog"
	// _ "os"
)

// 避免没有引用fmt的编译错误
// var _ = fmt.Println

func main() {
	//初始化命令行参数
	flag.Parse()

	glog.Info("hello, glog")
	glog.Warning("warning glog")
	glog.Error("error glog")

	glog.Infof("info %d", 1)
	glog.Warningf("warning %d", 2)
	glog.Errorf("error %d", 3)

	glog.V(3).Infoln("info with v 3")
	glog.V(2).Infoln("info with v 2")
	glog.V(1).Infoln("info with v 1")
	glog.V(0).Infoln("info with v 0")

	// 退出时调用，确保日志写入文件中
	glog.Flush()
}
