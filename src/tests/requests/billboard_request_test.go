package requests_test

import (
	"testing"

	"github.com/hackgame-org/fanclub_api/requests"
	"github.com/stretchr/testify/assert"
)

func TestBillboardRequest(t *testing.T) {
	// Instantiate a valid request
	validReq := requests.BillboardRequest{
		Title:       "test title",				
	}

	// Validate valid request. Should pass
	err := validReq.Validate()
	assert.NoError(t, err)

	// Invalid request: missing title
	invalidReq := requests.BillboardRequest{
		Description: "test description",		
	}

	// Validate invalid request. Should fail
	err = invalidReq.Validate()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "title")
}
