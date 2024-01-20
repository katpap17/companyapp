package actions

import (
	"github.com/gofrs/uuid"
	"github.com/katpap17/companyapp/repository"
	"github.com/katpap17/companyapp/utils"
)

func GetCompany(id uuid.UUID) (*repository.Company, error) {
	return repository.GetCompany(id)
}

func CreateCompany(company *repository.Company) (error, uuid.UUID) {
	id, err := uuid.NewV7()
	if err != nil {
		utils.Logger.Error(err.Error())
		return err, uuid.Nil
	}
	company.ID = id
	return repository.CreateCompany(company), id
}

func DeleteCompany(id uuid.UUID) error {
	return repository.DeleteCompany(id)
}

func UpdateCompany(company *repository.Company) error {
	return repository.UpdateCompany(company)
}
