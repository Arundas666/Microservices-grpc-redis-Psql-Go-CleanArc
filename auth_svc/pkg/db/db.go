package db

import (
	"auth_svc/pkg/config"
	"auth_svc/pkg/domain"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Initdb(c *config.Config) (*gorm.DB, error) {
	DBUrl := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", c.DbHost, c.DbUsername, c.DbPassword, c.DbName, c.DbPort)
	db, err := gorm.Open(postgres.Open(DBUrl), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&domain.User{})
	return db, err
}
