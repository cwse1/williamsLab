package middlewares

import (
	"williamsLab/handlers"
	"williamsLab/models"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
)

func ValidatePublication(e *core.RequestEvent) error {
	pub := models.Publication{}
	err := e.App.DB().Select("pmid", "abstract").From("publications").Where(dbx.Like("pmid", e.Request.PathValue("path")).Match(false, false)).One(&pub)
	if pub.PMID == e.Request.PathValue("path") && pub.Abstract != "" {
		return e.Next()
	} else {
		return handlers.HandleError(e, e.NotFoundError("invalid publication", err))
	}
}
