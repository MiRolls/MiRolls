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
	mysql := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.Config.Key.DB.Username, config.Config.Key.DB.Password, config.Config.Key.DB.Address, config.Config.Key.DB.Port, config.Config.Key.DB.Database)
	db, err := sqlx.Open("mysql", mysql)
	if err != nil {
		log.Fatal("[FATAL] Database connection pool couldn't run. err: " + err.Error())
	}
	Db = db
}
