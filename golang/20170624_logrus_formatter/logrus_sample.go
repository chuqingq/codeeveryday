package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

type XLogFormatter struct {
}

func (log *XLogFormatter) Format(entry *log.Entry) ([]byte, error) {
	b := &bytes.Buffer{}

	timestamp := time.Now().Format("2006-01-02 15:04:05.999 ")
	b.WriteString(timestamp)

	levelText := strings.ToUpper(entry.Level.String())[0:4]
	b.WriteString(levelText)

	fmt.Fprintf(b, " %s ", entry.Message)

	for k, v := range entry.Data {
		fmt.Fprintf(b, "%v=%v,", k, v)
	}

	b.WriteByte('\n')
	return b.Bytes(), nil
}

func main() {
	file, err := os.OpenFile("./logrus_sample.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return
	}
	log.SetOutput(file)
	log.SetFormatter(&XLogFormatter{})

	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
}

