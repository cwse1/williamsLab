package handlers

import (
	"williamsLab/views/pages"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/router"
)

func HandleError(e *core.RequestEvent, err *router.ApiError) error {
	return Render(e, pages.ErrorPage(err))
}
