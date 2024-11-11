package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"

	"github.com/nsqio/go-nsq"
)

type handler struct{}

func (*handler) HandleMessage(message *nsq.Message) error {
	bodyStr := string(message.Body)

	value, err := strconv.Atoi(bodyStr)
	if err != nil {
		return fmt.Errorf("failed to convert message body to int: %v", err)
	}

	if(value%2 == 0){
		fmt.Println(bodyStr+ " is even")
	} else {
		fmt.Println(bodyStr+ " is odd")
	}
	return nil
}

func main() {
	nsqtopic := "learningnsq"

	wg := sync.WaitGroup{}
	wg.Add(1)

	config := nsq.NewConfig()
	q, err := nsq.NewConsumer(nsqtopic, "channel", config)
	if err != nil {
		log.Panic("Error while connecting consumer to nsqd")
	}

	q.AddHandler(&handler{})

	err = q.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		log.Panic("Could not connect to nsqd")
	}

	wg.Wait()
}
