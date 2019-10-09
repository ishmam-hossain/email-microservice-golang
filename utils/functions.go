package functions

import (
	"log"

	"github.com/bitly/go-nsq"
)

var nsqWriter *nsq.Producer

// InitNSQ nsq initialization
func InitNSQ() {
	var err error
	nsqConfig := nsq.NewConfig()
	nsqWriter, err = nsq.NewProducer("127.0.0.1:4150", nsqConfig)

	if err != nil {
		log.Panic(err)
	}
}

// NsqPublish publishes data to queue
func NsqPublish(topicName string, data []byte) error {
	return nsqWriter.Publish(topicName, data)
}

// KillNSQ kills the connection 
func KillNSQ() {
	nsqWriter.Stop()
	log.Fatal("killed!")
}