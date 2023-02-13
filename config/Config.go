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
var defaultConfig = `server:
  port: 2333 #server port
  static: "vite/dist" # theme filepath
database: # about database
  username: "MiRolls"
  password: "JyF8KRTMY85zhaDN"
  protocol: "tcp"
  host: "101.42.50.87"
  port: 3306
  database: "mirolls"
site:
  name: "米卷"
  link: "wj.lmfans.cn" # site link and site name
  logo: "https://img.lmfans.cn/i/2023/01/26/uqwc3n.png" # url(will be favicon and logo, .png or .jpg file, 128*128 best)
  mainColor: "rgb(21, 127, 248)" #rgb or #xxxxxxx Your site color
  icp: "鲁ICP备2022023454号-25"
  lang: "zh"# Currently only supports English and Chinese（en and zh）
  needIcp: 1 #icp(if you in China, you maybe blank 1, if you in the USA or other, you can blank 0(this options controls hyperlinks at the footer))`

//goland:noinspection GoDeprecation
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

//goland:noinspection GoDeprecation
func MakeConfig() (bool, string) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	//filepath.Abs, return(string,err)
	if err != nil {
		return false, "Can't read filepath"
	}
	//WriteFile
	err = ioutil.WriteFile(filepath.Join(dir, "config", "config.yaml"), []byte(defaultConfig), 0644)
	if err != nil {
		return false, "Can't write config"
	}

	return true, "Success"
}

func ChangeConfig() {

}
