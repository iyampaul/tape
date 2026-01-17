// Package models contains data structures for configuration and actions.
package models

// RequestData represents the incoming request data as a map.
type RequestData map[string]string

// ActionObj defines the structure of an action as loaded from YAML.
type ActionObj struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Route       string `yaml:"route"`
	Method      string `yaml:"method"`
	Keygen      bool   `yaml:"generate_keys"`
	Command     string `yaml:"action"`
	Input       bool   `yaml:"accept_input"`
	Data        string `yaml:"data_field"`
	OutputWrite bool   `yaml:"output_write"`
	OutputFile  string `yaml:"output_file"`
	Key         string
}
