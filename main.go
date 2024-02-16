package main

import (
	"github.com/nattapat-w/go/config"
	"github.com/nattapat-w/go/database"
	"github.com/nattapat-w/go/server"
)

func main() {
	cfg := config.GetConfig()
	// db := database.NewMySQLDatabase()
	db := database.NewPostgresDatabase(&cfg)
	server.NewFiberServer(&cfg, db.GetDb()).Start()
}
