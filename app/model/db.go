package model

import (
	"fmt"
	"log"
	"time"

	"ms-api/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DbConnection connect to db using gorm
var DbConnection *gorm.DB

func connectDb() (db *gorm.DB, err error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		config.Config.DbHost,
		config.Config.DbPort,
		config.Config.DbUser,
		config.Config.DbName,
		config.Config.DbPassword,
		config.Config.DbSslMode)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func init() {
	db, err := connectDb()
	DbConnection = db
	if err != nil {
		log.Fatalln(err)
	}

	// Drop table
	if err := db.Migrator().DropTable(&Video{}); err != nil {
		log.Fatalln(err)
	}
	if err := db.Migrator().DropTable(&Analysis{}); err != nil {
		log.Fatalln(err)
	}
	if err := db.Migrator().DropTable(&Rate{}); err != nil {
		log.Fatalln(err)
	}

	// Auto Migrate
	if err := db.AutoMigrate(&Video{}); err != nil {
		log.Fatalln(err)
	}
	if err := db.AutoMigrate(&Analysis{}); err != nil {
		log.Fatalln(err)
	}
	if err := db.AutoMigrate(&Rate{}); err != nil {
		log.Fatalln(err)
	}

	// Create table
	if !db.Migrator().HasTable(&Video{}) {
		if err := db.Migrator().CreateTable(&Video{}); err != nil {
			log.Fatalln(err)
		}
	}
	if !db.Migrator().HasTable(&Analysis{}) {
		if err := db.Migrator().CreateTable(&Analysis{}); err != nil {
			log.Fatalln(err)
		}
	}
	if !db.Migrator().HasTable(&Rate{}) {
		if err := db.Migrator().CreateTable(&Rate{}); err != nil {
			log.Fatalln(err)
		}
	}

	for id := 1; id <= 4; id++ {
		video := Video{
			Title:     fmt.Sprintf("video %d", id),
			UserId:    fmt.Sprintf("user_%d", id),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		db.Create(&video)

		analysis := Analysis{
			UserId:    fmt.Sprintf("user_%d", id),
			VideoId:   video.Id,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		db.Create(&analysis)

		rate := Rate{
			UserId:    fmt.Sprintf("user_%d", id),
			VideoId:   video.Id,
			Value:     3,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		db.Create(&rate)
	}
}
