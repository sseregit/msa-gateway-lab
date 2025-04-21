package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"go-gateway/config"
	"log"
)

const (
	_allAcks = "all"
)

type Producer struct {
	cfg      config.Producer
	producer *kafka.Producer
}

func NewProducer(cfg config.Producer) Producer {
	url := cfg.URL
	id := cfg.ClientId
	acks := cfg.Acks

	if acks == "" {
		acks = _allAcks
	}

	conf := &kafka.ConfigMap{
		"bootstrap.servers": url,
		"client.id":         id,
		"acks":              acks,
	}

	if producer, err := kafka.NewProducer(conf); err != nil {
		panic(err.Error())
	} else {
		return Producer{
			cfg:      cfg,
			producer: producer,
		}
	}
}

func (p Producer) SendEvent(v []byte) {
	topic := p.cfg.Topic

	err := p.producer.Produce(
		&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: kafka.PartitionAny,
			},
			Value: v,
		}, nil)

	if err != nil {
		log.Println("Failed to send topic", string(v))
	} else {
		log.Println("Success to send message", string(v))
	}
}
