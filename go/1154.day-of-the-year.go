package main

import "time"

func dayOfYear(date string) int {
	d, _ := time.Parse("2006-01-02", date)
	return d.YearDay()
}
