// Package models Source
package models

import "github.com/jinzhu/gorm"

// SourcesPerPage limit of Sources pagination
const SourcesPerPage = 30

// Source source struct
type Source struct {
	gorm.Model
	URL    string `form:"url" json:"url" xml:"url" binding:"url,required"`
	Status bool   `form:"status" json:"status" xml:"status"`
	UserID int    `form:"user_id" json:"user_id" xml:"user_id"`
	user   User
}
