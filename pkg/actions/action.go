package actions

import (
	"log"
	"os"
	"tape/pkg/models"
)

type ActionList map[string]Action

type Action struct {
	Obj         models.ActionObj
	CoreHandler func(map[string]string) (bool, string)
}

func (event Action) Authenticate(request string) bool {
	return request == event.Obj.Key
}

func (event Action) Execute(reqData map[string]string) (bool, string) {
	// If this is a core action, use its handler
	if event.CoreHandler != nil {
		return event.CoreHandler(reqData)
	}

	// Execute runs the action's command with optional input and returns success and output/error message.
	cliArg := ""
	if event.Obj.Input {
		val, ok := reqData[event.Obj.Data]
		if !ok {
			log.Printf("Missing required input field: %s", event.Obj.Data)
			return false, "Missing required input field: " + event.Obj.Data
		}
		cliArg = val
	}
	out, err := cmd(event.Obj.Command, cliArg)

	if err != nil {
		log.Printf("Error executing action (%s): %v", event.Obj.Name, err.Error())
		return false, err.Error()
	}

	// Write output to file in 'output' directory if configured
	if event.Obj.OutputWrite && event.Obj.OutputFile != "" {
		outputDir := "output"
		if err := os.MkdirAll(outputDir, 0755); err != nil {
			log.Printf("Failed to create output directory: %v", err)
			return false, "Failed to create output directory: " + err.Error()
		}
		outputPath := outputDir + string(os.PathSeparator) + event.Obj.OutputFile
		fileErr := os.WriteFile(outputPath, out, 0644)
		if fileErr != nil {
			log.Printf("Failed to write output to file %s: %v", outputPath, fileErr)
			return false, "Failed to write output to file: " + fileErr.Error()
		}
	}

	return true, string(out)
}
