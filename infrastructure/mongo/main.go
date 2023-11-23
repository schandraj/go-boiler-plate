package mongo

import (
	"base-plate/config"
	"context"
	"go.elastic.co/apm/module/apmmongo/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Client struct {
	*mongo.Client
}

func NewClient(ctx context.Context) *Client {
	configURL := config.Cfg.GetString("MONGO_URL")

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(configURL), options.Client().SetMonitor(apmmongo.CommandMonitor()))
	if err != nil {
		log.Println(err)
	}

	return &Client{client}
}
