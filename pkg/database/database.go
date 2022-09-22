package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"pr0.bot/internal/queries"
	"pr0.bot/pkg/configs"
)

var MongoDBInstance *queries.MongoDB_Client

func newInstance(url string) *queries.MongoDB_Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		log.Fatalln(err)
	}

	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatalln(err)
	}

	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		log.Fatalln(err)
	}

	return &queries.MongoDB_Client{
		Client: client,
	}
}

func init() {
	MongoDBInstance = newInstance(configs.Config.Items.MongoDBURL)
}
