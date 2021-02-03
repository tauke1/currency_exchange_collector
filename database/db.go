package database

import (
	"currency_exchange_collector/config"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func NewDB() *gorm.DB {
	DBMS := "mysql"
	mySqlConfig := &mysql.Config{
		User:                 config.C.DbUser,
		Passwd:               config.C.DbPassword,
		Net:                  "tcp",
		Addr:                 config.C.DbHost,
		DBName:               config.C.DbName,
		AllowNativePasswords: true,
		Params: map[string]string{
			"parseTime": "true",
			"charset":   "utf8",
		},
	}

	db, err := gorm.Open(DBMS, mySqlConfig.FormatDSN())

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
