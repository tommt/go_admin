package database

import (
	"database/sql"
	"go_admin/src/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var sqlDB *sql.DB

func Connect() {
	dsn := "host=db user=postgres password=postgres dbname=postgres port=5432 sslmode=disable "
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	sqlDB, err := DB.DB()
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(models.User{})
}

// AutoMigrateDB will keep tables reflecting structs
// func AutoMigrateDB() {
// 	// if tables exists check if they reflects struts
// 	if err := DB.AutoMigrate(models.User{}); err != nil {
// 		panic(err)
// 	}

// }
