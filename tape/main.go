package main

import (
	"log"

	"tape/pkg/actions"
	"tape/pkg/config"
	"tape/pkg/events"
	"tape/pkg/handlers"
)

func main() {

	configFile := "config.yml"
	config := config.IngestConfig("config.yml")

	if config.LogEnabled {
		logFile := events.InitializeLogging(config)
		defer logFile.Close()
	}
	log.Printf("Processed configuration file: %v", configFile)

	log.Printf("Ingesting actions:")
	actList := actions.Initialize(config.ActionsPath, config)

	handlers.Initialize(config, actList)

}
