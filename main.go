package main

import (
	"MiRolls/config"
	"MiRolls/server"
)

func main() {
	config.InitConfig()
	server.Boot()
}
