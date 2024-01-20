package restserver

import "github.com/katpap17/companyapp/repository"

func validateCompany(company repository.Company) bool {
	if company.Name == "" || len(company.Name) > 15 {
		return false
	} else if len(company.Description) > 3000 {
		return false
	} else if company.Employees == 0 {
		return false
	} else if company.Registered == nil {
		return false
	} else if company.CompanyType == nil {
		return false
	}
	return true
}

func validateUser(user repository.User) bool {
	if user.Username == "" {
		return false
	} else if user.Password == "" {
		return false
	}
	return true
}
