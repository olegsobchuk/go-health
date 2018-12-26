// Package models User
package models

import (
  "reflect"
  "time"

  "github.com/gin-gonic/gin/binding"
  "github.com/jinzhu/gorm"
  "github.com/olegsobchuk/go-health/configs"
  "github.com/olegsobchuk/go-health/services/secret"
  validator "gopkg.in/go-playground/validator.v8"
)

// User user struct
type User struct {
  ID                   uint
  Email                string `gorm:"not null;unique" form:"email" json:"email" xml:"email" binding:"required,email,uniquseremail"`
  Username             string `gorm:"column:username" form:"username" json:"username" xml:"username" binding:"required"`
  Password             string `gorm:"-" form:"password" json:"password" xml:"password" binding:"required,min=6,max=24"`
  PasswordConfirmation string `gorm:"-" form:"password_confirmation" json:"password_confirmation" xml:"password_confirmation" binding:"eqfield=Password,required"`
  EncPassword          string `gorm:"column:enc_password"` // encripted password
  ConfirmedAt          time.Time
  UpdatedAt            time.Time  `form:"-"`
  CreatedAt            time.Time  `form:"-"`
  DeletedAt            *time.Time `form:"-"`
  Token                string `json: "token"`
  Sources              []Source
}

func init() {
  // regigtration custom validation
  if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
    v.RegisterValidation("uniquseremail", uniqEmail)
  }
}

// Create creates new user
func (user *User) Create() (*gorm.DB, error) {
  hash, err := secret.Hash(user.Password)
  if err != nil {
    return nil, err
  }
  user.EncPassword = string(hash)
  v := configs.DB.Create(user)
  return v, err
}

// #################
func uniqEmail(
  v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
  field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
  if email, ok := field.Interface().(string); ok {
    user := User{}
    // configs.DB.Where(map[string]interface{}{"email": email}).First(&User{})
    configs.DB.First(&user, map[string]interface{}{"email": email})
    if user.ID > 0 {
      return false
    }
  }
  return true
}
