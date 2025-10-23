package utils

import (
	"context"
	"io"
	"time"

	"williamsLab/models"

	"github.com/a-h/templ"
	"github.com/pocketbase/pocketbase/tools/types"
)

func ParseDate(d string) models.Date {
	parsed, _ := time.Parse(types.DefaultDateLayout, d)
	return models.Date{Year: parsed.Format("2006"), Month: parsed.Format("Jan"), Day: parsed.Format("02")}
}

func ParseHTML(h string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		_, err = io.WriteString(w, h)
		return
	})
}
