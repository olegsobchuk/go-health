package configs

import (
	"github.com/jinzhu/gorm"
	// postgres connection
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB database connection
var DB *gorm.DB

// InitDB initializes DB connection
func InitDB() {
	var err error
	DB, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres password=password dbname=health sslmode=disable")
	if err != nil {
		panic(err)
	}
}
