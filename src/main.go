package main

import (
	"log"

	"github.com/hackgame-org/fanclub_api/database"
	"github.com/hackgame-org/fanclub_api/handlers"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	// Load environmental variables form .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// Initialize ent client with mysql driver
	db, err := database.InitializeDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Initialize echo application
	e := echo.New()

	// Set up CORS config
	e.Use(middleware.CORS())

	// Set up logger
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())

	// API v1 group
	r := e.Group("/api/v1")

	// Users group APIs
	u := r.Group("/users")
	{
		uh := handlers.NewUserHandler(db)

		u.GET("", uh.GetUsers)
		u.GET("", uh.GetUser)
		u.GET("/:id/posts", uh.GetPostsByUserID)
	}

	// Posts group APIs
	p := r.Group("/posts")
	{
		// Initialize handler for posts
		ph := handlers.NewPostHandler(db)

		p.GET("", ph.GetPosts)
		p.GET("/:id", ph.GetPostByID)
		p.POST("", ph.CreatePost)
		p.PATCH("/:id", ph.UpdatePost)
		p.DELETE("/:id", ph.DeletePost)
	}
	
	// Categories group APIs
	c := r.Group("/categories")
	{
		ch := handlers.NewCategoryHandler(db)
		
		c.GET("", ch.GetCategories)
		c.GET("/:id/posts", ch.GetPostsByCategoryID)
		c.POST("", ch.CreateCategories)
		c.DELETE("/:id", ch.DeleteCategory)
	}

	// Subscriptions group APIs
	s := r.Group("/subscriptions")
	{
		sh := handlers.NewSubscriptionHandler(db)

		s.GET("", sh.GetSubscriptions)
		s.GET("/:id", sh.GetSubscriptions)
		s.POST("", sh.CreateSubscription)
		s.PATCH("/:id", sh.UpdateSubscription)
		s.DELETE("/:id", sh.DeleteSubscription)
	}

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}
