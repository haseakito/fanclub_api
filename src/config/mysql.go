package config

import (
	"context"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hackgame-org/fanclub_api/ent"
)

func InitializeDB() (*ent.Client, error) {
	// Establish database connection and initialize Ent client
	client, err := ent.Open("mysql", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Failed opening connection to mysql: %v", err)
	}

	// Run the auto migration
	if err := client.Schema.Create(context.Background()); err != nil {
		defer client.Close()
		log.Fatalf("Failed creating schema resources: %v", err)
	}

	return client, nil
}
