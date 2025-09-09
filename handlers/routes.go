package handlers

import (
	"williamsLab/services"
	"williamsLab/views/pages"

	"github.com/a-h/templ"
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
	return Render(e, pages.PublicationsPage())
}

func (h RouteHandler) About(e *core.RequestEvent) error {
	return Render(e, pages.AboutPage())
}
