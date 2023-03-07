package postgres

import (
	"eSearcher/internal/models"
)

func FillEmpty(params *models.SearchApplicantParams) {
	if len(params.Specialization) == 0 {
		params.Specialization = []int{1, 2, 3, 4, 5}
	}
	if len(params.Schedule) == 0 {
		params.Schedule = []int{1, 2, 3, 4, 5}
	}
	if len(params.Busyness) == 0 {
		params.Busyness = []int{1, 2, 3, 4, 5}
	}
}
