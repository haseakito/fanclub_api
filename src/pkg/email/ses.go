package email

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
)

var sourceEmailAddress = "info@hackgame.biz"

func NewSESClient(ctx context.Context) (*sesv2.Client, error) {
	// Load AWS configuration
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(*aws.String(os.Getenv("AWS_REGION"))))

	// If loading config failed, then throw an error
	if err != nil {
		return nil, fmt.Errorf("unable to load AWS SDK config: %w", err)
	}

	return sesv2.NewFromConfig(cfg), nil
}

func SendMail(ctx context.Context, to, subject, body string) error {
	// Initialize AWS SES client
	client, err := NewSESClient(ctx)
	if err != nil {
		return fmt.Errorf("failed to create SES client: %w", err)
	}

	// Construct input paramter for SES API
	input := &sesv2.SendEmailInput{
		FromEmailAddress: &sourceEmailAddress,
		Destination: &types.Destination{
			ToAddresses: []string{to},
		},
		Content: &types.EmailContent{
			Simple: &types.Message{
				Subject: &types.Content{
					Data: &subject,
				},
				Body: &types.Body{
					Text: &types.Content{
						Data: &body,
					},
				},
			},
		},
	}

	// Send email via AWS SES
	_, err = client.SendEmail(ctx, input)
	if err != nil {
		return fmt.Errorf("unable to send email: %w", err)
	}

	return nil
}
