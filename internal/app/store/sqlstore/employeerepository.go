package sqlstore

import (
	
	"github.com/Stanlyzoolo/bigbosscompany/internal/app/models"
)

type EmployeeRepository struct {
	store *Store
}

// Add...
func (r *EmployeeRepository) Add(u *models.Employee) error {
	return r.store.db.QueryRow(
		"INSERT INTO employee (name, secondName, surname, hiredate, position, companyid) 
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		 u.Name, u.SecondName, u.Surname, u.HireDate, u.Position, u.CompanyID,)
	 Scan(&u.ID)
}

// Find...
func (r *EmployeeRepository) Find(id int64) (*models.Employee, error) {
	u := &models.Employee{}
	if err:= r.store.db.QueryRow(
		"SELECT id, name, secondName, surname, hiredate, position, companyid FROM employee 
		WHERE id=$1",
		id,
		).Scan(
			&u.ID,
			&u.Name,
			&u.SecondName,
			&u.Surname, 
			&u.HireDate, 
			&u.Position, 
			&u.CompanyID,
		); err != nil {
			return nil, err
		}
		return u, nil
}

// Update...
func (r *EmployeeRepository) Update(u *models.Employee) (*models.Employee, error) {
	u := &models.Employee{}
	if err := r.store.db.QueryRow(
		"UPDATE employee SET (name, secondName, surname, hiredate, position, companyid) 
		WHERE (name=$1, secondName=$2, surname=$3, hiredate=$4, postion=$5, companyid=$6) 
		RETURNING id",
		 u.Name, u.SecondName, u.Surname, u.HireDate, u.Position, u.CompanyID,)
	 Scan(&u.ID); err != nil {
		return u, err
}

// UpdateWithForm...
func (r *EmployeeRepository) UpdateWithForm(u *models.Employee) (*models.Employee, error) {
	return nil, nil
}


// Delete...
func (r *EmployeeRepository) Delete(u *models.Employee) (error) {
	if err:= r.store.db.QueryRow(
		"DELETE FROM employee WHERE id=$1", u.ID
	); err != nil {
		return err
	}
	return err
}
