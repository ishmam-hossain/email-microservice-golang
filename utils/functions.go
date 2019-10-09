package functions

import (
	"github.com/bitly/go-nsq"
)


func nsqDataWriter() (interface{}) {
	nsqConfig := nsq.NewConfig()
	nsqWriter, err := nsq.NewProducer("127.0.0.1:4150", nsqConfig)
	if err != nil {
		return err
	}
	return nsqWriter
}

// NsqPublish publishes data to queue
func NsqPublish(topicName string, data []byte) error {
	nsqConfig := nsq.NewConfig()
	nsqWriter, err := nsq.NewProducer("127.0.0.1:4150", nsqConfig)
	if err != nil {
		return err
	}

	return nsqWriter.Publish(topicName, data)
}