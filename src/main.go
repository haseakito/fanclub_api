package main

import (
	"log"

	sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/hackgame-org/fanclub_api/config"
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
	db, _ := config.InitializeDB()

	// Initialize cloudinary client
	cld, _ := config.InitializeCloudinary()

	// Initialize sentry client
	config.InitializeSentry()

	// Initialize echo application
	e := echo.New()

	// Set up CORS config
	e.Use(middleware.CORS())

	// Logging middleware
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
	e.Use(sentryecho.New(sentryecho.Options{}))

	// API v1 group
	r := e.Group("/api/v1")

	// Users group APIs
	u := r.Group("/users")
	{
		uh := handlers.NewUserHandler(db)

		u.GET("", uh.GetUsers)
		u.GET("/:id", uh.GetUser)
		u.GET("/:id/posts", uh.GetPostsByUserID)
		u.PATCH("/:id", uh.UpdateUser)
	}

	// Posts group APIs
	p := r.Group("/posts")
	{
		// Initialize handler for posts
		ph := handlers.NewPostHandler(db, cld)

		p.GET("", ph.GetPosts)
		p.GET("/:id", ph.GetPostByID)
		p.POST("", ph.CreatePost)
		p.POST("/:id/upload", ph.UploadFiles)
		p.PATCH("/:id", ph.UpdatePost)
		p.DELETE("/:id", ph.DeletePost)
		p.DELETE("/:id/assets/:asset_id", ph.DeleteFile)
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

	// Billboards group APIs
	b := r.Group("/billboards")
	{
		bh := handlers.NewBillboardHandler(db, cld)

		b.GET("", bh.GetBillboards)
		b.GET("/:id", bh.GetBillboard)
		b.POST("", bh.CreateBillboard)
		b.POST("/:id/upload", bh.UploadFile)
		b.PATCH("/:id", bh.UpdateBillboard)
		b.DELETE("/:id", bh.DeleteBillboard)
	}

	// Webhook group APIs
	w := r.Group("/webhooks")
	{
		wh := handlers.NewWebhookHandler(db)

		w.POST("/users", wh.ClerkWebhook)
	}

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}
