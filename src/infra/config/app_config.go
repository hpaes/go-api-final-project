package config

import (
	"github.com/spf13/viper"
)

type (
	AppConfig struct {
		Application ApplicationConfig
		MySQL       MySQL
	}
	ApplicationConfig struct {
		Name   string
		Server ServerConfig
	}

	MySQL struct {
		Host     string
		Port     string
		User     string
		Password string
		DBName   string
	}

	ServerConfig struct {
		Port    string
		Timeout string
	}
)

func NewAppConfig() *viper.Viper {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("app")

	envVars := map[string]string{
		"name":           "APPLICATION_NAME",
		"server.port":    "SERVER_PORT",
		"server.timeout": "SERVER_TIMEOUT",
		"mysql.host":     "DB_HOST",
		"mysql.port":     "DB_PORT",
		"mysql.user":     "DB_USER",
		"mysql.password": "DB_PASSWORD",
		"mysql.dbname":   "DB_NAME",
	}

	for key, env := range envVars {
		viper.BindEnv(key, env)
	}

	return viper.GetViper()
}

func LoadConfig() (*AppConfig, error) {
	var appConfig AppConfig

	viper := NewAppConfig()

	appConfig.Application.Name = viper.GetString("name")
	appConfig.Application.Server.Port = viper.GetString("server.port")
	appConfig.Application.Server.Timeout = viper.GetString("server.timeout")
	appConfig.MySQL.Host = viper.GetString("mysql.host")
	appConfig.MySQL.Port = viper.GetString("mysql.port")
	appConfig.MySQL.User = viper.GetString("mysql.user")
	appConfig.MySQL.Password = viper.GetString("mysql.password")
	appConfig.MySQL.DBName = viper.GetString("mysql.dbname")

	return &appConfig, nil
}
