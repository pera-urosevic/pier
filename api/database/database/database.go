package database

import (
	"database/sql"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func DB() *sql.DB {
	if db == nil {
		cfg := mysql.Config{
			User:                 os.Getenv("DB_USER"),
			Passwd:               os.Getenv("DB_PASSWORD"),
			Net:                  "tcp",
			Addr:                 os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
			DBName:               os.Getenv("DB_NAME"),
			AllowNativePasswords: true,
			ParseTime:            true,
		}
		var err error
		db, err = sql.Open("mysql", cfg.FormatDSN())
		if err != nil {
			panic(err)
		}
	}
	return db
}
