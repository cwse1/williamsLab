package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		members := core.NewBaseCollection("members")

		members.Fields.Add(
			&core.TextField{
				Name: "name",
			},
			&core.FileField{
				Name:      "picture",
				MaxSelect: 1,
				MimeTypes: []string{"image/avif", "image/webp", "image/png", "image/jpeg"},
			},
			&core.EditorField{
				Name: "description",
			},
		)

		app.Save(members)

		alumni := core.NewBaseCollection("alumni")

		alumni.Fields.Add(
			&core.TextField{
				Name: "name",
			},
			&core.TextField{
				Name: "tenure",
			},
			&core.TextField{
				Name: "position",
			},
		)

		app.Save(alumni)
		return nil
	}, func(app core.App) error {
		// add down queries...

		return nil
	})
}
