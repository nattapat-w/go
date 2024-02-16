package main

import (
	"github.com/nattapat-w/go/config"
	"github.com/nattapat-w/go/database"
	"github.com/nattapat-w/go/order/entities"
)

func main() {
	cfg := config.GetConfig()

	db := database.NewPostgresDatabase(&cfg)

	//drop table if exists
	db.GetDb().Migrator().DropTable(&entities.Order{})

	orderMigrate(db)
}

func orderMigrate(db database.Database) {
	db.GetDb().Migrator().CreateTable(&entities.Order{})
	db.GetDb().CreateInBatches([]entities.Order{
		{Total: 1},
		{Total: 2},
		{Total: 2},
		{Total: 5},
		{Total: 3},
	}, 10)
}
