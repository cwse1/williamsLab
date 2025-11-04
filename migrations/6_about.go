package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		content := core.NewBaseCollection("about")

		content.Fields.Add(
			&core.NumberField{
				Name:     "priority",
				Required: true,
			},
			&core.RelationField{
				Name:         "content",
				CollectionId: ContentCollectionId,
			},
		)

		app.Save(content)

		members := core.NewBaseCollection("members")

		members.Fields.Add(
			&core.NumberField{
				Name:     "priority",
				Required: true,
			},
			&core.TextField{
				Name:     "name",
				Required: true,
			},
			&core.RelationField{
				Name:         "image",
				CollectionId: ImageCollectionId,
			},
			&core.EditorField{
				Name: "description",
			},
		)

		app.Save(members)

		alumni := core.NewBaseCollection("alumni")

		alumni.Fields.Add(
			&core.NumberField{
				Name:     "priority",
				Required: true,
			},
			&core.TextField{
				Name:     "name",
				Required: true,
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
