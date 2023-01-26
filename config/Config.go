package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	Server   *Server   `yaml:"server"`
	Database *Database `yaml:"database"`
	Site     *Site     `yaml:"site"`
}

type Server struct {
	Port   int    `yaml:"port"`
	Static string `yaml:"static"`
}

type Database struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Protocol string `yaml:"protocol"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
}

type Site struct {
	Name      string `yaml:"name"`
	Link      string `yaml:"link"`
	MainColor string `yaml:"mainColor"`
	WithColor string `yaml:"withColor"`
}

var config Config

func InitConfig() {
	configYaml, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatal("[FATAL ERROR]Cannot get config.yaml!" + err.Error())
	}
	err = yaml.Unmarshal(configYaml, &config)
	if err != nil {
		log.Fatal("[FATAL ERROR]Cannot read config.yaml!" + err.Error())
	}
}
