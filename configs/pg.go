package configs

import (
	"github.com/go-pg/pg"
)

// DB database connection
var DB *pg.DB

// Init initializes DB connection
func Init() {
	DB = pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "password",
		Database: "health",
		Addr:     "127.0.0.1:5432",
	})
}
