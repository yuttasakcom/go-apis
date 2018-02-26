package database

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Mysql db
type Mysql struct{}

// Connect mysql
func (d Mysql) Connect() *gorm.DB {
	db, err := gorm.Open("mysql", "root:secret@tcp(mysql:3306)/golang?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatal("Could not connect, error ", err.Error())
	}

	return db
}
