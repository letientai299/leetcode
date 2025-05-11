package lc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"text/template"
	"time"

	md "github.com/JohannesKaufmann/html-to-markdown/v2"
	"github.com/mitchellh/go-wordwrap"
	"github.com/pkg/errors"
)

const (
	graphqlURL      = "https://leetcode.com/graphql"
	jsonContentType = "application/json"
	contentWidth    = 77 // 80 minus 3 bytes for comment prefix
)

type Leetcode interface {
	Prepare(url string, lang Lang, force bool) (string, error)
	Read(url string) error
}

func New() Leetcode {
	lc := &leetcode{
		http: &http.Client{Timeout: 5 * time.Second},
	}
	return lc
}

type leetcode struct {
	http *http.Client
}

func (l leetcode) Read(raw string) error {
	u, err := l.parseURL(raw)
	if err != nil {
		return err
	}
	problem, err := l.downloadProblem(u)
	if err != nil {
		return err
	}

	cmt, err := l.prepareComment(problem, u)
	if err != nil {
		return err
	}

	fmt.Println(cmt.String())
	return nil
}

func (l leetcode) parseURL(raw string) (*url.URL, error) {
	u, err := url.Parse(raw)
	if err != nil {
		log.Printf("fail to parse url, raw=%v, err=%v", raw, err)
		return nil, err
	}

	u.RawPath = strings.TrimSuffix(u.Path, "/description/")
	u.Path = u.RawPath
	u.RawQuery = ""
	return u, nil
}

func (l leetcode) Prepare(raw string, lang Lang, force bool) (string, error) {
	u, err := l.parseURL(raw)
	if err != nil {
		return "", err
	}

	problem, err := l.downloadProblem(u)
	if err != nil {
		return "", err
	}

	outFileName := problem.FileName(lang)
	if l.fileExist(outFileName) && !force {
		log.Println("file exist: ", outFileName)
		return outFileName, nil
	}

	out, err := l.create(outFileName)
	if err != nil {
		log.Printf("fail to create, file=%v, err=%v", outFileName, err)
		return "", err
	}

	defer func() { _ = out.Close() }()

	cmt, err := l.prepareComment(problem, u)
	if err != nil {
		return "", err
	}

	code := problem.GetCodeSnippet(lang)
	data := map[string]string{
		"comment": cmt.String(),
		"code":    code,
		"package": problem.PackageName(lang),
	}

	tpl, err := l.loadTemplate(lang.templateName)
	if err != nil {
		return "", err
	}

	if err = tpl.Execute(out, data); err != nil {
		log.Printf("fail to execute template, err=%v", err)
		return "", err
	}

	return outFileName, nil
}

func (l leetcode) prepareComment(problem Problem, u *url.URL) (strings.Builder, error) {
	var sb strings.Builder
	l.cmtLine(&sb, problem.Title+" ("+problem.Difficulty+")")
	l.cmtLine(&sb, u.String())
	l.cmtLine(&sb, "")

	content, err := md.ConvertString(problem.Content)
	if err != nil {
		return sb, err
	}

	rep := strings.NewReplacer("\n", "\n// ", " ", " ")
	ss := strings.Split(content, "\n")
	lines := make([]string, 0, len(ss))

	for _, s := range ss {
		s = wordwrap.WrapString(s, contentWidth)
		s = strings.TrimSpace(rep.Replace(s))
		lines = append(lines, s)
	}

	lines = slices.Compact(lines)

	for _, s := range lines {
		l.cmtLine(&sb, s)
	}

	langs := make([]string, 0, len(problem.Snippets))
	for _, snip := range problem.Snippets {
		langs = append(langs, snip.LangSlug)
	}

	l.cmtLine(&sb, "")

	langLines := "Langs: " + strings.Join(langs, ", ")
	langLines = wordwrap.WrapString(langLines, contentWidth)
	for _, line := range strings.Split(langLines, "\n") {
		l.cmtLine(&sb, line)
	}
	return sb, nil
}

func (l leetcode) extractTitleSlug(u *url.URL) string {
	s := strings.TrimSpace(u.Path)
	n := len(s)
	if s[n-1] == '/' {
		n--
		s = s[:n]
	}
	slug := s[strings.LastIndexByte(s, '/')+1:]
	return slug
}

func (l leetcode) downloadProblem(u *url.URL) (Problem, error) {
	tpl, err := l.loadTemplate(tplQuestionData)
	if err != nil {
		return Problem{}, err
	}

	titleSlug := l.extractTitleSlug(u)
	tplData := map[string]string{
		varTitleSlug: titleSlug,
	}

	var buf bytes.Buffer
	if err = tpl.Execute(&buf, tplData); err != nil {
		log.Printf("fail to execute template, err=%v", err)
		return Problem{}, err
	}

	resp, err := l.http.Post(graphqlURL, jsonContentType, &buf)
	if err != nil {
		log.Printf("fail to send graphql request, err=%v", err)
		return Problem{}, err
	}

	m := make(map[string]map[string]Problem)
	if err = json.NewDecoder(resp.Body).Decode(&m); err != nil {
		log.Printf("fail to decode, err=%v", err)
		return Problem{}, err
	}

	problem := m["data"]["question"]
	return problem, nil
}

func (l leetcode) cmtLine(sb *strings.Builder, s string) {
	if len(s) == 0 {
		sb.WriteString("//")
	} else {
		sb.WriteString("// ")
	}
	sb.WriteString(s)
	sb.WriteString("\n")
}

func (l leetcode) loadTemplate(name string) (*template.Template, error) {
	tplContent, err := loadEmbeddedTemplate(name)
	if err != nil {
		log.Printf("fail to load %v, err=%v", name, err)
		return nil, err
	}

	tpl, err := template.New(name).Parse(tplContent)
	if err != nil {
		log.Printf("fail to parse template, err=%v", err)
		return nil, err
	}

	return tpl, nil
}

func (l leetcode) fileExist(s string) bool {
	_, err := os.Open(s)
	return !os.IsNotExist(err)
}

func (l leetcode) create(name string) (*os.File, error) {
	parent := filepath.Dir(name)
	if err := os.MkdirAll(parent, os.ModePerm); err != nil {
		return nil, errors.Wrapf(err, "fail to create dir: %v", parent)
	}

	return os.Create(name)
}
