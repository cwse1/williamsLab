package middlewares

import (
	"williamsLab/handlers"
	"williamsLab/models"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
)

func ValidatePublication(e *core.RequestEvent) error {
	pub := models.Publication{}
	err := e.App.DB().Select("pmid").From("publications").Where(dbx.Like("pmid", e.Request.PathValue("path")).Match(false, false)).One(&pub)
	if string(pub.PMID) == e.Request.PathValue("path") {
		return e.Next()
	} else {
		return handlers.HandleError(e, e.NotFoundError("invalid publication", err))
	}
}
