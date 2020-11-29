package models

// Company...
type Company struct {
	ID        int64     `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	LegalForm LegalForm `json:"legalform,omitempty"`
}

// LegalForm...
type LegalForm struct {
	OOO, ZAO, IP string
}
