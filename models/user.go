package models

import (
	"time"

	"github.com/olegsobchuk/go-health/configs"
	"github.com/olegsobchuk/go-health/services/secret"
)

// User user struct
type User struct {
	ID                   int    `form:"id" json:"id" xml:"id"`
	Email                string `sql:",unique:idx_users_on_email" form:"email" json:"email" xml:"email" binding:"required,email"`
	Username             string `sql:"username" form:"username" json:"username" xml:"username" binding:"required"`
	Password             string `sql:"-" form:"password" json:"password" xml:"password" binding:"required,min=6,max=24"`
	PasswordConfirmation string `sql:"-" form:"password_confirmation" json:"password_confirmation" xml:"password_confirmation" binding:"eqfield=Password,required"`
	EncPassword          string `sql:"enc_password"` // encripted password
	ConfirmedAt          time.Time
	UpdatedAt            time.Time `sql:"default:now()"`
	CreatedAt            time.Time `sql:"default:now()"`
}

// Create validator
func (user *User) Create() error {
	hash, err := secret.Hash(user.Password)
	if err != nil {
		return err
	}
	user.EncPassword = string(hash)
	time := time.Now()
	user.UpdatedAt = time
	user.CreatedAt = time
	err = configs.DB.Insert(user)
	return err
}
