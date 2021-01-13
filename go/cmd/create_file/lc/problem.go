package lc

import "strings"

type Problem struct {
	FrontendID string        `json:"questionFrontendId"`
	Title      string        `json:"title"`
	Content    string        `json:"content"`
	Difficulty string        `json:"difficulty"`
	Snippets   []CodeSnippet `json:"codeSnippets"`
}

func (p Problem) GetCodeSnippet(lang Lang) string {
	for _, sp := range p.Snippets {
		if sp.LangSlug == string(lang) {
			return sp.Code
		}
	}

	return ""
}

type CodeSnippet struct {
	LangSlug string `json:"langSlug"`
	Code     string `json:"code"`
}

func (p Problem) FileName(lang Lang) string {
	switch lang {
	case Go:
		return p.goFileName()
	case C:
		return p.cFileName()
	}
	return ""
}

func (p Problem) goFileName() string {
	var sb strings.Builder
	sb.WriteString(p.FrontendID)
	sb.WriteString(".")
	title := strings.ToLower(p.Title)

	for _, s := range strings.Split(title, " ") {
		s = strings.ToLower(s)
		sb.WriteString(s)
		sb.WriteString("-")
	}

	s := sb.String()
	s = s[:len(s)-1]
	s += ".go"
	return s
}

func (p Problem) cFileName() string {
	goName := p.goFileName()
	n := len(goName)
	return goName[:n-3] + ".c"
}
