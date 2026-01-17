package config

import (
	"os"
	"tape/pkg/events"
	"tape/pkg/models"

	"gopkg.in/yaml.v3"
)

func IngestConfig(configFile string) *models.ConfigObj {

	inFile, err := os.ReadFile(configFile)
	events.CheckError(err)

	config := &models.ConfigObj{}
	err = yaml.Unmarshal(inFile, config)
	events.CheckError(err)

	return (config)
}
