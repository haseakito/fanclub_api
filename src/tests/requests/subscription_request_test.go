package requests_test

import (
	"testing"

	"github.com/hackgame-org/fanclub_api/requests"
	"github.com/stretchr/testify/assert"
)

func TestSubscriptionRequest(t *testing.T) {
	// Instantiate a valid request
	validReq := requests.SubscriptionRequest{
		Name:       "test name",
		UserID:      "user123",
		Description: nil,
		Price:       nil,
		TrialPeriod: nil,
	}

	// Validate valid request
	err := validReq.Validate()
	assert.NoError(t, err)

	// Invalid request: missing name
	invalidReq := requests.SubscriptionRequest{
		Name:       "",
		UserID:      "user123",
		Description: nil,
		Price:       nil,
		TrialPeriod: nil,
	}

	// Validate invalid request
	err = invalidReq.Validate()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "name")

	// Invalid request: missing user id
	invalidReq = requests.SubscriptionRequest{
		Name:       "test name",
		UserID:      "",
		Description: nil,
		Price:       nil,
		TrialPeriod: nil,
	}

	// Validate invalid request
	err = invalidReq.Validate()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "user_id")

	// Invalid request: missing description
	invalidReq = requests.SubscriptionRequest{
		Name:       "test name",
		UserID:      "user123",
		Price:       nil,
		TrialPeriod: nil,
	}

	// Validate invalid request
	err = invalidReq.Validate()
	assert.Nil(t, err)

	// Invalid request: missing price
	invalidReq = requests.SubscriptionRequest{
		Name:       "test name",
		UserID:      "user123",
		Description: nil,
		TrialPeriod: nil,
	}

	// Validate invalid request
	err = invalidReq.Validate()
	assert.Nil(t, err)

	// Invalid request: missing trial period
	invalidReq = requests.SubscriptionRequest{
		Name:       "test name",
		UserID:      "user123",
		Description: nil,
		Price:       nil,
	}

	// Validate invalid request
	err = invalidReq.Validate()
	assert.Nil(t, err)
}
