package main

import (
	"strings"
)

func reformatDate(date string) string {
	ss := strings.Split(date, " ")
	d, m, y := ss[0], ss[1], ss[2]
	for i, c := range d {
		if c >= '0' && c <= '9' {
			continue
		}
		d = d[:i]
		break
	}

	if len(d) < 2 {
		d = "0" + d
	}

	switch m {
	case "Jan":
		y += "-01-"
	case "Feb":
		y += "-02-"
	case "Mar":
		y += "-03-"
	case "Apr":
		y += "-04-"
	case "May":
		y += "-05-"
	case "Jun":
		y += "-06-"
	case "Jul":
		y += "-07-"
	case "Aug":
		y += "-08-"
	case "Sep":
		y += "-09-"
	case "Oct":
		y += "-10-"
	case "Nov":
		y += "-11-"
	case "Dec":
		y += "-12-"
	}

	return y + d
}
