package model

import (
	"fmt"

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

// func init() {
// 	db, err := connectDb()
// 	DbConnection = db
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	// Auto Migrate
// 	if err := db.AutoMigrate(&Video{}); err != nil {
// 		log.Fatalln(err)
// 	}
// 	if err := db.AutoMigrate(&View{}); err != nil {
// 		log.Fatalln(err)
// 	}
// 	if err := db.AutoMigrate(&Rate{}); err != nil {
// 		log.Fatalln(err)
// 	}

// 	// Drop table
// 	if err := db.Migrator().DropTable(&Video{}); err != nil {
// 		log.Fatalln(err)
// 	}

// 	// Create table
// 	if !db.Migrator().HasTable(&Video{}) {
// 		if err := db.Migrator().CreateTable(&Video{}); err != nil {
// 			log.Fatalln(err)
// 		}
// 	}
// 	if !db.Migrator().HasTable(&View{}) {
// 		if err := db.Migrator().CreateTable(&View{}); err != nil {
// 			log.Fatalln(err)
// 		}
// 	}
// 	if !db.Migrator().HasTable(&Rate{}) {
// 		if err := db.Migrator().CreateTable(&Rate{}); err != nil {
// 			log.Fatalln(err)
// 		}
// 	}

// 	for id := 1; id <= 4; id++ {
// 		content := Video{}
// 		content.Title = fmt.Sprintf("video %d", id)
// 		content.CreatedAt = time.Now()
// 		content.UpdatedAt = time.Now()
// 		db.Create(&content)
// 	}
// }
