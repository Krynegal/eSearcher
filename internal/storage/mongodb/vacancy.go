package mongodb

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type VacancyCollection struct {
	collection *mongo.Collection
}

func NewVacancyCollection(database *mongo.Database, collection string) *VacancyCollection {
	return &VacancyCollection{
		collection: database.Collection(collection),
	}
}

func (vc *VacancyCollection) Create(name, desc string) (string, error) {
	res, err := vc.collection.InsertOne(context.TODO(), bson.D{
		{
			"name", name,
		},
		{
			"description", desc,
		},
	})
	if err != nil {
		return "-1", err
	}
	if _, ok := res.InsertedID.(primitive.ObjectID); !ok {
		return "-1", errors.New("vacancy id cast error")
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}
