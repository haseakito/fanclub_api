package main

import (
	"context"
	"log"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hackgame-org/fanclub_api/api/ent"
)

func main() {
	// Initialize the database client
	client, err := ent.Open("mysql", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}
	defer client.Close()

	if err := GenerateCategories(client); err != nil {
		log.Fatalf("failed to create categories: %v", err)
	}

	// Log the result
	log.Println("All seed data generated successfully.")
}

func GenerateCategories(client *ent.Client) error {
	ctx := context.Background()

	// List of predefined categories
	categories := []string{"Technology", "Health", "Finance", "Education", "Entertainment"}

	// WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Channels to limit concurrent goroutines
	ch := make(chan bool, 5)

	// Create posts for the specified user
	for _, name := range categories {
		wg.Add(1)
		// Start a goroutine
		go func(name string) {
			defer wg.Done()
			ch <- true
			defer func() { <-ch }()

			// Create a new post
			_, err := client.Category.
				Create().
				SetName(name).
				Save(ctx)
			if err != nil {
				log.Printf("Failed to create category '%s': %v", name, err)
			}
		}(name)
	}
	// Wait for all goroutines to complete
	wg.Wait()

	// Log the result of the post creation
	log.Printf("Successfully created categories")

	return nil
}
