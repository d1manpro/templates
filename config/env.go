package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var envconfig *EnvConfig

const (
	EnvDebug = "DEBUG"

	EnvToken   = "TOKEN"
	EnvAdminID = "ADMIN_ID"

	EnvDBHost     = "DB_HOST"
	EnvDBPort     = "DB_PORT"
	EnvDBName     = "DB_NAME"
	EnvDBUser     = "DB_USER"
	EnvDBPassword = "DB_PASSWORD"
)

type EnvConfig struct {
	Debug bool

	Token   string
	AdminID int

	DB Database
}

type Database struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}

func Load(path string, debug bool) error {
	err := godotenv.Load(path)
	if err != nil {
		return fmt.Errorf("error loading env: %w", err)
	}

	if v, err := strconv.ParseBool(os.Getenv(EnvDebug)); err == nil && v {
		debug = v
	}

	envconfig = &EnvConfig{
		Debug:   debug,
		Token:   requireEnv(EnvToken),
		AdminID: requireEnvInt(EnvAdminID),
		DB: Database{
			Host:     requireEnv(EnvDBHost),
			Port:     requireEnv(EnvDBPort),
			Name:     requireEnv(EnvDBName),
			User:     requireEnv(EnvDBUser),
			Password: requireEnv(EnvDBPassword),
		},
	}
	return nil
}

func requireEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		panic("env " + key + " is required")
	}
	return v
}

func requireEnvInt(key string) int {
	v := os.Getenv(key)
	if v == "" {
		panic("env " + key + " is required")
	}
	i, err := strconv.Atoi(v)
	if err != nil {
		panic("env " + key + " must be integer. Error: " + err.Error())
	}
	return i
}

func Get() *EnvConfig {
	if envconfig == nil {
		panic("config not loaded")
	}
	return envconfig
}
