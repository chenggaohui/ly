package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	MysqlConf MysqlConf `json:"mysql" yaml:"mysql"`
}

type MysqlConf struct {
	Host string `json:"host" yaml:"host"`
	Port int    `json:"port" yaml:"port"`
	User string `json:"user" yaml:"user"`
	Pwd  string `json:"pwd" yaml:"pwd"`
	Db   string `json:"db" yaml:"db"`
}

var config *Config

func InitConfig() {
	wd, err := os.Getwd()
	configFilePath := wd + "/config/config.yaml"
	println(configFilePath)
	yamlFile, err := os.ReadFile(configFilePath)
	if err != nil {
		panic(err)
	}
	cfg := &Config{}
	err = yaml.Unmarshal(yamlFile, cfg)
	if err != nil {
		panic(err)
	}
	config = cfg
	println(config.MysqlConf.Host)
	println(config.MysqlConf.Port)
	return
}
