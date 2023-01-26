package main

import (
	"MiRollsBackend/config"
	"MiRollsBackend/server"
)

func main() {
	config.InitConfig()
	server.Boot()
}
