package store

import "github.com/Stanlyzoolo/bigbosscompany/internal/app/models"

type CompanyRepository interface {
	Place(*models.Company) error
	Find(int32) (*models.Company, error)
	Update(*models.Company, error)
	GetList(*models.Company, error)
	Delete(*models.Company, error)
}