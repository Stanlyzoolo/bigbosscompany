package sqlstore

import (
	"github.com/Stanlyzoolo/bigbosscompany/internal/app/models"
)

type CompanyRepository struct {
	store *Store
}

// Place...
func (r *CompanyRepository) Place(c *models.Company) error {
	return r.store.db.QueryRow(
		"INSERT INTO company (name, legalform,) 
		VALUES ($1, $2) RETURNING id"
		 c.Name, c.LegalForm)
	 Scan(&c.ID)
}

// FindCompany...
func (r *CompanyRepository) Find(id int32) (*models.Company, error) {
	c := &models.Company{}
	if err:= r.store.db.QueryRow(
		"SELECT id, name, legalform FROM company 
		WHERE id=$1, name=$2, legalform=$3",
		id, 
		).Scan(
			&c.ID,
			&c.Name,
			&c.LegalForm,
		); err != nil {
			return nil, err
		}
		return c, nil
}

// UpdateCompany...
func (r *CompanyRepository) Update(c *models.Company) (*models.Company, error) {
	c := &models.Company{}
	if err := r.store.db.QueryRow(
		"UPDATE company SET (name, legalform) 
		WHERE (name=$1, legalform=$2) 
		RETURNING id",
		 c.Name, c.LegalForm,)
	 Scan(&c.ID); err != nil {
		return c, err
}

// GetList...
func (r *CompanyRepository) GetList(c *models.Company) (*models.Company, error) {
	c := &models.Company{}
	if err:= r.store.db.QueryRow(
		"SELECT  * FROM company"); err != nil {
			return nil, err
		}
		return c, nil
}


// Delete...
func (r *CompanyRepository) Delete(c *models.Company) (error) {
	if err:= r.store.db.QueryRow(
		"DELETE FROM company WHERE id=$1", c.ID
	); err != nil {
		return err
	}
}
