package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

var yamlconfig *YamlConfig

type YamlConfig struct {
	Debug bool `yaml:"debug"`

	Token string `yaml:"token"`

	AdminID  int   `yaml:"admin_id"`
	StaffIDs []int `yaml:"staff_ids"`

	Roles map[int]string `yaml:"roles"`

	DB YamlDatabase `yaml:"database"`
}

type YamlDatabase struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func LoadConfig(path string, debug bool) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var cfg *YamlConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return err
	}

	if debug {
		cfg.Debug = debug
	}

	yamlconfig = cfg

	return nil
}

func GetYaml() *YamlConfig {
	if yamlconfig == nil {
		panic("config not loaded")
	}
	return yamlconfig
}
