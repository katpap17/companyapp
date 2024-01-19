package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type companyType int

const (
	Corporations companyType = iota
	NonProfit
	Cooperative
	SoleProprietorshipt
)

type Company struct {
	ID          uuid.UUID
	Name        string
	Description string
	Employees   int
	Registered  bool
	Type        companyType
}

func GetCompany(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/companies/"):]
	company := Company{
		ID:          uuid.New(),
		Name:        "XM",
		Description: id,
	}
	json.NewEncoder(w).Encode(company)
}
