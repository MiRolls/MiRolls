package database

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"strconv"
	"strings"
)

// ROLLS 创建Rolls表SQL语句
const ROLLS = "CREATE TABLE `rolls` (`id` int(11) NOT NULL AUTO_INCREMENT, `roll` text, `code` text, `link` text NOT NULL, PRIMARY KEY (`id`)) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4"

// ANSWER 创建answer表SQL语句
const ANSWER = "CREATE TABLE `answer` (`id` int(11) NOT NULL AUTO_INCREMENT, `answer` text, `code` text NOT NULL, PRIMARY KEY (`id`)) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4"

func Init(userName string, password string, ip string, dbName string) error {
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", strconv.Itoa(3306), ")/", dbName, "?charset=utf8"}, "")
	// 建立数据库连接
	db, err := sqlx.Open("mysql", path)
	if err != nil {
		return errors.New("[ERROR] Can't connect to database")
	}
	// 数据库表创建判断
	_, err = db.Query("DROP TABLE IF EXISTS `rolls`")
	if err != nil {
		return err
	}
	// 建立Rolls数据库表
	_, err = db.Query(ROLLS)
	if err != nil {
		return err
	}
	// 同上
	_, err = db.Query(ANSWER)
	if err != nil {
		return err
	}
	// 锁表
	//_, err = db.Query("LOCK TABLES `rolls` WRITE")
	//if err != nil {
	//	return err
	//}
	// 解锁
	//_, err = db.Query("UNLOCK TABLES")
	//if err != nil {
	//	return err
	//}
	//Write db
	_ = db.Close()
	defer func(open sqlx.DB) {
		_ = open.Close()
	}(*db) // Close
	return nil
}
