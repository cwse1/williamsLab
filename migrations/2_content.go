package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

var ContentCollectionId string
var ImageCollectionId string

var BlockTypes = []string{
	"text",
	"image",
	"imageLeft",
	"imageRight",
	"imageTable",
}

func init() {
	m.Register(func(app core.App) error {
		images := core.NewBaseCollection("images")

		images.Fields.Add(
			&core.FileField{
				Name:      "image",
				MaxSelect: 1,
				MimeTypes: []string{"image/avif", "image/webp", "image/png", "image/jpeg"},
			},
			&core.EditorField{
				Name: "caption",
			},
			&core.TextField{
				Name: "alt",
			},
		)

		ImageCollectionId = images.Id

		app.Save(images)

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
			&core.RelationField{
				Name: "images",
				MaxSelect: 6,
				CollectionId: ImageCollectionId,
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
