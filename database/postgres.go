package database

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Postgres db
type Postgres struct{}

// Connect postgres
func (d Postgres) Connect() *gorm.DB {
	db, err := gorm.Open("postgres", "root:secret@postgres/golang?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatal("Could not connect, error ", err.Error())
	}

	return db
}
