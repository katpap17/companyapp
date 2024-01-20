package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Logger = logrus.New()

const logFileName string = "LOGFILE"
const defaultLogFile string = "companyapp.log"

func SetUpLogging() {
	filepath := GetEnv(logFileName, defaultLogFile)
	logFile, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		Logger.Fatal("Error opening log file:", err)
	}
	Logger.SetOutput(logFile)
	Logger.SetFormatter(&logrus.JSONFormatter{})
	Logger.SetLevel(logrus.DebugLevel)
}
