package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

var ContentCollectionId string

func init() {
	m.Register(func(app core.App) error {
		content := core.NewBaseCollection("content")

		content.Fields.Add(
			&core.SelectField{
				Name:   "type",
				Values: []string{"text", "image", "imageLeft"},
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
