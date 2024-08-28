package db

import (
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbConn *gorm.DB

func ConnectToDB() error {
	connStr := "user=postgres password=q123 dbname=postgres host=localhost port=5433 sslmode=disable"

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return err
	}

	fmt.Println("Connected to database")

	dbConn = db
	return nil
}

func CloseDBConn() error {
	// err := dbConn.Close()
	// if err != nil {
	// 	return err
	// }

	return nil
}

func GetDBConn() *gorm.DB {
	return dbConn
}
