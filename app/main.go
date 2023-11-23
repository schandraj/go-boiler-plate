package app

import (
	"base-plate/infrastructure/http"
	"base-plate/infrastructure/kafka"
	"base-plate/infrastructure/mongo"
	"base-plate/infrastructure/redis"
	"base-plate/infrastructure/sql"
	"context"
)

type Container struct {
	KafkaPublisher  *kafka.Publisher
	KafkaSubscriber *kafka.Subscriber
	Cache           *redis.Client
	MongoClient     *mongo.Client
	HttpClient      *http.Client
	SqlClient       *sql.Client
	PSqlClient      *sql.Client
}

func NewContainer(ctx context.Context) *Container {
	kp := kafka.NewPublisher()
	ks := kafka.NewSubscriber()
	c := redis.NewClient()
	mc := mongo.NewClient(ctx)
	hc := http.NewClient()
	sc := sql.NewClient()
	psc := sql.NewPostgreSQLClient()

	return &Container{
		KafkaPublisher:  kp,
		KafkaSubscriber: ks,
		Cache:           c,
		MongoClient:     mc,
		HttpClient:      hc,
		SqlClient:       sc,
		PSqlClient:      psc,
	}
}
