package store


type Store interface {
	Employee() EmployeeRepository
	Company() CompanyRepository
}
