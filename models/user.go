package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/olegsobchuk/go-health/configs"
	"github.com/olegsobchuk/go-health/services/secret"
)

// User user struct
type User struct {
	gorm.Model
	Email                string `gorm:"not null;unique" form:"email" json:"email" xml:"email" binding:"required,email"`
	Username             string `gorm:"column:username" form:"username" json:"username" xml:"username" binding:"required"`
	Password             string `gorm:"-" form:"password" json:"password" xml:"password" binding:"required,min=6,max=24"`
	PasswordConfirmation string `gorm:"-" form:"password_confirmation" json:"password_confirmation" xml:"password_confirmation" binding:"eqfield=Password,required"`
	EncPassword          string `gorm:"column:enc_password"` // encripted password
	ConfirmedAt          time.Time
	Sources              []Source
}

// Create validator
func (user *User) Create() (*gorm.DB, error) {
	hash, err := secret.Hash(user.Password)
	if err != nil {
		return nil, err
	}
	user.EncPassword = string(hash)
	fmt.Printf("%+v\n", configs.DB)
	v := configs.DB.Create(user)
	return v, err
}
