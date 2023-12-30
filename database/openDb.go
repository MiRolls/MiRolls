package database

import (
	"MiRolls/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

var Db *sqlx.DB

func Open() {
	mysql := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", config.Configs.Database.Username, config.Configs.Database.Password, config.Configs.Database.Protocol, config.Configs.Database.Host, config.Configs.Database.Port, config.Configs.Database.Database)
	db, err := sqlx.Open("mysql", mysql)
	if err != nil {
		log.Fatal("[FATAL] Database connection pool couldn't run. err: " + err.Error())
	}
	Db = db
}
