package kafka

import (
	"base-plate/config"
	"github.com/Shopify/sarama"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"log"
	"strings"
)

type Publisher struct {
	*kafka.Publisher
}

func NewPublisher() *Publisher {
	saramaSubscriberConfig := kafka.DefaultSaramaSubscriberConfig()
	saramaSubscriberConfig.Consumer.Offsets.Initial = sarama.OffsetOldest

	publisher, err := kafka.NewPublisher(
		kafka.PublisherConfig{
			Brokers:   strings.Split(config.Cfg.GetString("KAFKA_BROKERS"), ","),
			Marshaler: kafka.DefaultMarshaler{},
		},
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		log.Println(err)
	}

	return &Publisher{publisher}
}
