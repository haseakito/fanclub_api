package requests_test

import (
	"mime/multipart"
	"net/textproto"
	"testing"

	"github.com/hackgame-org/fanclub_api/api/requests"
	"github.com/stretchr/testify/assert"
)

func TestValidateFile(t *testing.T) {
	// Instantiate a valid request: mime type is .jpg
	validReq := &multipart.FileHeader{
		Filename: "test.jpg",
		Size:     1024 * 1024, // 1MB
		Header:   make(textproto.MIMEHeader),
	}
	validReq.Header.Set("Content-Type", "image/jpeg")

	// Validate valid request. Should pass
	err := requests.ValidateFile(validReq)
	assert.NoError(t, err)

	// Instantiate a valid request: mime type is .png
	validReq = &multipart.FileHeader{
		Filename: "test.png",
		Size:     1024 * 1024, // 1MB
		Header:   make(textproto.MIMEHeader),
	}
	validReq.Header.Set("Content-Type", "image/png")

	// Validate valid request. Should pass
	err = requests.ValidateFile(validReq)
	assert.NoError(t, err)

	// Invalid request: file size exceeds limit
	invalidReq := &multipart.FileHeader{
		Filename: "test.jpg",
		Size:     6 * 1024 * 1024, // 6MB
		Header:   make(textproto.MIMEHeader),
	}
	invalidReq.Header.Set("Content-Type", "image/jpeg")

	// Validate invalid request. Should fail
	err = requests.ValidateFile(invalidReq)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "file size exceeds the limit of 5MB")

	// Invalid request: file type is not allowed
	invalidReq = &multipart.FileHeader{
		Filename: "test.pdf",
		Size:     1 * 1024 * 1024, // 1MB
		Header:   make(textproto.MIMEHeader),
	}
	invalidReq.Header.Set("Content-Type", "image/pdf")

	// Validate invalid request. Should fail
	err = requests.ValidateFile(invalidReq)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "file type must be JPEG or PNG")
}

func TestValidateVideo(t *testing.T) {
	// Instantiate a valid request: mime type is .mpeg
	validReq := &multipart.FileHeader{
		Filename: "test.mpeg",
		Size:     5 * 1024 * 1024, // 5MB
		Header:   make(textproto.MIMEHeader),
	}
	validReq.Header.Set("Content-Type", "video/mpeg")

	// Validate valid request. Should pass
	err := requests.ValidateVideo(validReq)
	assert.NoError(t, err)

	// Instantiate a valid request: mime type is .mp4
	validReq = &multipart.FileHeader{
		Filename: "test.mp4",
		Size:     5 * 1024 * 1024, // 1MB
		Header:   make(textproto.MIMEHeader),
	}
	validReq.Header.Set("Content-Type", "video/mp4")

	// Validate valid request. Should pass
	err = requests.ValidateVideo(validReq)
	assert.NoError(t, err)

	// Invalid request: file size exceeds limit
	invalidReq := &multipart.FileHeader{
		Filename: "test.mpeg",
		Size:     2 * 1024 * 1024 * 1024, // 2GB
		Header:   make(textproto.MIMEHeader),
	}
	invalidReq.Header.Set("Content-Type", "video/mpeg")

	// Validate invalid request. Should fail
	err = requests.ValidateVideo(invalidReq)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "file size exceeds the limit of 1GB")

	// Invalid request: file type is not allowed
	invalidReq = &multipart.FileHeader{
		Filename: "test.mov",
		Size:     1 * 1024 * 1024, // 1MB
		Header:   make(textproto.MIMEHeader),
	}
	invalidReq.Header.Set("Content-Type", "video/quicktime")

	// Validate invalid request. Should fail
	err = requests.ValidateVideo(invalidReq)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "file type must be MPEG or MP4")
}