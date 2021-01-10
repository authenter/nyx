package config

import (
	"fmt"

	"github.com/authenter/nyx/pkg/db"
	"github.com/authenter/nyx/pkg/log"

	"github.com/spf13/viper"
)

// Config object holds all configurations
type Config struct {
	Server   serverConfig
	Log      logConfig
	Database databaseConfig
}

type serverConfig struct {
	Host string
	Port int
}

type logConfig struct {
	Level log.Level
	File  string
}

type databaseConfig struct {
	Username string
	Password string
	Driver   db.Driver
	Host     string
	Port     int
	Name     string
}

// LoadConfig returns filled Config object
func LoadConfig() *Config {
	setDefaults()
	loadFromViper()

	logLevel, err := log.ParseLevel(viper.GetString("log.level"))
	if err != nil {
		panic(fmt.Errorf("failed to parse log level: %s", err))
	}

	dbDriver, err := db.ParseDriver(viper.GetString("database.driver"))
	if err != nil {
		panic(fmt.Errorf("failed to parse database driver: %s", err))
	}

	return &Config{
		Server: serverConfig{
			Host: viper.GetString("server.host"),
			Port: viper.GetInt("server.port"),
		},
		Log: logConfig{
			Level: logLevel,
			File:  viper.GetString("log.file"),
		},
		Database: databaseConfig{
			Username: viper.GetString("database.username"),
			Password: viper.GetString("database.password"),
			Driver:   dbDriver,
			Host:     viper.GetString("database.host"),
			Port:     viper.GetInt("database.port"),
			Name:     viper.GetString("database.name"),
		},
	}
}

func loadFromViper() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to load config file: %s", err))
	}

	viper.WatchConfig()
}
