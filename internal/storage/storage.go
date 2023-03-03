package storage

type VacancyStorage interface {
	Create(name, desc string) (string, error)
}

type EmployeeStorage interface {
}

type EmployerStorage interface {
}

type Storage struct {
	VacancyStorage
	EmployeeStorage
	EmployerStorage
}
