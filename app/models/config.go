package models

// Config model
type Config struct {
	Database  Database `json:"database"`
	Env       string   `json:"env"`
	ProjectID string   `json:"project_id"`
}

// Database model
type Database struct {
	AppName  string `json:"app_name"`
	Host     string `json:"host"`
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
}
