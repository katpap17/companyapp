package restserver

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/katpap17/companyapp/utils"
)

type CompanyService struct {
	srv *http.Server
}

func (c *CompanyService) start(router *mux.Router, IPPort string) {
	c.srv = &http.Server{
		Addr:         IPPort,
		WriteTimeout: time.Second * time.Duration(utils.Cfg.Server.WriteTimeout),
		ReadTimeout:  time.Second * time.Duration(utils.Cfg.Server.ReadTimeout),
		IdleTimeout:  time.Second * time.Duration(utils.Cfg.Server.IdleTimeout),
		Handler:      router,
	}
	c.srv.ListenAndServe()
}

func (c *CompanyService) handleRequests(router *mux.Router) {
	router.HandleFunc(COMPANY_URL, GetCompany).Methods(http.MethodGet)
	router.HandleFunc(LOGIN_URL, Login).Methods(http.MethodPost)
	router.HandleFunc(COMPANIES_URL, AuthMiddleware(CreateCompany)).Methods(http.MethodPost)
	router.HandleFunc(COMPANY_URL, AuthMiddleware(UpdateCompany)).Methods(http.MethodPatch)
	router.HandleFunc(COMPANY_URL, AuthMiddleware(DeleteCompany)).Methods(http.MethodDelete)
}
