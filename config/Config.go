package config

import (
	"errors"
	"gopkg.in/yaml.v2"
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
var defaultConfig = new(Config)

//goland:noinspection GoDeprecation
func InitConfig() (bool, int) {
	_, err := os.ReadDir("config")
	if err != nil {
		return false, 0
	}

	configYaml, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		return false, 0
	}
	err = yaml.Unmarshal(configYaml, &Configs)
	if err != nil {
		log.Fatal("[FATAL ERROR]Can't read config.yaml!" + err.Error())
	}
	return true, 1
}

//goland:noinspection GoDeprecation
func MakeConfig() error {
	err := os.MkdirAll("config", 0755)
	if err != nil {
		return err
	}
	_, err = os.Create("./config/config.yaml")
	if err != nil {
		return err
	}
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return errors.New("can't read filepath")
	}
	//WriteFile
	defaultConfigByte, err := yaml.Marshal(defaultConfig)
	if err != nil {
		//return errors.New("cant to byte")
		return err
	}
	err = os.WriteFile(filepath.Join(dir, "config", "config.yaml"), defaultConfigByte, 0644)
	if err != nil {
		return err
	}
	InitConfig()
	//return true, "Success"
	return nil
}

//goland:noinspection GoDeprecation
func ChangeSite(value *Site) error {
	InitConfig() //init Config
	//Change Site Module
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return errors.New("can't read filepath")
	}
	//goland:noinspection GoDeprecation
	defaultConfig.Site = value
	defaultConfigByte, err := yaml.Marshal(&defaultConfig)
	err = os.WriteFile(filepath.Join(dir, "config", "config.yaml"), defaultConfigByte, 0644)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

//goland:noinspection GoDeprecation
func ChangeDatabase(value *Database) error {
	InitConfig() //init Config
	//Change Site Module
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return errors.New("can't read filepath")
	}
	//goland:noinspection GoDeprecation
	defaultConfig.Database = value
	defaultConfigByte, err := yaml.Marshal(&defaultConfig)
	err = os.WriteFile(filepath.Join(dir, "config", "config.yaml"), defaultConfigByte, 0644)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func ChangeServer(value *Server) error {
	InitConfig() //init Config
	//Change Site Module
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return errors.New("can't read filepath")
	}
	//goland:noinspection GoDeprecation
	defaultConfig.Server = value
	defaultConfigByte, err := yaml.Marshal(&defaultConfig)
	err = os.WriteFile(filepath.Join(dir, "config", "config.yaml"), defaultConfigByte, 0644)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func Destroy() error {
	return os.Remove("config/config.yaml")
}
