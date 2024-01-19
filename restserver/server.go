package restserver

import (
	"os"

	"github.com/gorilla/mux"
)

func StartServer() {
	ipPort := os.Getenv(IP_PORT)
	if ipPort == "" {
		ipPort = ":8000"
	}
	router := mux.NewRouter()
	srv := CompanyService{}
	srv.handleRequests(router)
	srv.start(router, ipPort)
}
