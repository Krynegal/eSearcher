package mongodb

import (
	"eSearcher/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func applyAllFilters(params *models.SearchVacancyParams) bson.D {
	return bson.D{
		containInName(params.Name),
		containInTags(params.Tags),
	}
}

func containInName(word string) bson.E {
	if word != "" {
		return bson.E{Key: "name", Value: primitive.Regex{Pattern: word}}
	}
	return bson.E{}
}

func containInTags(tags []string) bson.E {
	if len(tags) != 0 {
		return bson.E{Key: "tags", Value: bson.M{"$in": tags}}
	}
	return bson.E{}
}
