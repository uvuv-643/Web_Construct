package internal

import (
	"context"
	"github.com/go-pg/pg/v10"
	"log"
)

func ConnectToDB() *pg.DB {
	db := pg.Connect(&pg.Options{
		User:     "your_user",
		Password: "your_password",
		Database: "your_database",
	})
	if err := db.Ping(context.Background()); err != nil {
		log.Fatalf("Error connecting to DB: %v", err)
	}
	return db
}
