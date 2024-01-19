package restserver

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type Company struct {
	ID          uuid.UUID `json:id,omitempty`
	Name        string
	Description string
	Employees   int
	Registered  bool
	Type        companyType
}

func mockCompany() (Company, error) {
	return Company{ID: uuid.Nil, Name: "XM"}, nil
}

func GetCompany(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len(PARSE_URL):]
	company, err := mockCompany()
	if err != nil {
		handleErrorResponse(w, FAILED_GET, http.StatusInternalServerError)
		return
	}
	if company == (Company{}) {
		handleErrorResponse(w, FAILED_GET, http.StatusNotFound)
		return
	}
	company.Description = id
	res := map[string]interface{}{COMPANY: company}
	handleSuccessFullResponse(w, res)
}

func CreateCompany(w http.ResponseWriter, r *http.Request) {
	var company Company
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&company)
	if err != nil {
		handleErrorResponse(w, FAILED_CREATE, http.StatusBadRequest)
		return
	}
	_, err = mockCompany()
	if err != nil {
		handleErrorResponse(w, FAILED_CREATE, http.StatusInternalServerError)
		return
	}
	handleSuccessFullResponse(w, nil)
}

func UpdateCompany(w http.ResponseWriter, r *http.Request) {
	var company Company
	if r.Body == nil {
		handleErrorResponse(w, FAILED_UPDATE, http.StatusBadRequest)
	}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&company)
	if err != nil {
		handleErrorResponse(w, FAILED_UPDATE, http.StatusBadRequest)
		return
	}
	_, err = mockCompany()
	if err != nil {
		handleErrorResponse(w, FAILED_UPDATE, http.StatusInternalServerError)
		return
	}
	handleSuccessFullResponse(w, nil)
}

func DeleteCompany(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len(PARSE_URL):]
	company, err := mockCompany()
	if err != nil {
		handleErrorResponse(w, FAILED_DELETE, http.StatusInternalServerError)
		return
	}
	if company == (Company{}) {
		handleErrorResponse(w, FAILED_DELETE, http.StatusNotFound)
		return
	}
	company.Description = id
	company.ID = uuid.New()
	handleSuccessFullResponse(w, nil)
}

func handleErrorResponse(w http.ResponseWriter, err string, status int) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(err)
}

func handleSuccessFullResponse(w http.ResponseWriter, body map[string]interface{}) {
	w.WriteHeader(http.StatusOK)
	if body != nil {
		body[STATUS] = SUCCESS
	} else {
		body = map[string]interface{}{STATUS: SUCCESS}
	}

	json.NewEncoder(w).Encode(body)
}
