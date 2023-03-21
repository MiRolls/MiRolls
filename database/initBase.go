package database

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"strconv"
	"strings"
)

const ROLLS = "CREATE TABLE `rolls` (`id` int(11) NOT NULL AUTO_INCREMENT, `roll` text, `code` text, `link` text NOT NULL, PRIMARY KEY (`id`)) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4"
const ANSWER = "CREATE TABLE `answer` (`id` int(11) NOT NULL AUTO_INCREMENT, `answer` text, `link` text NOT NULL, PRIMARY KEY (`id`)) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf9mb4"

func Init(userName string, password string, ip string, dbName string) error {
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", strconv.Itoa(3306), ")/", dbName, "?charset=utf8"}, "")
	db, err := sqlx.Open("mysql", path)
	defer func() {
		_ = db.Close()
	}()
	if err != nil {
		return errors.New("can't connect to database")
	}
	defer func(open sqlx.DB) {
		_ = open.Close()
	}(*db) // Close
	_, _ = db.Query("DROP TABLE IF EXISTS `rolls`")
	_, _ = db.Query(ROLLS)
	_, _ = db.Query(ANSWER)
	_, _ = db.Query("LOCK TABLES `rolls` WRITE")
	_, _ = db.Query("UNLOCK TABLES")
	//Write db
	return nil
}
