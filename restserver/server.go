package restserver

import (
	"github.com/gorilla/mux"
	"github.com/katpap17/companyapp/utils"
)

func StartServer() {
	ipPort := utils.GetEnv(IP_PORT, DEFAULT_IP_PORT)
	router := mux.NewRouter()
	srv := CompanyService{}
	srv.handleRequests(router)
	srv.start(router, ipPort)
}
