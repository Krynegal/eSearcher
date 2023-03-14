package mongodb

import (
	"context"
	"eSearcher/configs"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func New(cfg *configs.Config) (*mongo.Database, error) {
	mongoDBURL := fmt.Sprintf("mongodb://%s:%s", cfg.Mongo.MongoHost, cfg.Mongo.MongoPort)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoDBURL))
	if err != nil {
		return nil, err
	}
	if err = client.Ping(context.TODO(), nil); err != nil {
		return nil, fmt.Errorf("mongo %v", err)
	}

	return client.Database(cfg.Mongo.MongoName), nil
}
