package config

import (
	"log"
	"os"
	"strconv"

	"github.com/inasalifatus/bank-payment/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var PORT int

func LoadEnv(key string) string {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv(key)
}

func InitDB() {
	//databaseURL := os.Getenv("DATABASE_URL")
	var err error
	DB, err = gorm.Open(mysql.Open(LoadEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitMigrate()
}

func InitPort() {
	PORT, _ = strconv.Atoi(LoadEnv("PORT"))

}

func InitMigrate() {
	DB.AutoMigrate(&models.Users{})
}
