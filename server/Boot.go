package server

import (
	"MiRolls/config"
	"MiRolls/database"
	"MiRolls/mainProgram"
	"fmt"
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
	database.Open()
	mainProgram.Run()
}
