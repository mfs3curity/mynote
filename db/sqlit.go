package db

import (
	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"gorm.io/gorm"
)

var dbClient *gorm.DB

func InitSQLite() {
	var err error
	dbClient, err = gorm.Open(sqlite.Open("mynote.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func Getdb() *gorm.DB {
	return dbClient
}

func CloseDB() {
	s, _ := dbClient.DB()
	s.Close()
}
