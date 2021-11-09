package models

type Config struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	Port     string `json:"port"`
	Database string `json:"database"`
}
