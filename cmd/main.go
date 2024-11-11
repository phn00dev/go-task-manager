package main

import (
	"fmt"
	"log"

	"github.com/phn00dev/go-task-manager/internal/app"
	"github.com/phn00dev/go-task-manager/internal/setup/constructor"
)

func main() {
	fmt.Println("Golang task manager application")

	getDependencies, err := app.GetDependencies()
	if err != nil {
		log.Println("get dependencies error", err)
		return
	}
	constructor.Build(getDependencies)
	runServer := fmt.Sprintf("%s:%s", getDependencies.Config.HttpServer.HttpHost, getDependencies.Config.HttpServer.HttpPort)
	log.Printf("http://%s \n", runServer)
	newApp := app.NewApp(getDependencies.Config)
	if errRunServer := newApp.Run(runServer); errRunServer != nil {
		log.Println("run server error", errRunServer)
		return
	}
}
