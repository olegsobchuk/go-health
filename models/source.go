// Package models Source
package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/olegsobchuk/go-health/configs"
)

// SourcesPerPage limit of Sources pagination
const SourcesPerPage = 30

// Source source struct
type Source struct {
	ID        uint
	URL       string `form:"url" json:"url" xml:"url" binding:"url,required"`
	Status    bool   `form:"status" json:"status" xml:"status" binding:"exists"`
	UserID    uint   `form:"user_id" json:"user_id" xml:"user_id"`
	user      User
	CreatedAt time.Time  `form:"-"`
	UpdatedAt time.Time  `form:"-"`
	DeletedAt *time.Time `form:"-"`
}

// Create save instance to DB
func (source *Source) Create() (*gorm.DB, error) {
	result := configs.DB.Create(source)
	return result, nil
}
