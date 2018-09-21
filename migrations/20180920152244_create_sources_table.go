package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	migrations.Register(func(db migrations.DB) error {
		fmt.Println("creating table sources...")
		if _, err := db.Exec(`CREATE TABLE sources()`); err != nil {
			return err
		}
		_, err := db.Exec(`
      ALTER TABLE sources
      ADD COLUMN id SERIAL PRIMARY KEY,
      ADD COLUMN url text,
      ADD COLUMN status boolean,
      ADD COLUMN user_id integer,
      ADD COLUMN created_at timestamp without time zone NOT NULL,
      ADD COLUMN updated_at timestamp without time zone NOT NULL;
    `)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping table sources...")
		_, err := db.Exec(`DROP TABLE sources`)
		return err
	})
}
