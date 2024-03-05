package config

import (
	"log"
	"os"

	"github.com/getsentry/sentry-go"
)

func InitializeSentry() error {
	// Initialize the sentry client with sentry secrets
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              os.Getenv("SENTRY_DSN"),
		TracesSampleRate: 1.0,
	})
	if err != nil {
		log.Fatalf("Failed to initialize sentry client %v", err)
		return err
	}

	return nil
}
