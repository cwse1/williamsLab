package migrations

import (
	"os"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		settings := app.Settings()

		settings.Meta.AppName = "williamsLab"

		// remove unused `users` collection
		c, err := app.FindCollectionByNameOrId("users")
		if err != nil {
			return err
		}

		err = app.Delete(c)
		if err != nil {
			return err
		}

		// Create initial superuser
		su, err := app.FindCollectionByNameOrId(core.CollectionNameSuperusers)
		if err != nil {
			return err
		}

		user := core.NewRecord(su)
		user.Set("email", os.Getenv("SUPERUSER"))
		user.Set("password", os.Getenv("SUPASS"))

		err = app.Save(user)
		if err != nil {
			return err
		}

		return nil
	}, func(app core.App) error {
		// add down queries...

		return nil
	})
}
