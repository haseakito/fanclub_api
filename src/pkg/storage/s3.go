package storage

import (
	"context"
	"fmt"
	"mime/multipart"
	"net/url"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func NewS3Client(ctx context.Context) (*s3.Client, error) {
	// Load AWS configuration
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(*aws.String(os.Getenv("AWS_REGION"))))

	// If loading config failed, then throw an error
	if err != nil {
		return nil, fmt.Errorf("unable to load AWS SDK config: %w", err)
	}

	return s3.NewFromConfig(cfg), nil
}

func UploadFile(ctx context.Context, file multipart.File, content_type, key string) (string, error) {
	// Initialize AWS S3 client
	client, _ := NewS3Client(ctx)

	// Upload file to AWS S3
	_, err := client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("BUCKET_NAME")),
		Key:    aws.String(key),
		Body:   file,
		ContentType: aws.String(content_type),
	})

	// If upload failed, then throw an error
	if err != nil {
		return "", fmt.Errorf("failed to upload file to AWS S3: %w", err)
	}

	// Return the CloudFront URL
	return fmt.Sprintf("https://%s/%s", os.Getenv("CLOUDFRONT_DOMAIN"), key), nil
}

func DeleteFile(ctx context.Context, cloudFrontURL string) error {
	// Initialize AWS S3 client
	client, _ := NewS3Client(ctx)

	// Extract key from cloudfront url
	key, err := extractKey(cloudFrontURL)
	if err != nil {
		return err
	}

	// Construct S3 url
	s3URL := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", os.Getenv("BUCKET_NAME"), key)

	// Delete file from AWS S3
	_, err = client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(os.Getenv("BUCKET_NAME")),
		Key:    aws.String(s3URL),
	})

	// If deletion failed, then throw an error
	if err != nil {
		return fmt.Errorf("failed to delete file from AWS S3: %w", err)
	}

	return nil
}

func extractKey(cfURL string) (string, error) {
	// Parse the cloudfront url
	parsedURL, err := url.Parse(cfURL)
	if err != nil {
		return "", err
	}

	// Remove the leading '/' from the path
	key := strings.TrimPrefix(parsedURL.Path, "/")
	return key, nil
}
