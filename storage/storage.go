package storage

import (
	"database/sql"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DB() (*gorm.DB, *sql.DB, error) {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + name + "?parseTime=True"
	db, err := gorm.Open(mysql.New(mysql.Config{DSN: dsn}), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	con, err := db.DB()
	if err != nil {
		return nil, nil, err
	}

	return db, con, err
}
