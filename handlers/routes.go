package handlers

import (
	"williamsLab/models"
	"williamsLab/services"
	"williamsLab/views/pages"

	"github.com/a-h/templ"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
)

type RouteHandler struct {
	PubService services.PubService
}

func Render(e *core.RequestEvent, t templ.Component) error {
	return t.Render(e.Request.Context(), e.Response)
}

func (h RouteHandler) Home(e *core.RequestEvent) error {
	return Render(e, pages.HomePage())
}

func (h RouteHandler) Research(e *core.RequestEvent) error {
	return Render(e, pages.ResearchPage())
}

func (h RouteHandler) Publications(e *core.RequestEvent) error {
	pubs := []models.Publication{}
	e.App.DB().Select("*").From("publications").All(&pubs)
	return Render(e, pages.PublicationsPage(pubs))
}

func (h RouteHandler) PublicationAbstract(e *core.RequestEvent) error {
	pub := models.Publication{}
	e.App.DB().Select("*").From("publications").Where(dbx.Like("pmid", e.Request.PathValue("path")).Match(false, false)).One(&pub)
	return Render(e, pages.PubAbstract(e.Request.PathValue("path"), pub))
}

func (h RouteHandler) About(e *core.RequestEvent) error {
	return Render(e, pages.AboutPage())
}
