package models

type ConfigObj struct {
	Hostname       string                 `yaml:"hostname"`
	ListenPort     int                    `yaml:"listen_port"`
	TlsEnabled     bool                   `yaml:"tls_enabled"`
	TlsCertificate string                 `yaml:"tls_certificate"`
	TlsKey         string                 `yaml:"tls_key"`
	ActionsPath    string                 `yaml:"actions_path"`
	LogEnabled     bool                   `yaml:"log_enabled"`
	LogPath        string                 `yaml:"log_path"`
	FileStore      string                 `yaml:"file_store"`
	CoreActions    map[string]interface{} `yaml:"core_actions"`
}
