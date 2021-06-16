package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() {
	dsn := "host=db user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Kathmandu"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	sqlDB, err := db.DB()
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.

	if err != nil {
		panic("Could not connect with the database")
	}
	sqlDB.SetMaxIdleConns(10)
}
