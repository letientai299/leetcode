package lc

import (
  "fmt"
  "strconv"
  "strings"
)

type Problem struct {
  FrontendID string        `json:"questionFrontendId"`
  Title      string        `json:"title"`
  Content    string        `json:"content"`
  Difficulty string        `json:"difficulty"`
  Snippets   []CodeSnippet `json:"codeSnippets"`
}

func (p Problem) GetCodeSnippet(lang Lang) string {
  for _, sp := range p.Snippets {
    if sp.LangSlug == lang.name {
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
  switch lang.ext {
  case "go":
    return p.goFileName()
  case "java":
    return p.javaFileName()
  default:
    return p.scriptableFileName(lang.ext)
  }
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

func (p Problem) scriptableFileName(ext string) string {
  goName := p.goFileName()
  n := len(goName)
  return goName[:n-3] + "." + ext
}

func (p Problem) javaFileName() string {
  id, _ := strconv.Atoi(p.FrontendID)
  return fmt.Sprintf("p%04d", id) + "/Solution.java"
}

func (p Problem) PackageName(lang Lang) string {
  switch lang.ext {
  case "java":
    id, _ := strconv.Atoi(p.FrontendID)
    return fmt.Sprintf("p%04d", id)
  default:
    return ""
  }
}
