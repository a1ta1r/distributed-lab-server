package app

import (
	"github.com/a1ta1r/distributed-lab-server/internal/entity"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var connection *gorm.DB

func GetConnection() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("MYSQL_DSN")

	if connection == nil {
		connection, err = gorm.Open("mysql", dsn)
		connection = connection.Set("gorm:auto_preload", true)
		if err != nil {
			panic(err)
		}
	}
	return connection
}

func UpdateSchema() {
	if connection == nil {
		GetConnection()
	}
	connection.AutoMigrate(
		&entity.Project{},
		&entity.Status{},
		&entity.Task{},
		&entity.Team{},
		&entity.User{},
	)
}
