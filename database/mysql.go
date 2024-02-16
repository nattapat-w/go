package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mysqlDB struct {
	Db *gorm.DB
}

func NewMySQLDatabase() Database {
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return &mysqlDB{Db: db}
}

// implement database GetDb
func (p *mysqlDB) GetDb() *gorm.DB {
	return p.Db
}
