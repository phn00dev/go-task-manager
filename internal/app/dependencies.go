package app

import (
	"net/http"

	"gorm.io/gorm"

	"github.com/phn00dev/go-task-manager/pkg/config"
	dbconnect "github.com/phn00dev/go-task-manager/pkg/database/db_connect"
	httpclient "github.com/phn00dev/go-task-manager/pkg/http_client"
)

type Dependencies struct {
	DB         *gorm.DB
	HttpClient *http.Client
	HttpServer *http.Server
	Config     *config.Config
}

func GetDependencies() (*Dependencies, error) {
	getConfig, err := config.GetConfig()
	if err != nil {
		return nil, err
	}
	// db config
	newDB := dbconnect.NewDBConnnect(getConfig)
	getDB, err := newDB.GetDBConnect()
	if err != nil {
		return nil, err
	}

	// get seeder
	//dbSeeder := seeders.NewDbSeeder(getDB)
	//if err := dbSeeder.GetDBSeeder(); err != nil {
	//	return nil, err
	//}

	// get httpClient
	newHttpClient := httpclient.NewHttpClient()
	newHttpServer := httpclient.NewHttpServer()
	return &Dependencies{
		DB:         getDB,
		Config:     getConfig,
		HttpClient: newHttpClient,
		HttpServer: newHttpServer,
	}, nil
}
