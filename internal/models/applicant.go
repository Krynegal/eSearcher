package models

type Applicant struct {
	ID       string
	Name     string
	Lastname string
	Busyness []int
	Schedule []int
}

type SearchApplicantParams struct {
	Limit          int64
	Offset         int64
	Name           string
	Lastname       string
	Specialization []int
	Busyness       []int
	Schedule       []int
}
