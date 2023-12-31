package database

import (
	"fmt"
	"go-nat-project/config"
	user_models "go-nat-project/models"

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
	p := config.Config("DB_PORT")

	port, err := strconv.ParseInt(p, 10, 32)

	log.Println("Databae port is: ", port)

	if err != nil {
		log.Fatal("Error parsing port (str -> int)")
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", config.Config("DB_HOST"), config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"), port)
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

	var models = []interface{}{&user_models.User{}}

	db.AutoMigrate(models...)

	DB = Database{Db: db}
}
