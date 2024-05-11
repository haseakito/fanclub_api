package main

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/brianvoe/gofakeit/v6"

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

	// Channels to limit concurrent goroutines
	ch := make(chan bool, 5)

	// WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ch <- true
			defer func() { <-ch }()

			// Seed user data
			user, err := GenerateRandomUser(client)
			if err != nil {
				log.Fatalf("failed to create user: %v", err)
			}

			// Seed posts data for each user
			if err := GenerateRandomPosts(client, user, 10); err != nil {
				log.Fatalf("failed to create post for user: %v", err)
			}
		}()
	}

	if err := GenerateCategories(client); err != nil {
		log.Fatalf("failed to create categories: %v", err)
	}
	// Wait for all goroutines to complete
	wg.Wait()

	// Log the result
	log.Println("All seed data generated successfully.")
}

func GenerateRandomUser(client *ent.Client) (*ent.User, error) {
	ctx := context.Background()

	// Create a new user
	user, err := client.User.
		Create().
		SetName(gofakeit.Name()).
		SetEmail(gofakeit.Email()).
		SetPassword(gofakeit.Password(true, true, true, false, false, 10)).
		SetProfileImageURL(gofakeit.ImageURL(100, 100)).
		Save(ctx)
	// If creation failed, then throw an error
	if err != nil {
		log.Printf("Failed to create user: %v", err)
		return nil, err
	}

	// Log the result of the user creation
	log.Printf("Created user: %v\n", user.ID)

	return user, nil
}

func GenerateRandomPosts(client *ent.Client, user *ent.User, count int) error {
	ctx := context.Background()

	// WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Create posts for the specified user
	for i := 0; i < count; i++ {
		wg.Add(1)
		// Start a goroutine
		go func() {
			defer wg.Done()

			// Create a new post
			_, err := client.Post.
				Create().
				SetTitle(gofakeit.ProductName()).
				SetDescription(gofakeit.Sentence(gofakeit.Number(30, 50))).
				SetPrice(int64(gofakeit.Price(0, 1000))).
				SetThumbnailURL(gofakeit.ImageURL(1000, 1000)).
				SetUser(user).
				Save(ctx)
			if err != nil {
				log.Printf("Failed to create post for user %v: %v", user.ID, err)
			}
		}()
	}
	// Wait for all goroutines to complete
	wg.Wait()

	// Log the result of the post creation
	log.Printf("Created %d posts for user %v\n", count, user.ID)

	return nil
}

func GenerateCategories(client *ent.Client) error {
	ctx := context.Background()

	// List of predefined categories
	categories := []string{"Technology", "Health", "Finance", "Education", "Entertainment"}

	// WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Create posts for the specified user
	for _, name := range categories {
		wg.Add(1)
		// Start a goroutine
		go func(name string) {
			defer wg.Done()

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
	log.Printf("Created categories")

	return nil
}
