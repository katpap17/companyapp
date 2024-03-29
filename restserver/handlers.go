package restserver

import (
	"encoding/json"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/katpap17/companyapp/actions"
	"github.com/katpap17/companyapp/repository"
	"github.com/katpap17/companyapp/utils"
)

func GetCompany(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len(PARSE_URL):]
	parsedID, err := uuid.FromString(id)
	if err != nil {
		utils.Logger.Error(err)
		handleErrorResponse(w, FAILED_GET, http.StatusBadRequest)
		return
	}
	utils.Logger.Debug("fetching company with id: ", parsedID)
	company, err := actions.GetCompany(parsedID)
	if err != nil {
		utils.Logger.Error("Failed to fetch company with error: ", err)
		handleErrorResponse(w, FAILED_GET, http.StatusInternalServerError)
		return
	}
	if company == nil {
		utils.Logger.Error("Company not found ", parsedID)
		handleErrorResponse(w, FAILED_GET, http.StatusNotFound)
		return
	}
	res := map[string]interface{}{COMPANY: company}
	handleSuccessFullResponse(w, res)
}

func CreateCompany(w http.ResponseWriter, r *http.Request) {
	var company repository.Company
	err := json.NewDecoder(r.Body).Decode(&company)
	if err != nil {
		utils.Logger.Error("Company not decoded with err: ", err)
		handleErrorResponse(w, FAILED_CREATE, http.StatusBadRequest)
		return
	}
	utils.Logger.Debug("Validating company: ", company)
	validated := validateCompany(company)
	if !validated {
		utils.Logger.Info("Company not validated successfully")
		handleErrorResponse(w, FAILED_CREATE, http.StatusBadRequest)
		return
	}
	utils.Logger.Debug("Creating company: ", company)
	err, id := actions.CreateCompany(&company)
	if err != nil {
		utils.Logger.Error("Company not created with err: ", err)
		handleErrorResponse(w, FAILED_CREATE, http.StatusInternalServerError)
		return
	}
	res := map[string]interface{}{"id": id.String()}
	handleSuccessFullResponse(w, res)
}

func UpdateCompany(w http.ResponseWriter, r *http.Request) {
	var company repository.Company
	id := r.URL.Path[len(PARSE_URL):]
	parsedID, err := uuid.FromString(id)
	if err != nil {
		utils.Logger.Error(err)
		handleErrorResponse(w, FAILED_GET, http.StatusBadRequest)
		return
	}
	if r.Body == nil {
		utils.Logger.Error("Body request is empty")
		handleErrorResponse(w, FAILED_UPDATE, http.StatusBadRequest)
	}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&company)
	if err != nil {
		utils.Logger.Error("Company not decoded with err: ", err)
		handleErrorResponse(w, FAILED_UPDATE, http.StatusBadRequest)
		return
	}
	company.ID = parsedID
	utils.Logger.Debug("Validating company: ", company)
	validated := validateCompany(company)
	if !validated {
		utils.Logger.Info("Company not validated successfully")
		handleErrorResponse(w, FAILED_CREATE, http.StatusBadRequest)
		return
	}
	utils.Logger.Debug("Updating company: ", company)
	err = actions.UpdateCompany(&company)
	if err != nil {
		utils.Logger.Error("Company not updated with err: ", err)
		handleErrorResponse(w, FAILED_UPDATE, http.StatusInternalServerError)
		return
	}
	handleSuccessFullResponse(w, nil)
}

func DeleteCompany(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len(PARSE_URL):]
	parsedID, err := uuid.FromString(id)
	if err != nil {
		utils.Logger.Error(parsedID)
		handleErrorResponse(w, FAILED_GET, http.StatusBadRequest)
		return
	}
	utils.Logger.Debug("Deleting company with id: ", parsedID)
	err = repository.DeleteCompany(parsedID)
	if err != nil {
		handleErrorResponse(w, FAILED_DELETE, http.StatusInternalServerError)
		return
	}
	handleSuccessFullResponse(w, nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user repository.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		utils.Logger.Error("User not decoded with err: ", err)
		handleErrorResponse(w, FAILED_UPDATE, http.StatusBadRequest)
		return
	}
	utils.Logger.Debug("Validating user: ", user.Username)
	validated := validateUser(user)
	if !validated {
		utils.Logger.Info("User not validated successfully")
		handleErrorResponse(w, FAILED_CREATE, http.StatusBadRequest)
		return
	}
	token, err := actions.Login(&user)
	if err != nil {
		utils.Logger.Error("Failed to generate token with error: ", err)
		handleErrorResponse(w, FAILED_LOGIN, http.StatusInternalServerError)
		return
	} else if token == "" {
		handleErrorResponse(w, FAILED_LOGIN, http.StatusUnauthorized)
		return
	}
	res := map[string]interface{}{TOKEN: token}
	handleSuccessFullResponse(w, res)
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
