package events

import (
	"log"
	"os"
	"path/filepath"
	"tape/pkg/models"
)

func InitializeLogging(config *models.ConfigObj) *os.File {
	logDir := config.LogPath
	logPath := filepath.Join(logDir, "event.log")

	// Ensure the log directory exists
	if err := os.MkdirAll(logDir, 0755); err != nil {
		CheckError(err)
	}

	_, err := os.Stat(logPath)
	if os.IsNotExist(err) {
		f, createErr := os.Create(logPath)
		if createErr != nil {
			CheckError(createErr)
		} else {
			f.Close()
		}
	}

	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	CheckError(err)

	log.SetOutput(logFile)
	return logFile
}
