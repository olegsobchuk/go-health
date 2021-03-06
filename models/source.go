// Package models Source
package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/olegsobchuk/go-health/configs"
)

// SourcesPerPage limit of Sources pagination
const (
	SourcesPerPage = 30
	limit          = 1000
)

// Source source struct
type Source struct {
	ID        uint
	URL       string     `form:"url" json:"url" xml:"url" binding:"required"`
	Status    bool       `form:"status" json:"status" xml:"status" binding:"exists"`
	UserID    uint       `form:"user_id" json:"user_id" xml:"user_id"`
	CreatedAt time.Time  `form:"-"`
	UpdatedAt time.Time  `form:"-"`
	DeletedAt *time.Time `form:"-"`
}

// Scopes

// AllActive gets all active sources
func AllActive(db *gorm.DB) *gorm.DB {
	return db.Model(&Source{}).Where("status = ?", true)
}

// Instance functions

// Create save instance to DB
func (source *Source) Create() (*gorm.DB, error) {
	result := configs.DB.Create(source)
	return result, nil
}

// CountActive counts value of active Sources
func (source *Source) CountActive() int {
	var count int
	configs.DB.Model(&source).Where("status = ?", true).Count(&count)
	return count
}

// AddToKVStorage adds Sources to Key-Value storage
func (source *Source) AddToKVStorage() {
	var count int
	configs.DB.Scopes(AllActive).Count(&count)
	var sources []Source
	rounds := count / limit
	for i := 0; i <= rounds; i++ {
		configs.DB.Debug().Scopes(AllActive).Select([]string{"id", "url"}).Offset(i * limit).Limit(limit).Find(&sources)
		fmt.Printf("%+v\n", sources)
		for _, source := range sources {
			configs.KVRegisterSource(strID, source.URL)
		}
	}
}
