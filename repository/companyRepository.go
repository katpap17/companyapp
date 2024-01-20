package repository

import (
	"github.com/gofrs/uuid"
	"github.com/katpap17/companyapp/utils"
)

type companyType int

const (
	Corporations companyType = iota
	NonProfit
	Cooperative
	SoleProprietorshipt
)

type Company struct {
	ID          uuid.UUID   `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Employees   int         `json:"employees"`
	Registered  bool        `json:"registered"`
	CompanyType companyType `json:"companyType"`
}

func (Company) TableName() string {
	return "company"
}

type CompanyRepository struct {
	Repository Repository
}

var companyRepository CompanyRepository

func SetCompanyRepository(db DBHandler) {
	companyRepository = CompanyRepository{Repository: Repository{db: db}}
}

func GetCompany(id uuid.UUID) (*Company, error) {
	return companyRepository.get(id)
}

func CreateCompany(company *Company) error {
	return companyRepository.create(company)
}

func DeleteCompany(id uuid.UUID) error {
	return companyRepository.delete(id)
}

func UpdateCompany(company *Company) error {
	return companyRepository.update(company)
}

func (r *CompanyRepository) get(id uuid.UUID) (*Company, error) {
	var company Company
	if err := r.Repository.db.First(&company, id).Error; err != nil {
		utils.Logger.Error(err.Error())
		return nil, err
	}
	return &company, nil

}

func (r *CompanyRepository) create(company *Company) error {
	id, err := uuid.NewV7()
	if err != nil {
		utils.Logger.Error(err.Error())
		return err
	}
	company.ID = id
	result := r.Repository.db.Create(company)
	if result.Error != nil {
		utils.Logger.Error(result.Error.Error())
		return result.Error
	}
	return nil

}

func (r *CompanyRepository) update(company *Company) error {
	result := r.Repository.db.Save(company)
	if result.Error != nil {
		utils.Logger.Error(result.Error.Error())
		return result.Error
	}
	return nil

}

func (r *CompanyRepository) delete(id uuid.UUID) error {
	result := r.Repository.db.Delete(&Company{}, id)
	if result.Error != nil {
		utils.Logger.Error(result.Error.Error())
		return result.Error
	}
	return nil

}
