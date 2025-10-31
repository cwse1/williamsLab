package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

var ContentCollectionId string

var BlockTypes = []string{
	"text",
	"image",
	"imageLeft",
	"imageRight",
}

func init() {
	m.Register(func(app core.App) error {
		content := core.NewBaseCollection("content")

		content.Fields.Add(
			&core.SelectField{
				Name:     "type",
				Values:   BlockTypes,
				Required: true,
			},
			&core.EditorField{
				Name: "content",
			},
			&core.FileField{
				Name:      "image",
				MaxSelect: 6,
				MimeTypes: []string{"image/avif", "image/webp", "image/png", "image/jpeg"},
			},
		)

		ContentCollectionId = content.Id

		app.Save(content)
		return nil
	}, func(app core.App) error {
		// add down queries...

		return nil
	})
}
