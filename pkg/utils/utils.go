package utils

import (
	cst "Telegram/pkg/constants"
	"fmt"
	"time"
)

func GetCurrentWeek() int {
	startDate, _ := time.Parse(cst.Layout, fmt.Sprintf("%d-09-01", time.Now().Year()))

	daysDiff := time.Now().Sub(startDate).Hours() / 24
	weeksDiff := int(daysDiff) / 7

	if weeksDiff%2 != 0 {
		return 1
	}
	return 2
}
