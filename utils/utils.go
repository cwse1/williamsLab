package utils

import (
	"strconv"
	"time"

	"williamsLab/models"

	"github.com/pocketbase/pocketbase/tools/types"
)

func ParseDate(d string) models.Date {
	parsed, _ := time.Parse(types.DefaultDateLayout, d)
	return models.Date{Year: strconv.Itoa(parsed.Year()), Month: strconv.Itoa(int(parsed.Month())), Day: strconv.Itoa(parsed.Day())}
}
