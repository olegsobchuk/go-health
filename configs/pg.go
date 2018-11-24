package configs

import (
	"github.com/jinzhu/gorm"
	// postgres connection
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB database connection
var DB *gorm.DB

// Init initializes DB connection
func Init() {
	var err error
	DB, err = gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres password=password dbname=health")
	if err != nil {
		panic(err)
	}
}
