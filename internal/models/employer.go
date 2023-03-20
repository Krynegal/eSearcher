package models

type Employer struct {
	ID     int            `json:"id"`
	Info   EmployerInfo   `json:"info"`
	Sphere EmployerSphere `json:"sphere"`
}

type EmployerInfo struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type EmployerSphere struct {
	Sphere  []int `json:"sphere"`
	Added   []int `json:"added,omitempty"`
	Deleted []int `json:"deleted,omitempty"`
}
