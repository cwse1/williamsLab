package handlers

import (
	"williamsLab/models"
	"williamsLab/views/pages"

	"github.com/a-h/templ"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

type RouteHandler struct {
}

func Render(e *core.RequestEvent, t templ.Component) error {
	return t.Render(e.Request.Context(), e.Response)
}

func (h RouteHandler) Error(e *core.RequestEvent) error {
	return Render(e, pages.ErrorPage(apis.NewNotFoundError("Route Not Found.", nil)))
}

func (h RouteHandler) Home(e *core.RequestEvent) error {
	content, _ := e.App.FindRecordsByFilter("home", "", "+priority", -1, 0)
	if errs := e.App.ExpandRecords(content, []string{"content", "content.images"}, nil); len(errs) > 0 {
		return Render(e, pages.ErrorPage(apis.NewApiError(400, "Failed to fetch page content.", errs)))
	}
	return Render(e, pages.HomePage(content))
}

func (h RouteHandler) Research(e *core.RequestEvent) error {
	content, _ := e.App.FindRecordsByFilter("research", "", "+priority", -1, 0)
	if errs := e.App.ExpandRecords(content, []string{"content", "content.images"}, nil); len(errs) > 0 {
		return Render(e, pages.ErrorPage(apis.NewApiError(400, "Failed to fetch page content.", errs)))
	}
	return Render(e, pages.ResearchPage(content))
}

func (h RouteHandler) Publications(e *core.RequestEvent) error {
	pubs := []models.Publication{}
	e.App.DB().Select("*").From("publications").OrderBy("date DESC").All(&pubs)
	return Render(e, pages.PublicationsPage(pubs))
}

func (h RouteHandler) PublicationAbstract(e *core.RequestEvent) error {
	pub := models.Publication{}
	e.App.DB().Select("*").From("publications").Where(dbx.Like("pmid", e.Request.PathValue("path")).Match(false, false)).One(&pub)
	return Render(e, pages.PubAbstract(e.Request.PathValue("path"), pub))
}

func (h RouteHandler) About(e *core.RequestEvent) error {
	content, _ := e.App.FindRecordsByFilter("about", "", "+priority", -1, 0)
	if errs := e.App.ExpandRecords(content, []string{"content", "content.images"}, nil); len(errs) > 0 {
		return Render(e, pages.ErrorPage(apis.NewApiError(400, "Failed to fetch page content.", errs)))
	}
	members, _ := e.App.FindRecordsByFilter("members", "", "+priority", -1, 0)
	if errs := e.App.ExpandRecords(members, []string{"image"}, nil); len(errs) > 0 {
		return Render(e, pages.ErrorPage(apis.NewApiError(400, "Failed to fetch page content.", errs)))
	}
	alumni, _ := e.App.FindRecordsByFilter("alumni", "", "+priority", -1, 0)
	return Render(e, pages.AboutPage(members, alumni, content))
}
