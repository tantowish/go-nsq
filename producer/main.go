package main

import (
	"log"
	"strconv"
	"time"

	"github.com/nsqio/go-nsq"
)

func main() {
	nsqtopic := "learningnsq"

	config := nsq.NewConfig()
	w, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		log.Panic("Error while connecting producer to nsqd")
	}

	for i := 0; i < 50; i++ {
		time.Sleep(1 * time.Second)
		err := w.Publish(nsqtopic, []byte(strconv.Itoa(i)))
		if err != nil {
			log.Panic("Publishing of message on topic : " + nsqtopic + " failed")
		}
	}

	w.Stop()
}
