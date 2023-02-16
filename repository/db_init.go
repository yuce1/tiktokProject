package repository

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB = nil
	err error
)

func InitDB(dsn string) error {

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("[FATAL] %s", err)
		os.Exit(-1)
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Printf("[ERROR] %s", err)
		return err
	}

	err = db.AutoMigrate(&Video{})
	if err != nil {
		log.Printf("[ERROR] %s", err)
		return err
	}

	err = db.AutoMigrate(&Comment{})
	if err != nil {
		log.Printf("[ERROR] %s", err)
		return err
	}

	err = db.AutoMigrate(&Relation{})
	if err != nil {
		log.Printf("[ERROR] %s", err)
		return err
	}

	err = db.AutoMigrate(&Favorite{})
	if err != nil {
		log.Printf("[ERROR] %s", err)
		return err
	}

	return nil
}
