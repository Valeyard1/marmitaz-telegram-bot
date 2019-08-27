package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// User ...
// 0 if the person hasn't ordered yet
type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(100)"`
	Order    int    `gorm:"default:0"`
	UserID   int64  `gorm:"unique;not null"`
}

func initializeDatabase() *gorm.DB {

	// dbname set on TableName
	db, err := gorm.Open("sqlite3", DATABASE_HOST)
	if err != nil {
		log.Fatalf("Failed to connect to database %s\n%s", DATABASE_HOST, err.Error())
	}

	// Migrate the schema
	db.AutoMigrate(&User{})
	return db
}
