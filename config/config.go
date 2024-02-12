package config

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
)

var (
	Config *JsonConfig
)

var (
	configPath string
	cloudKey   string
)

// init 配置初始化
func init() {
	flag.StringVar(&configPath, "c", "config.json", "config path")
	flag.StringVar(&cloudKey, "cloud", "", "cloudKey")
	_, err := os.Stat(configPath)
	if err != nil {
		log.Panicln("[ERROR] 未找到config.json")
		return
	}
	Reader, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal("[ERROR] 读取config.json失败")
		return
	}
	err = json.Unmarshal(Reader, &Config)
	if err != nil {
		log.Fatal("[ERROR] config.json不是正确的json文件,请检查格式")
		return
	}
}
