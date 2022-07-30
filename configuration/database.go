package configuration

import (
	"fiber-rest-api/pkg/users"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func DbInit() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("err loaginf .env")
		log.Fatal("Error loading .env file")
	}

	DatabaseName := os.Getenv("DATABASE_NAME")
	DatabaseUSername := os.Getenv("DATABASE_USERNAME")
	DatabasePassword := os.Getenv("DATABASE_PASSWORD")
	DatabaseHost := os.Getenv("DATABASE_HOST")
	DatabasePort := os.Getenv("DATABASE_PORT")

	stringConfig := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DatabaseUSername, DatabasePassword, DatabaseHost, DatabasePort, DatabaseName)

	db, err := gorm.Open(mysql.Open(stringConfig), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database! \n", err.Error())
		os.Exit(2)
	}

	log.Println("Connected to the database successfully")
	db.Logger = logger.Default.LogMode(logger.Info)
	// TODO: Add Migration
	// db.AutoMigrate(&users.User{})

	Database = DbInstance{Db: db}
}

func RunMigrations() {
	err := Database.Db.AutoMigrate(&users.User{})
	if err != nil {
		log.Fatal("err : ", err)
	}

	fmt.Println("Database Migrated")
}
