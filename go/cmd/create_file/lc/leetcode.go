package lc

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/gobuffalo/packr/v2"
	wordwrap "github.com/mitchellh/go-wordwrap"
)

//go:generate packr2 clean
//go:generate packr2

const (
	graphqlURL      = "https://leetcode.com/graphql"
	jsonContentType = "application/json"
	contentWidth    = 77 // 80 minus 3 bytes for comment prefix
)

type Leetcode interface {
	Prepare(url string, lang Lang) (string, error)
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

func (l leetcode) Prepare(url string, lang Lang) (string, error) {
	problem, err := l.downloadProblem(url)
	if err != nil {
		return "", err
	}

	outFileName := problem.FileName(lang)
	if l.fileExist(outFileName) {
		log.Println("file exist: ", outFileName)
		return outFileName, nil
	}

	out, err := l.create(outFileName)
	if err != nil {
		log.Printf("fail to crate, file=%v, err=%v", outFileName, err)
		return "", err
	}

	defer func() { _ = out.Close() }()

	cmt, err := l.prepareComment(problem, url)
	if err != nil {
		return "", err
	}

	code := problem.GetCodeSnippet(lang)
	data := map[string]string{
		"comment": cmt.String(),
		"code":    code,
		"package": problem.PackageName(lang),
	}

	tpl, err := l.loadTplByLang(lang)
	if err != nil {
		return "", err
	}

	if err = tpl.Execute(out, data); err != nil {
		log.Printf("fail to execute template, err=%v", err)
		return "", err
	}

	return outFileName, nil
}

func (l leetcode) prepareComment(problem Problem, url string) (strings.Builder, error) {
	var sb strings.Builder
	l.cmtLine(&sb, problem.Title)
	l.cmtLine(&sb, "")
	l.cmtLine(&sb, problem.Difficulty)
	l.cmtLine(&sb, "")
	l.cmtLine(&sb, url)
	l.cmtLine(&sb, "")

	conv := md.NewConverter("", true, nil)
	content, err := conv.ConvertString(problem.Content)
	if err != nil {
		return sb, err
	}

	rep := strings.NewReplacer("\n", "\n// ", " ", " ")
	for _, s := range strings.Split(content, "\n") {
		s = wordwrap.WrapString(s, contentWidth)
		s = rep.Replace(s)
		l.cmtLine(&sb, s)
	}

	return sb, nil
}

func (l leetcode) extractTitleSlug(url string) string {
	url = strings.TrimSpace(url)
	n := len(url)
	if url[n-1] == '/' {
		n--
		url = url[:n]
	}
	slug := url[strings.LastIndexByte(url, '/')+1:]
	return slug
}

func (l leetcode) downloadProblem(url string) (Problem, error) {
	tpl, err := l.loadTemplate(tplQuestionData)
	if err != nil {
		return Problem{}, err
	}

	titleSlug := l.extractTitleSlug(url)
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

func (l leetcode) loadTplByLang(lang Lang) (*template.Template, error) {
	name := string(lang) + "_file.tpl"
	return l.loadTemplate(name)
}

func (l leetcode) loadTemplate(name string) (*template.Template, error) {
	tplContent, err := getDefaultBox().FindString(name)
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

func getDefaultBox() *packr.Box {
	// this must be string literal, can't use string constant, as the packr2 tool
	// won't be able to recognize it
	// also, the packr2 tool must be executed within this folder,
	return packr.New("default box", "./tpl")
}

func (l leetcode) fileExist(s string) bool {
	_, err := os.Open(s)
	return !os.IsNotExist(err)
}

func (l leetcode) create(name string) (*os.File, error) {
	parent := filepath.Dir(name)
	if err := os.MkdirAll(parent, os.ModePerm); err != nil {
		return nil, err
	}

	return os.Create(name)
}
