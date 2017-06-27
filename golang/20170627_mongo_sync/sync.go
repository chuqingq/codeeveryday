package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const CONFIG_FILE = "./sync.conf"

func main() {
	config, err := InitConfig(CONFIG_FILE)
	if err != nil {
		log.Fatalf("InitConfig error: %v\n", err)
	}
	log.Printf("config: %+v", config)

	config.Start()
	select {}
}

type ConnectInfo struct {
	Addrs    string `json:"addrs"`
	SSL      bool   `json:"ssl"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Config struct {
	Src     ConnectInfo `json:"src"`
	Dst     ConnectInfo `json:"dst"`
	BufSize int         `json:"bufsize"`
	Ns      []string    `json:"ns"`
	Ts      uint64      `json:"ts"`

	SrcSess    *mgo.Session `json:"-"`
	DstSess    *mgo.Session `json:"-"`
	OplogChan  chan Oplog   `json:"-"`
	OpsToApply []Oplog      `json:"-"`
}

func InitConfig(filename string) (*Config, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(b, &config)
	if err != nil {
		return nil, err
	}

	config.OplogChan = make(chan Oplog, config.BufSize)
	config.OpsToApply = make([]Oplog, 0, config.BufSize)

	return &config, nil
}

func (config *Config) WriteFile() error {
	b, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(CONFIG_FILE, b, 0666)
	if err != nil {
		return err
	}
	return nil
}

func (config *Config) Start() {
	go config.startSrc()
	go config.startDst()
}

type Oplog struct {
	Timestamp bson.MongoTimestamp `bson:"ts"`
	HistoryID int64               `bson:"h"`
	Version   int                 `bson:"v"`
	Operation string              `bson:"op"`
	Namespace string              `bson:"ns"`
	Object    bson.D              `bson:"o"`
	Query     bson.D              `bson:"o2"`
}

func (config *Config) startSrc() {
	var err error
	for {
		config.SrcSess, err = connectMongo(&config.Src)
		if err != nil {
			log.Printf("connectSrc error: %v\n", err)
			time.Sleep(5 * time.Second)
			continue
		}
		log.Printf("connectMongo(src) success\n")

		err = config.readOplog()
		if err != nil {
			log.Printf("readOplog error: %v\n", err)
			time.Sleep(5 * time.Second)
			continue
		}
	}
}

func (config *Config) startDst() {
	var err error
	for {
		config.DstSess, err = connectMongo(&config.Dst)
		if err != nil {
			log.Printf("connectSrc error: %v\n", err)
			time.Sleep(5 * time.Second)
			continue
		}
		log.Printf("connectMongo(dst) success\n")

		err = config.writeOplog()
		if err != nil {
			log.Printf("writeOplog error: %v\n", err)
			time.Sleep(5 * time.Second)
			continue
		}
	}
}

func connectMongo(connectInfo *ConnectInfo) (*mgo.Session, error) {
	password := connectInfo.Password

	dialInfo := &mgo.DialInfo{
		Addrs:    strings.Split(connectInfo.Addrs, ","),
		Username: connectInfo.Username,
		Password: password,
		Timeout:  10 * time.Second,
	}

	return mgo.DialWithInfo(dialInfo)
}

func (config *Config) readOplog() error {
	log.Printf("readOplog()")
	oplog := config.SrcSess.DB("local").C("oplog.rs")

	dbRegExs := []bson.RegEx{}
	for _, dbname := range config.Ns {
		dbRegExs = append(dbRegExs, bson.RegEx{Pattern: dbname})
	}
	log.Printf("dbRegExs: %+v", dbRegExs)

	iter := oplog.Find(bson.M{"ts": bson.M{"$gt": bson.MongoTimestamp(config.Ts)}, "ns": bson.M{"$in": dbRegExs}}).Tail(1 * time.Second)
	defer iter.Close()

	var oplogEntry Oplog
	for {
		for iter.Next(&oplogEntry) {
			if oplogEntry.Operation == "n" {
				// log.Printf("skipping no-op for namespace `%v`", oplogEntry.Namespace)
				continue
			}

			config.OplogChan <- oplogEntry
		}

		err := iter.Err()
		if err != nil {
			iter.Close()
			return err
		}

		if iter.Timeout() {
			continue
		}
	}
}

type ApplyOpsResponse struct {
	Ok     bool   `bson:"ok"`
	ErrMsg string `bson:"errmsg"`
}

func (config *Config) writeOplog() error {
	log.Printf("writeOplog()")

	for {
		select {
		case oplogEntry := <-config.OplogChan:
			config.OpsToApply = append(config.OpsToApply, oplogEntry)
			if len(config.OpsToApply) == cap(config.OpsToApply) {
				config.applyOps()
			}
		case <-time.After(1 * time.Second):
			config.applyOps()
		}
	}
}

func (config *Config) applyOps() error {
	if len(config.OpsToApply) == 0 {
		return nil
	}

	log.Printf("%v opCount: %v", time.Now(), len(config.OpsToApply))

	var applyOpsResponse ApplyOpsResponse
	err := config.DstSess.Run(bson.M{"applyOps": config.OpsToApply}, &applyOpsResponse)
	if err != nil {
		return err
	}

	if !applyOpsResponse.Ok {
		return fmt.Errorf("server gave error applying ops: %v", applyOpsResponse.ErrMsg)
	}

	last := config.OpsToApply[len(config.OpsToApply)-1]
	config.Ts = uint64(last.Timestamp)
	err = config.WriteFile()
	if err != nil {
		log.Printf("config.WriteFile error: %v", err)
		return err
	}

	config.OpsToApply = config.OpsToApply[:0:cap(config.OpsToApply)]

	return nil
}

