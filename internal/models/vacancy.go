package models

type Vacancy struct {
	Name           string `json:"name"`
	Description    string `json:"description"`
	Specialization int
	Tags           []string `json:"tags"`
}

type SearchVacancyParams struct {
	Limit          int64
	Offset         int64
	Name           string
	Experience     int
	Schedule       []int
	Busyness       []int
	Specialization []int
	Tags           []string
}
