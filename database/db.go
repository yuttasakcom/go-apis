package database

import "github.com/jinzhu/gorm"

// DB interface
type DB interface {
	Connect() *gorm.DB
}

// ConnectDB func
func ConnectDB(db DB) *gorm.DB {
	return db.Connect()
}
