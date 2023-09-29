package database

import (
	"fmt"
	"go-nat-project/config"
	models "go-nat-project/models"

	"log"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	Db *gorm.DB
}

var DB = Database{}

func Connect() {
	env := config.ENV
	p := env.DbPort

	port, err := strconv.ParseInt(p, 10, 32)

	log.Println("Databae port is: ", port)

	if err != nil {
		log.Fatal("Error parsing port (str -> int)")
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", env.DbHost, env.DbUser, env.DbPass, env.DbName, port)
	log.Println("Database connection string: ", dsn)
	// postgres connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Error connecting to database")
		// os exit 2 is a general error
		os.Exit(2)
	}

	log.Println("Database connected")
	db.Logger = db.Logger.LogMode(logger.Info)
	log.Println("Database logger set to info")

	log.Println("Database migration started")

	var models = []interface{}{&models.User{}}

	db.AutoMigrate(models...)

	DB = Database{Db: db}

	// defer disconnect(db *gorm.DB){
	// 	// TODO : close database connection
	// 	db.Close()
	// }
}
