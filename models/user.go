package models

// User user struct
type User struct {
	Email                string `form:"email" json:"email" xml:"email" binding:"required"`
	Username             string
	Password             string `form:"password" json:"password" xml:"password" binding:"required"`
	PasswordConfirmation string `form:"password_confirmation" json:"password_confirmation" xml:"password_confirmation" binding:"required"`
	EncPassword          string // encripted password
}
