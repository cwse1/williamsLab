package middlewares

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pocketbase/pocketbase/core"
)

type ValidateRecord struct{}

func (h ValidateRecord) ValidateMultiFiles(e *core.RecordEvent) error {
	record := e.Record

	contentType := record.GetString("type")

	fileSlice := record.GetUnsavedFiles("image")

	var maxImages int
	switch contentType {
	case "text":
		maxImages = 0
	case "image":
		maxImages = 1
	case "imageLeft":
		maxImages = 1
	}

	if len(fileSlice) > maxImages {
		return validation.NewError("validation_max_select", "Too many images for content type.")
	}

	return e.Next()
}
