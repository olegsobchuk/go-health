package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.Register(func(db migrations.DB) error {
		fmt.Println("creating table users...")
		if _, err := db.Exec(`CREATE TABLE users()`); err != nil {
			return err
		}
		_, err := db.Exec(`
      ALTER TABLE users
      ADD COLUMN id SERIAL PRIMARY KEY,
      ADD COLUMN email text,
      ADD COLUMN username text,
      ADD COLUMN enc_password text,
			ADD COLUMN confirmed_at timestamp without time zone,
      ADD COLUMN created_at timestamp without time zone NOT NULL,
      ADD COLUMN updated_at timestamp without time zone NOT NULL;
			ALTER TABLE users ADD CONSTRAINT users_on_email UNIQUE (email);
    `)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping table users...")
		_, err := db.Exec(`DROP TABLE users`)
		return err
	})
}
