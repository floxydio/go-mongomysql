package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var MySQLClient *gorm.DB

func ConnectMysql() {
	dsnMySQL := os.Getenv("MYSQL_URL")
	db, err := gorm.Open(mysql.Open(dsnMySQL), &gorm.Config{})

	if err != nil {
		log.Fatal("MySQL Not Connected")
	}
	MySQLClient = db
}
