package restserver

const (
	IP_PORT         = "IPPORT"
	DEFAULT_IP_PORT = ":8000"
	PARSE_URL       = "/companies/"
	LOGIN_URL       = "/login"
	COMPANY_URL     = "/companies/{id}"
	COMPANIES_URL   = "/companies"
	COMPANY         = "company"
	TOKEN           = "token"
	STATUS          = "status"
	SUCCESS         = "success"
	ERROR           = "error"
	FAILED_GET      = "Resource not found"
	FAILED_CREATE   = "Failed to create resource"
	FAILED_UPDATE   = "Failed to update resource"
	FAILED_DELETE   = "Failed to delete resource"
	FAILED_LOGIN    = "Login failed"
)

type companyType int

const (
	Corporations companyType = iota
	NonProfit
	Cooperative
	SoleProprietorshipt
)
