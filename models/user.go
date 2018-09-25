package models

// User user struct
type User struct {
	ID                   int    `form:"id" json:"id" xml:"id"`
	Email                string `form:"email" json:"email" xml:"email" binding:"required,email"`
	Username             string `form:"username" json:"username" xml:"username" binding:"required"`
	Password             string `form:"password" json:"password" xml:"password" binding:"required,min=6,max=24"`
	PasswordConfirmation string `form:"password_confirmation" json:"password_confirmation" xml:"password_confirmation" binding:"eqfield=Password,required"`
	EncPassword          string // encripted password
}

// Valid validator
// func (user *User) Valid bool {
// user.email
// }
