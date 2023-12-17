package server

import (
	"MiRolls/config"
	"MiRolls/install"
	"MiRolls/mainProgram"
	"fmt"
	"log"
)

func Boot() {
	isSuccess, errCode := config.InitConfig()
	if config.CloudConfigs.Activity == false {
		log.Fatal(fmt.Sprintf("[ERROR] The current back-end version %s is no longer supported by the author. It is recommended that you download the latest version %s!\n", config.BackendVersion, config.CloudConfigs.LatestBackendVersion))
	}
	if !isSuccess && errCode == 0 {
		//Install
		install.Run()
	} else {
		mainProgram.Run()
	}

}
