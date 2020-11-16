package main

import (
	"time"
)

func daysBetweenDates(date1 string, date2 string) int {
	a, _ := time.Parse("2006-01-02", date1)
	b, _ := time.Parse("2006-01-02", date2)
	if a.Before(b) {
		a, b = b, a
	}
	return int(a.Sub(b).Hours() / 24)
}

