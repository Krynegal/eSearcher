package models

type Response struct {
	ApplicantID int    `json:"applicant_id"`
	VacancyID   string `json:"vacancy_id"`
	StatusID    int    `json:"status_id"`
}
