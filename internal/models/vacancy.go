package models

type Vacancy struct {
	ID             string   `json:"id"`
	Name           string   `json:"name"`
	Description    string   `json:"description"`
	Salary         int      `json:"salary"`
	Experience     int      `json:"experience"`
	Busyness       []int    `json:"busyness"`
	Schedule       []int    `json:"schedule"`
	Specialization []int    `json:"specialization"`
	Tags           []string `json:"tags"`
	Status         bool     `json:"status"`
}

type SearchVacancyParams struct {
	Limit          int64    `json:"limit"`
	Offset         int64    `json:"offset"`
	Name           string   `json:"name"`
	Description    string   `json:"description"`
	Salary         int      `json:"salary"`
	Experience     int      `json:"experience"`
	Busyness       []int    `json:"busyness"`
	Schedule       []int    `json:"schedule"`
	Specialization []int    `json:"specialization"`
	Tags           []string `json:"tags"`
	Status         bool     `json:"status"`
}
