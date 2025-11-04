package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		researchContent := core.NewBaseCollection("research")

		researchContent.Fields.Add(
			&core.NumberField{
				Name: "priority",
				Required: true,
			},
			&core.TextField{
				Name:   "topic",
				Min: 5,
				Required: true,
			},
			&core.RelationField{
				Name:         "content",
				CollectionId: ContentCollectionId,
			},
		)

		app.Save(researchContent)
		return nil
	}, func(app core.App) error {
		// add down queries...

		return nil
	})
}
