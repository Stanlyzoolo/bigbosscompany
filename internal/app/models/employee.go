package models

// Employee ...
type Employee struct {
	ID         int64    `json:"id,omitempty"`
	Name       string   `json:"name,omitempty"`
	SecondName string   `json:"secondname,omitempty"`
	Surname    string   `json:"surname,omitempty"`
	HireDate   string   `json:"hiredate,omitempty"`
	Position   Position `json:"position,omitempty"`
	CompanyID  int64    `json:"companyid,omitempty"`
}

// Position...
type Position struct {
	developer, manager,  
	qualityAssurance, businessAnalyst  string 
}
