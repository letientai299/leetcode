package lc

import "strings"

type Lang string

const (
	Java    Lang = "java"
	C       Lang = "c"
	CPP     Lang = "cpp"
	Go      Lang = "golang"
	JS      Lang = "javascript"
	Python  Lang = "python"
	Python3 Lang = "python3"
)

func ValidateLang(s string) (Lang, bool) {
	s = strings.ToLower(s)
	switch {
	case s == string(Go) || s == "go":
		return Go, true
	case s == string(C):
		return C, true
	case s == string(Java):
		return Java, true
	}

	return "", false
}
