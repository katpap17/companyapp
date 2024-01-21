package restserver

import "github.com/katpap17/companyapp/repository"

func validateCompany(company repository.Company) bool {
	return validateCompanyName(company.Name) &&
		validateCompanyDescription(company.Description) &&
		validateCompanyEmployees(company.Employees) &&
		validateCompanyRegistered(company.Registered) &&
		validateCompanyType(company.CompanyType)
}

func validateCompanyName(name string) bool {
	return name != "" && len(name) <= 15
}

func validateCompanyDescription(description string) bool {
	return len(description) <= 3000
}

func validateCompanyEmployees(employees int) bool {
	return employees > 0
}

func validateCompanyRegistered(registered *bool) bool {
	return registered != nil
}

func validateCompanyType(companyType *repository.CompanyType) bool {
	return companyType != nil
}

func validateUser(user repository.User) bool {
	return user.Username != "" && user.Password != ""
}
