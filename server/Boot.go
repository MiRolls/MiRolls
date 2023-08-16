package server

import (
	"MiRolls/config"
	"MiRolls/install"
	"MiRolls/mainProgram"
)

func Boot() {

	isSuccess, errCode := config.InitConfig()
	if !isSuccess && errCode == 0 {
		//Install
		install.Run()
	} else {
		mainProgram.Run()
	}

}
