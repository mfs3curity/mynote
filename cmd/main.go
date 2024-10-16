package main

import (
	"github.com/mfs3curity/mynote/api"
	"github.com/mfs3curity/mynote/db"
	"github.com/mfs3curity/mynote/db/migrations"
)

func main() {
	// init db
	db.InitSQLite()
	// migrations
	migrations.UP()
	api.InitServer()
}
