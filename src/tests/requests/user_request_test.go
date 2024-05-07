package requests_test

import (
	"testing"

	"github.com/hackgame-org/fanclub_api/api/requests"
	"github.com/stretchr/testify/assert"
)

func TestUserRequest(t *testing.T) {
	// Instantiate a valid request
	validReq := requests.UserRequest{
		Username: "user-123",
		Bio:      "test user description",
		Url:      "http://twitter.com/",
	}

	// Validate valid request
	err := validReq.Validate()
	assert.NoError(t, err)

	// valid request: missing username
	validReq = requests.UserRequest{
		Bio:      "test user description",
		Url:      "http://twitter.com/",
	}

	// Validate valid request
	err = validReq.Validate()
	assert.NoError(t, err)
	
	invalidReq := requests.UserRequest{
		Username: "user-123",
		Bio:      "test user description",
		Url:      "not a url",
	}

	// Validate invalid request
	err = invalidReq.Validate()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "url")
}
