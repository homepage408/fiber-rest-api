package configuration

import (
	"os"
	"strconv"
)

type Configuration struct {
	Service  Service
	Hosts    Hosts
	Database DatabaseConf
}

type Service struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Address string `json:"address"`
	Port    int    `json:"port"`
}

type Hosts struct {
	Database DatabaseConf
}

type DatabaseConf struct {
	Name     string
	Username string
	Password string
	Host     string
	Port     int
}

func (conf *Configuration) Init() {
	databasePort, _ := strconv.Atoi(os.Getenv("DATABASE_PORT"))
	servicePort, _ := strconv.Atoi(os.Getenv("SERVICE_PORT"))

	conf.Database = DatabaseConf{
		Name:     os.Getenv("DATABASE_NAME"),
		Username: os.Getenv("DATABASE_USERNAME"),
		Password: os.Getenv("DATABASE_PASSWORD"),
		Host:     os.Getenv("DATABASE_HOST"),
		Port:     databasePort,
	}

	conf.Service = Service{
		Name:    os.Getenv("SERVICE_NAME"),
		Version: os.Getenv("SERVICE_VERSION"),
		Address: os.Getenv("SERVICE_ADDRESS"),
		Port:    servicePort,
	}
}
