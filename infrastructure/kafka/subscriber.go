package kafka

import (
	"base-plate/config"
	"github.com/Shopify/sarama"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"log"
	"strings"
)

type Subscriber struct {
	*kafka.Subscriber
}

func NewSubscriber() *Subscriber {
	saramaSubscriberConfig := kafka.DefaultSaramaSubscriberConfig()
	// equivalent of auto.offset.reset: earliest
	saramaSubscriberConfig.Consumer.Offsets.Initial = sarama.OffsetOldest

	subscriber, err := kafka.NewSubscriber(
		kafka.SubscriberConfig{
			Brokers:               strings.Split(config.Cfg.GetString("KAFKA_BROKER"), ","),
			Unmarshaler:           kafka.DefaultMarshaler{},
			OverwriteSaramaConfig: saramaSubscriberConfig,
			ConsumerGroup:         "interaction",
		},
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		log.Println(err)
	}

	return &Subscriber{subscriber}
}
