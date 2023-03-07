package mongodb

import (
	"eSearcher/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func applyAllFilters(params *models.SearchVacancyParams) bson.D {
	return bson.D{
		containsName(params.Name),
		containsTags(params.Tags),
		containsExperience(params.Experience),
		containSchedule(params.Schedule),
		containBusyness(params.Busyness),
		containSpecialization(params.Specialization),
	}
}

func containsName(word string) bson.E {
	if word != "" {
		return bson.E{Key: "name", Value: primitive.Regex{Pattern: word}}
	}
	return bson.E{}
}

func containsExperience(experience int) bson.E {
	if experience != 0 {
		return bson.E{Key: "experience", Value: experience}
	}
	return bson.E{}
}

func containSchedule(schedule []int) bson.E {
	if len(schedule) != 0 {
		return bson.E{Key: "tags", Value: bson.M{"$in": schedule}}
	}
	return bson.E{}
}

func containBusyness(busyness []int) bson.E {
	if len(busyness) != 0 {
		return bson.E{Key: "tags", Value: bson.M{"$in": busyness}}
	}
	return bson.E{}
}

func containSpecialization(specialization []int) bson.E {
	if len(specialization) != 0 {
		return bson.E{Key: "tags", Value: bson.M{"$in": specialization}}
	}
	return bson.E{}
}

func containsTags(tags []string) bson.E {
	if len(tags) != 0 {
		return bson.E{Key: "tags", Value: bson.M{"$in": tags}}
	}
	return bson.E{}
}
