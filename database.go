package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(100)"`
	UserID   int64  `gorm:"unique;not null"`
}

func (User) TableName() string {
	return "marmitaz.public." + DATABASE_NAME
}

func initializeDatabase() *gorm.DB {

	// dbname set on TableName
	dbUri := fmt.Sprintf("host=%s user=%s port=5432 dbname='' sslmode=disable password=%s", DATABASE_HOST, DATABASE_USER,
		DATABASE_PASSWORD) //Build connection string
	db, err := gorm.Open("postgres", dbUri)
	if err != nil {
		log.Fatalf("Failed to connect to database %s\n%s", DATABASE_HOST, err.Error())
	}

	// Migrate the schema
	db.AutoMigrate(&User{})
	return db
}
