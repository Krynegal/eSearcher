package service

type VacancyService interface {
	CreateVacancy(name, desc string) error
}

type EmployeeService interface {
}

type EmployerService interface {
}

type Services struct {
	VacancyService
	EmployeeService
	EmployerService
}
