package actions

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"os"
	"path/filepath"
	"tape/pkg/events"
	"tape/pkg/models"

	"gopkg.in/yaml.v3"
)

func Initialize(actionPath string, config *models.ConfigObj) ActionList {
	actionConfigs := enumerate(actionPath)
	actions := parse(actionPath, actionConfigs)

	coreActions := GetCoreActions(structToMap(config))
	for route, core := range coreActions {
		if core.Enabled {
			key := core.Key
			if core.Keygen {
				k, err := genKey()
				if err != nil {
					key = "None"
				} else {
					key = k
				}
			}
			actions[route] = Action{
				Obj: models.ActionObj{
					Name:        core.Name,
					Description: core.Description,
					Route:       route,
					Method:      core.Method,
					Keygen:      core.Keygen,
					Command:     core.Command,
					Input:       core.Input,
					Data:        core.Data,
					OutputWrite: core.OutputWrite,
					OutputFile:  core.OutputFile,
					Key:         key,
				},
				CoreHandler: core.Handler,
			}
			log.Printf("Core Action (%s) Route: (%s) Key: (%s) State: (enabled=%v)", core.Name, route, key, core.Enabled)
		}
	}
	return actions
}

func structToMap(cfg *models.ConfigObj) map[string]interface{} {
	cfgMap := map[string]interface{}{
		"hostname":        cfg.Hostname,
		"listen_port":     cfg.ListenPort,
		"tls_enabled":     cfg.TlsEnabled,
		"tls_certificate": cfg.TlsCertificate,
		"tls_key":         cfg.TlsKey,
		"actions_path":    cfg.ActionsPath,
		"log_enabled":     cfg.LogEnabled,
		"log_path":        cfg.LogPath,
		"file_store":      cfg.FileStore,
		"core_actions":    cfg.CoreActions,
	}
	return cfgMap
}

func enumerate(actionPath string) []os.DirEntry {
	actionConfigs, err := os.ReadDir(actionPath)
	events.CheckError(err)

	return (actionConfigs)
}

func parse(actionPath string, actionConfigs []os.DirEntry) ActionList {

	actionsParsed := ActionList{}

	for _, file := range actionConfigs {
		action, err := ingest(filepath.Join(actionPath, file.Name()))

		if err != nil {
			log.Printf("Failed parsing %s - Error: %v", file.Name(), err)
			continue
		}

		// Prevents duplicate routes
		if _, exists := actionsParsed[action.Obj.Route]; exists {
			log.Printf("Duplicate route detected: %s in file %s. Skipping.", action.Obj.Route, file.Name())
			continue
		}

		actionsParsed[action.Obj.Route] = action
		log.Printf("Action (%s) Route: (%s) Key: (%s)", file.Name(), action.Obj.Route, action.Obj.Key)
	}

	// If no actions parsed, end processing
	return actionsParsed

}

func ingest(filePath string) (Action, error) {

	buffer, err := os.ReadFile(filePath)
	events.CheckError(err)

	actionConfig := &models.ActionObj{}
	err = yaml.Unmarshal(buffer, actionConfig)
	if err != nil {
		return Action{Obj: *actionConfig}, err
	}

	if actionConfig.Keygen {
		actionConfig.Key, err = genKey()
		events.CheckError(err)
	} else {
		actionConfig.Key = ""
	}

	action := Action{
		Obj: *actionConfig}

	return action, err
}

func genKey() (string, error) {
	key := make([]byte, 48)

	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(key), nil
}
