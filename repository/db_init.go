package repository

import (
	"fmt"
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
		fmt.Println(err)
		os.Exit(-1)
	}
	err = db.AutoMigrate(&User{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&Video{})
	if err != nil {
		return err
	}
	return nil
}
