package requests_test

import (
	"testing"

	"github.com/hackgame-org/fanclub_api/api/requests"
	"github.com/stretchr/testify/assert"
)

func TestPostRequest(t *testing.T) {
	// Instantiate a valid request: without categories and subscriptions
	validReq := requests.PostRequest{
		Title:       "test title",
		Description: nil,
		Price:       nil,
	}

	// Validate valid request
	err := validReq.Validate()
	assert.NoError(t, err)

	// Instantiate a valid request: with categories
	validReq = requests.PostRequest{
		Title:       "test title",
		CategoryIDs: []string{"category1", "category2"},
		Description: nil,
		Price:       nil,
	}

	// Validate valid request
	err = validReq.Validate()
	assert.NoError(t, err)

	// Instantiate a valid request: with subscriptions
	validReq = requests.PostRequest{
		Title:           "test title",
		SubscriptionIDs: []string{"plan1", "plan2"},
		Description:     nil,
		Price:           nil,
	}

	// Validate valid request
	err = validReq.Validate()
	assert.NoError(t, err)

	// Instantiate a valid request: with categories and subscriptions
	validReq = requests.PostRequest{
		Title:           "test title",
		CategoryIDs:     []string{"category1", "category2"},
		SubscriptionIDs: []string{"plan1", "plan2"},
		Description:     nil,
		Price:           nil,
	}

	// Validate valid request
	err = validReq.Validate()
	assert.NoError(t, err)

	// Invalid request: missing title
	invalidReq := requests.PostRequest{
		Title:       "",
		Description: nil,
		Price:       nil,
	}

	// Validate invalid request
	err = invalidReq.Validate()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "title")

	// Invalid request: missing description
	invalidReq = requests.PostRequest{
		Title: "test title",

		Price: nil,
	}

	// Validate invalid request
	err = invalidReq.Validate()
	assert.Nil(t, err)

	// Invalid request: missing price
	invalidReq = requests.PostRequest{
		Title:       "test title",
		Description: nil,
	}

	// Validate invalid request
	err = invalidReq.Validate()
	assert.Nil(t, err)
}
