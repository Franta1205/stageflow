package initializers

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
	"sync"
)

type Database struct {
	DB *gorm.DB
}

var (
	instance *Database
	mutex    sync.Mutex
)

func ConnectDB() {
	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s port=%s host=%s sslmode=disable",
		LoadEnvVariable("DB_USER"), LoadEnvVariable("PASSWORD"), LoadEnvVariable("DB_NAME"), LoadEnvVariable("PORT"), LoadEnvVariable("HOST"))
	db, err := gorm.Open("postgres", dbInfo)
	if err != nil {
		log.Printf("could not connect to db: %v", err)
		return
	}
	err = db.DB().Ping()
	if err != nil {
		log.Printf("failed to ping db: %v", err)
		return
	}

	instance = &Database{
		DB: db,
	}

	log.Println("Database connected successfully")
}

func CloseDB() {
	mutex.Lock()
	defer mutex.Unlock()

	if instance != nil && instance.DB != nil {
		if err := instance.DB.Close(); err != nil {
			log.Printf("Error closing database connection: %v", err)
		}
		instance = nil
	}
}

func GetDB() *gorm.DB {
	return instance.DB
}
