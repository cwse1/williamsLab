package main

import (
	"log"
	"os"

	"williamsLab/handlers"
	"williamsLab/middlewares"
	_ "williamsLab/migrations"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

func main() {
	app := pocketbase.New()

	routes := handlers.RouteHandler{}

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		Automigrate: app.IsDev(),
	})

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		se.Router.GET("/static/{path...}", apis.Static(os.DirFS("./static"), false))

		se.Router.GET("/{$}", routes.Home)
		se.Router.GET("/publications", routes.Publications)
		se.Router.GET("/publications/{path}", routes.PublicationAbstract).BindFunc(middlewares.ValidatePublication)
		se.Router.GET("/about", routes.About)
		se.Router.GET("/research", routes.Research)

		se.Router.GET("/", routes.Error)

		return se.Next()
	})

	err := app.Start()
	if err != nil {
		log.Fatal(err)
	}
}
