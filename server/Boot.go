package server

import (
	"MiRolls/config"
	"MiRolls/database"
	"MiRolls/install"
	"MiRolls/mainProgram"
	"fmt"
	"log"
)

func Boot() {
	// 声明软件版权
	fmt.Print(`
  __  __   _   ____            _   _       
 |  \/  | (_) |  _ \    ___   | | | |  ___ 
 | |\/| | | | | |_) |  / _ \  | | | | / __|
 | |  | | | | |  _ <  | (_) | | | | | \__ \
 |_|  |_| |_| |_| \_\  \___/  |_| |_| |___/
 V` + config.BackendVersion + `	Commit #` + config.LastCommit + ` StaticVer` + config.RequiredStaticVersion + ` InstallerVer` + config.InstallerVersion + `
`)

	isSuccess, errCode := config.InitConfig()
	if config.CloudConfigs.Activity == false {
		log.Fatal(fmt.Sprintf("[ERROR] The current back-end version %s is no longer supported by the author. It is recommended that you download the latest version %s!\n", config.BackendVersion, config.CloudConfigs.LatestBackendVersion))
	}
	if !isSuccess && errCode == 0 {
		//Install
		install.Run()
	} else {
		database.Open()
		mainProgram.Run()
	}

}
