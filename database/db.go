package database

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// DB interface
type DB interface {
	Connect() *gorm.DB
}

// ConnectDB func
func ConnectDB(db DB) *gorm.DB {
	return db.Connect()
}

// Mysql db
type Mysql struct{}

// Connect mysql
func (d Mysql) Connect() *gorm.DB {
	db, err := gorm.Open("mysql", "root:secret@/golang?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatal("Could not connect, error ", err.Error())
	}

	return db
}

// Postgres db
type Postgres struct{}

// Connect postgres
func (d Postgres) Connect() *gorm.DB {
	db, err := gorm.Open("postgres", "root:secret@/golang?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatal("Could not connect, error ", err.Error())
	}

	return db
}
