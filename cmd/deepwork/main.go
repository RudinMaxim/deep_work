package main

import (
	"log"
	"os"

	"deep-work-app/internal/app"
	"deep-work-app/internal/config"
	"deep-work-app/internal/interfaces/cli"
)

func main() {
	_config, err := config.LoadConfig("../../config/config.json")
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
		os.Exit(1)
	}

	application := app.NewApplication(_config)
	cliHandler := cli.NewCLIHandler(application)

	cliHandler.Execute()
}
