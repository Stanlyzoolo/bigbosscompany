package store

import "github.com/Stanlyzoolo/bigbosscompany/internal/app/models"

type EmployeeRepository interface {
	Add(*models.Employee) error
	Find(int64) (*models.Employee, error)
	Update(*models.Employee, error)
	Delete(*models.Employee, error)
}