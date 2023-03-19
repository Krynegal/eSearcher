package models

type Applicant struct {
	ID              string                    `json:"id"`
	Info            ApplicantInfo             `json:"info"`
	Experiences     []ApplicantExperience     `json:"experiences"`
	Educations      []ApplicantEducation      `json:"educations"`
	Languages       []ApplicantLanguage       `json:"languages"`
	Specializations []ApplicantSpecialization `json:"specializations"`
	Busyness        []int                     `json:"busyness"`
	Schedule        []int                     `json:"schedule"`
}

type ApplicantInfo struct {
	Name        string  `json:"name"`
	Lastname    string  `json:"lastname"`
	Status      int     `json:"status"`
	Phone       *string `json:"phone"`
	Birthday    *string `json:"birthday"`
	Description *string `json:"description"`
	Male        bool    `json:"male"`
}

type ApplicantExperience struct {
	Start        *string `json:"start,omitempty"`
	Finish       *string `json:"finish,omitempty"`
	Organization *string `json:"organization,omitempty"`
	Position     *string `json:"position,omitempty"`
	Duties       *string `json:"duties,omitempty"`
	Skills       *string `json:"skills,omitempty"`
}

type ApplicantEducation struct {
	Institution    *int    `json:"institution_id"`
	Grade          *int    `json:"grade"`
	Faculty        *string `json:"faculty"`
	Specialization *int    `json:"specialization"`
	Finish         *string `json:"finish"`
}

type ApplicantSpecialization struct {
	Specialization *int `json:"specialization"`
	Salary         *int `json:"salary"`
}

type ApplicantLanguage struct {
	Language *int `json:"language"`
	Level    *int `json:"level"`
}

type SearchApplicantParams struct {
	Limit          int64  `json:"limit"`
	Offset         int64  `json:"offset"`
	Name           string `json:"name"`
	Lastname       string `json:"lastname"`
	Specialization []int  `json:"specialization"`
	Busyness       []int  `json:"busyness"`
	Schedule       []int  `json:"schedule"`
}
