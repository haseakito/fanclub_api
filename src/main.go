package main

import (
	"log"
	"os"

	handlers "github.com/hackgame-org/fanclub_api/api/handler"
	middlewares "github.com/hackgame-org/fanclub_api/api/middleware"
	"github.com/hackgame-org/fanclub_api/internal/database"
	"github.com/hackgame-org/fanclub_api/internal/redis"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/muxinc/mux-go"
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
	db, _ := database.Initialize()

	// Initialize redis client
	rdb := redis.Initialize()

	// Initialize Mux client
	mux := muxgo.NewAPIClient(
		muxgo.NewConfiguration(
			muxgo.WithBasicAuth(os.Getenv("MUX_TOKEN_ID"), os.Getenv("MUX_TOKEN_SECRET")),
		))

	// Initialize echo application
	e := echo.New()

	// Set up CORS config
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", os.Getenv("FRONTEND_URL")},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	// Logging middleware
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())

	// API v1 group
	r := e.Group("/api/v1")

	// Auth group APIs
	a := r.Group("/auth")
	{
		ah := handlers.NewAuthHandler(db)

		a.POST("/sign-up", ah.Signup)
		a.POST("/sign-in", ah.Signin)
		a.POST("/verify-email", ah.VerifyEmail)
		a.POST("/verify-email/resend", ah.VerifyEmail)
		a.POST("/reset-password", ah.ResetPassword)
		a.POST("/verify-reset-password", ah.VerifyPasswordReset)
	}

	// Users group APIs
	u := r.Group("/users")
	{
		uh := handlers.NewUserHandler(db, rdb)

		u.POST("/upload/profile_picture", uh.UploadProfilePicture, middlewares.AuthMiddleware)
		u.POST("/:id/follow", uh.FollowUser)
		u.POST("/:id/unfollow", uh.UnfollowUser)
		u.GET("", uh.GetUsers, middlewares.AuthMiddleware)
		u.GET("/:id", uh.GetUser)
		u.GET("/:id/followers", uh.GetFollowers)
		u.GET("/:id/following", uh.GetFollowing)
		u.PATCH("/:id", uh.UpdateUserProfile)
		u.DELETE("", uh.DeleteUser, middlewares.AuthMiddleware)
	}

	// Posts group APIs
	p := r.Group("/posts")
	{
		// Initialize handler for posts
		ph := handlers.NewPostHandler(db, mux, rdb)

		p.POST("", ph.CreatePost)
		p.POST("/:id/upload/video", ph.UploadVideo)
		p.POST("/:id/upload/thumnail", ph.UploadThumnail)
		p.GET("", ph.GetPosts)
		p.GET("/:id", ph.GetPostByID)		
		p.PATCH("/:id", ph.UpdatePost)
		p.DELETE("/:id", ph.DeletePost)
	}

	// Likes group APIs
	l := r.Group("/likes")
	{
		lh := handlers.NewLikeHandler(db)

		l.POST("/create", lh.CreateLike)
		l.POST("/delete", lh.DeleteLike)
	}

	// Categories group APIs
	c := r.Group("/categories")
	{
		ch := handlers.NewCategoryHandler(db)

		c.POST("", ch.CreateCategories)
		c.GET("", ch.GetCategories)
		c.GET("/:id/posts", ch.GetPostsByCategoryID)
		c.DELETE("/:id", ch.DeleteCategory)
	}

	// Subscriptions group APIs
	s := r.Group("/subscriptions")
	{
		sh := handlers.NewSubscriptionHandler(db)

		s.GET("", sh.GetSubscriptions)
		s.GET("/:id", sh.GetSubscription)
		s.POST("", sh.CreateSubscription)
		s.PATCH("/:id", sh.UpdateSubscription)
		s.DELETE("/:id", sh.DeleteSubscription)
	}

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}
