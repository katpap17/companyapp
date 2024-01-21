package utils

import (
	"encoding/json"
	"os"
)

var Cfg Configuration

type Server struct {
	WriteTimeout int `json:"write_timeout"`
	ReadTimeout  int `json:"read_timeout"`
	IdleTimeout  int `json:"idle_timeout"`
}

type Configuration struct {
	Server Server `json:"server"`
}

func LoadConfig() error {
	confFile := GetEnv("CONF_FILE_PATH", "config.json")

	file, err := os.Open(confFile)
	if err != nil {
		SetDefaultConfiguration()
		return err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Cfg)
	if err != nil {
		SetDefaultConfiguration()
		return err
	}
	return nil
}

func SetDefaultConfiguration() {
	Cfg = Configuration{
		Server: Server{
			WriteTimeout: 300,
			ReadTimeout:  300,
			IdleTimeout:  300,
		},
	}
}
