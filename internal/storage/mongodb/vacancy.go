package mongodb

import (
	"context"
	"eSearcher/internal/models"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type VacancyCollection struct {
	collection *mongo.Collection
}

func NewVacancyCollection(database *mongo.Database, collection string) *VacancyCollection {
	return &VacancyCollection{
		collection: database.Collection(collection),
	}
}

func (vc *VacancyCollection) Create(vacancy *models.Vacancy) (string, error) {
	res, err := vc.collection.InsertOne(context.TODO(), bson.D{
		{
			"name", vacancy.Name,
		},
		{
			"description", vacancy.Description,
		},
		{
			"tags", vacancy.Tags,
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

func (vc *VacancyCollection) Search(params *models.SearchVacancyParams) ([]*models.Vacancy, error) {
	filter := applyAllFilters(params)
	var opts options.FindOptions
	opts.SetLimit(params.Limit)
	opts.SetSkip(params.Offset)
	cursor, err := vc.collection.Find(context.TODO(), filter, &opts)
	if err != nil {
		return nil, err
	}
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}
	var vacancies []*models.Vacancy
	for _, result := range results {
		var vacancy *models.Vacancy
		bsonBytes, _ := bson.Marshal(result)
		bson.Unmarshal(bsonBytes, &vacancy)
		vacancies = append(vacancies, vacancy)
		fmt.Printf("%+v", vacancy)
	}
	return vacancies, nil
}
