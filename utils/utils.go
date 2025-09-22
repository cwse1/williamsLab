package utils

import (
	"time"

	"williamsLab/models"

	"github.com/pocketbase/pocketbase/tools/types"
)

func ParseDate(d string) models.Date {
	parsed, _ := time.Parse(types.DefaultDateLayout, d)
	return models.Date{Year: parsed.Format("2006"), Month: parsed.Format("Jan"), Day: parsed.Format("02")}
}
