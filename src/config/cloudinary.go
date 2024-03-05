package config

import (
	"log"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
)

func InitializeCloudinary() (*cloudinary.Cloudinary, error) {
	// Initialize the cloudinary client with cloudinary secrets
	cld, err := cloudinary.NewFromParams(
		os.Getenv("CLOUDINARY_CLOUD_NAME"),
		os.Getenv("CLOUDINARY_API_KEY"),
		os.Getenv("CLOUDINARY_API_SECRET_KEY"),
	)
	if err != nil {
		log.Fatalf("Failed to initialize cloudinary client %v", err)
		return nil, err
	}

	return cld, nil
}
