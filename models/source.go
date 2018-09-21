package models

// Source source struct
type Source struct {
	ID     int    `form:"id" json:"id" xml:"id"`
	URL    string `form:"url" json:"url" xml:"url"`
	Status bool   `form:"status" json:"status" xml:"status"`
	UserID int    `form:"user_id" json:"user_id" xml:"user_id"`
	user   *User
}
