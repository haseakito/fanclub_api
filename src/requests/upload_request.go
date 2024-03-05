package requests

import (
	"mime/multipart"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func ValidateFile(file *multipart.FileHeader) error {
	// If there is no file, throw an error
	if file == nil {
		return validation.NewError("file", "file is required")
	}

	// Check file size (Max size 2MB limit)
	const maxFileSize = 5 * 1024 * 1024
	if file.Size > maxFileSize {
		return validation.NewError("file", "file size exceeds the limit of 5MB")
	}

	// Allowed file types (jpeg or png)
	allowedTypes := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
	}

	// Check file type
	contentType := file.Header.Get("Content-Type")
	if !allowedTypes[contentType] {
		return validation.NewError("file", "file type must be JPEG or PNG")
	}

	return nil
}

func ValidateVideo(file *multipart.FileHeader) error {
	// If there is no file, throw an error
	if file == nil {
		return validation.NewError("file", "file is required")
	}

	// Check file size (Max size 1GB limit)
	const maxFileSize = 1 * 1024 * 1024 * 1024
	if file.Size > maxFileSize {
		return validation.NewError("file", "file size exceeds the limit of 1GB")
	}

	// Allowed file types (MPEG or MP4)
	allowedTypes := map[string]bool{
		"video/mpeg": true,
		"video/mp4":  true,
	}

	// Check file type
	contentType := file.Header.Get("Content-Type")
	if !allowedTypes[contentType] {
		return validation.NewError("file", "file type must be MPEG  or MP4")
	}

	return nil
}
