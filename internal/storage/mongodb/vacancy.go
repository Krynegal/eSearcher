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
func (vc *VacancyCollection) GetEmployerVacancies(uid int) ([]*models.Vacancy, error) {
	cursor, err := vc.collection.Find(context.TODO(), bson.D{bson.E{Key: "employerID", Value: uid}})
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
	}
	return vacancies, err
}

func (vc *VacancyCollection) GetByID(id string) (*models.Vacancy, error) {
	cursor, err := vc.collection.Find(context.TODO(), bson.D{bson.E{Key: "_id", Value: id}})
	var result bson.M
	if err = cursor.All(context.TODO(), &result); err != nil {
		return nil, err
	}
	var vacancy models.Vacancy
	bsonBytes, _ := bson.Marshal(result)
	bson.Unmarshal(bsonBytes, &vacancy)
	return &vacancy, err
}

func (vc *VacancyCollection) Create(vacancy *models.Vacancy) (string, error) {
	res, err := vc.collection.InsertOne(context.TODO(), bson.D{
		{"name", vacancy.Name},
		{"description", vacancy.Description},
		{"salary", vacancy.Salary},
		{"experience", vacancy.Experience},
		{"busyness", vacancy.Busyness},
		{"schedule", vacancy.Schedule},
		{"specialization", vacancy.Specialization},
		{"tags", vacancy.Tags},
		{"banned", vacancy.Banned},
	})
	if err != nil {
		return "-1", err
	}
	if _, ok := res.InsertedID.(primitive.ObjectID); !ok {
		return "-1", errors.New("vacancy id cast error")
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (vc *VacancyCollection) Update(vacancy *models.Vacancy) error {
	id, _ := primitive.ObjectIDFromHex(vacancy.ID)
	filter := bson.D{{"_id", id}}
	update := bson.D{
		{"$set", bson.D{{"name", vacancy.Name}}},
		{"$set", bson.D{{"description", vacancy.Description}}},
		{"$set", bson.D{{"salary", vacancy.Salary}}},
		{"$set", bson.D{{"experience", vacancy.Experience}}},
		{"$set", bson.D{{"busyness", vacancy.Busyness}}},
		{"$set", bson.D{{"schedule", vacancy.Schedule}}},
		{"$set", bson.D{{"specialization", vacancy.Specialization}}},
		{"$set", bson.D{{"tags", vacancy.Tags}}},
		{"$set", bson.D{{"banned", vacancy.Banned}}},
	}
	_, err := vc.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
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
