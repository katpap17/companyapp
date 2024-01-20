package main

import (
	"github.com/katpap17/companyapp/database"
	"github.com/katpap17/companyapp/repository"
	server "github.com/katpap17/companyapp/restserver"
	"github.com/katpap17/companyapp/utils"
)

func main() {
	utils.SetUpLogging()

	utils.Logger.Trace("Setting up db")
	db, err := database.Init()
	if err != nil {
		utils.Logger.Error("Error in connecting to database")
	}
	repository.SetCompanyRepository(db)

	defer utils.Logger.Trace("Starting server")
	server.StartServer()

}
