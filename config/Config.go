package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
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
	Logo      string `yaml:"logo"`
	MainColor string `yaml:"mainColor"`
	Icp       string `yaml:"icp"`
	Lang      string `yaml:"lang"`
	NeedIcp   int    `yaml:"needIcp"`
}

var Configs Config

func InitConfig() (bool, int) {
	configYaml, err := ioutil.ReadFile("./config/config.yaml")
	if err != nil {
		return false, 0
	}
	err = yaml.Unmarshal(configYaml, &Configs)
	if err != nil {
		log.Fatal("[FATAL ERROR]Can't read config.yaml!" + err.Error())
	}
	return true, 1
}

func MakeConfig(content string) (bool, string) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	//filepath.Abs, return(string,err)
	if err != nil {
		return false, "Can't read filepath"
	}
	//WriteFile
	err = ioutil.WriteFile(filepath.Join(dir, "config", "config.yaml"), []byte(content), 0644)
	if err != nil {
		return false, "Can't write config"
	}

	return true, "Success"
}

func ChangeConfig() {

}
