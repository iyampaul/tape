package actions

import (
	"encoding/json"
	"log"
)

// CoreAction defines a built-in action's configuration and handler
// Handler receives the request data and returns (success, response string)
type CoreAction struct {
	Name        string
	Description string
	Enabled     bool
	Route       string
	Method      string
	Keygen      bool
	Command     string
	Input       bool
	Data        string
	OutputWrite bool
	OutputFile  string
	Key         string
	Handler     func(map[string]string) (bool, string)
}

// GetCoreActions returns all available core actions, configured by config.yml
func GetCoreActions(config map[string]interface{}) map[string]CoreAction {
	coreActions := map[string]CoreAction{}

	logCfg, ok := config["core_actions"].(map[string]interface{})
	if ok {
		logActionCfg, ok := logCfg["log"].(map[string]interface{})
		if ok {
			enabled, _ := logActionCfg["enabled"].(bool)
			method, _ := logActionCfg["method"].(string)
			keygen, _ := logActionCfg["keygen"].(bool)
			input, _ := logActionCfg["input"].(bool)
			data, _ := logActionCfg["data"].(string)
			outputWrite, _ := logActionCfg["output_write"].(bool)
			outputFile, _ := logActionCfg["output_file"].(string)
			route, _ := logActionCfg["route"].(string)
			description, _ := logActionCfg["description"].(string)
			// state := logActionCfg["state"] // Not used in struct, but available

			if enabled {
				var key string
				if keygen {
					k, err := genKey()
					if err != nil {
						key = "None"
					} else {
						key = k
					}
				} else {
					key = ""
				}
				coreActions["log"] = CoreAction{
					Name:        "log",
					Description: description,
					Enabled:     enabled,
					Route:       route,
					Method:      method,
					Keygen:      keygen,
					Command:     "",
					Input:       input,
					Data:        data,
					OutputWrite: outputWrite,
					OutputFile:  outputFile,
					Key:         key,
					Handler: func(req map[string]string) (bool, string) {
						b, _ := json.Marshal(req)
						log.Printf("[core:log] %s", string(b))
						return true, "Logged"
					},
				}
			}
		}
	}
	return coreActions
}
