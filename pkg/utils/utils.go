package utils

import (
	cst "Telegram/pkg/constants"
	"fmt"
	"time"
)

func GetCurrentWeek() int {
	startDate, _ := time.Parse(cst.FullDateLayout, fmt.Sprintf("%d-09-01", time.Now().Year()))

	now := time.Now().Truncate(24 * time.Hour)
	daysDiff := now.Sub(startDate).Hours() / 24
	weeksDiff := int(daysDiff) / 7

	if weeksDiff%2 != 0 {
		return 2
	}
	return 1
}

func IsDate(date string) bool {
	_, err := time.Parse(cst.DateLayout, date)
	return err == nil
}

func GetWeekDay(date string) int {
	year := time.Now().Year()
	date = fmt.Sprintf("%d-%s", year, date)
	t, _ := time.Parse(cst.FullDateLayout, date)
	if t.Weekday() == 0 {
		return 6
	}
	return int(t.Weekday()) - 1
}
