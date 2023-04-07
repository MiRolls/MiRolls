package database

import (
	"MiRolls/config"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func Open() (error, *sqlx.DB) {
	mysql := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", config.Configs.Database.Username, config.Configs.Database.Password, config.Configs.Database.Protocol, config.Configs.Database.Host, config.Configs.Database.Port, config.Configs.Database.Database)
	db, err := sqlx.Open("mysql", mysql)
	if err != nil {
		return err, nil
	}
	return nil, db
}
