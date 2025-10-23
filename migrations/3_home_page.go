package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		homeContent := core.NewBaseCollection("home")

		homeContent.Fields.Add(
			&core.NumberField{
				Name: "priority",
				Required: true,
			},
			&core.RelationField{
				Name:         "content",
				CollectionId: ContentCollectionId,
			},
		)

		app.Save(homeContent)
		return nil
	}, func(app core.App) error {
		// add down queries...

		return nil
	})
}
