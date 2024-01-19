package restserver

const (
	IP_PORT       = "IPPORT"
	PARSE_URL     = "/companies/"
	COMPANY_URL   = "/companies/{id}"
	COMPANIES_URL = "/companies"
	COMPANY       = "company"
	STATUS        = "status"
	SUCCESS       = "success"
	ERROR         = "error"
	FAILED_GET    = "Resource not found"
	FAILED_CREATE = "Failed to create resource"
	FAILED_UPDATE = "Failed to update resource"
	FAILED_DELETE = "Failed to delete resource"
)

type companyType int

const (
	Corporations companyType = iota
	NonProfit
	Cooperative
	SoleProprietorshipt
)
