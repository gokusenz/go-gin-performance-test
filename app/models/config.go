package models

// Config struct
type Config struct {
	Database DatabaseConf `yaml:"database"`
}

// DatabaseConf struct
type DatabaseConf struct {
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	AppName  string `yaml:"app_name"`
}
