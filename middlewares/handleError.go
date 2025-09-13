package middlewares

import (
	"williamsLab/handlers"
	"williamsLab/views/pages"

	"github.com/pocketbase/pocketbase/core"
)

func HandleError(e *core.RequestEvent) error {
	return handlers.Render(e, pages.ErrorPage())
}
