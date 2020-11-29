package sqlstore

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Store struct {
	db                 *sql.DB
	employeeRepository *EmployeeRepository
	companyRepository  *CompanyRepository
}

// New...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// Employee...
func (s *Store) Employee() store.EmployeeRepository {
	if s.employeeRepository != nil {
		return s.employeeRepository
	}

	s.employeeRepository = &EmployeeRepository{
		store: s,
	}

	return s.employeeRepository
}

// Company
func (s *Store) Company() store.CompanyRepository {
	if s.companyRepository != nil {
		return s.companyRepository
	}

	s.companyRepository = &CompanyRepository{
		store: s,
	}

	return s.companyRepository
}

// Пример обращения: store.Employee().Add()
// 						store.Company().Add()
