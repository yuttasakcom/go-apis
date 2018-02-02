package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/yuttasakcom/go-apis/database"
	"golang.org/x/crypto/bcrypt"
)

// User model
type User struct {
	ID        uuid.UUID
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Create user
func (u *User) Create() {
	db := database.ConnectDB(database.Mysql{})

	u.ID = uuid.Must(uuid.NewV4())

	hpwd, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		return
	}

	u.Password = string(hpwd)
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	db.NewRecord(u)
	db.Create(&u)
	db.NewRecord(u)
}
