package models

type Vacancy struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}

type SearchVacancyParams struct {
	Limit  int64
	Offset int64
	Name   string
	Tags   []string
}
