package models

type Applicant struct {
	ID              int                       `json:"id"`
	Info            ApplicantInfo             `json:"info"`
	Experiences     []ApplicantExperience     `json:"experiences"`
	Educations      []ApplicantEducation      `json:"educations"`
	Languages       []ApplicantLanguage       `json:"languages"`
	Specializations []ApplicantSpecialization `json:"specializations"`
	Busyness        ApplicantBusyness         `json:"busyness"`
	Schedule        ApplicantSchedule         `json:"schedule"`
}

type ApplicantInfo struct {
	Name        string `json:"name"`
	Lastname    string `json:"lastname"`
	Status      int    `json:"status"`
	Phone       string `json:"phone"`
	Birthday    string `json:"birthday"`
	Description string `json:"description"`
	Male        bool   `json:"male"`
}

type ApplicantExperience struct {
	ID           int    `json:"id"`
	Start        string `json:"start"`
	Finish       string `json:"finish"`
	Organization string `json:"organization"`
	Position     string `json:"position"`
	Duties       string `json:"duties"`
	Skills       string `json:"skills"`
}

type ApplicantEducation struct {
	ID             int    `json:"id"`
	Institution    int    `json:"institution_id"`
	Grade          int    `json:"grade"`
	Faculty        string `json:"faculty"`
	Specialization int    `json:"specialization"`
	Finish         string `json:"finish"`
}

type ApplicantSpecialization struct {
	ID             int `json:"id"`
	Specialization int `json:"specialization"`
	Salary         int `json:"salary"`
}

type ApplicantLanguage struct {
	ID       int `json:"id"`
	Language int `json:"language"`
	Level    int `json:"level"`
}

type ApplicantBusyness struct {
	Busyness []int `json:"busyness"`
	Deleted  []int `json:"deleted,omitempty"`
	Added    []int `json:"added,omitempty"`
}

type ApplicantSchedule struct {
	Schedule []int `json:"schedule"`
	Deleted  []int `json:"deleted,omitempty"`
	Added    []int `json:"added,omitempty"`
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
