package main

import (
	"github.com/katpap17/companyapp/actions"
	"github.com/katpap17/companyapp/database"
	"github.com/katpap17/companyapp/repository"
	server "github.com/katpap17/companyapp/restserver"
	"github.com/katpap17/companyapp/utils"
)

const DEFAULT_USER = "DEFAULT_USER"
const USER_PASS = "USER_PASS"

func main() {
	utils.SetUpLogging()
	var err error
	err = utils.LoadConfig()
	if err != nil {
		utils.Logger.Error("Error loading configuration:", err)
	}

	utils.Logger.Trace("Setting up db")
	db, err := database.Init()
	if err != nil {
		utils.Logger.Error("Error in connecting to database")
	}
	repository.SetDB(db)

	utils.Logger.Trace("Creating default user")
	CreateDefaultUser()

	defer utils.Logger.Trace("Starting server")
	server.StartServer()

}

func CreateDefaultUser() {
	userName := utils.GetEnv(DEFAULT_USER, "")
	userPass := utils.GetEnv(USER_PASS, "")
	if userName == "" || userPass == "" {
		utils.Logger.Warn("Default user vars not provided")
		return
	}
	err := actions.CreateUser(userName, userPass)
	if err != nil {
		utils.Logger.Error("Default user not created with error: ", err)
		return
	}
	utils.Logger.Trace("Default user created successfully: ", userName)
}
