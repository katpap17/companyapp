package main

import (
	server "github.com/katpap17/companyapp/restserver"
	"github.com/katpap17/companyapp/utils"
)

func main() {
	utils.SetUpLogging()
	utils.Logger.Trace("Starting server")
	server.StartServer()

}
