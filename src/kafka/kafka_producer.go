package kafka

import (
	"github.com/IBM/sarama"
)

var (
	kafkaBrokers = []string{"localhost:9092"}
)

func ProduceToKafka(topic, message string) (int32, int64, error) {
	producer, err := sarama.NewSyncProducer(kafkaBrokers, nil)
	if err != nil {
		return 0, 0, err
	}
	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		return 0, 0, err
	}

	return partition, offset, nil
}
