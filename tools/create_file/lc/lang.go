package lc

import (
	"embed"
	"fmt"
	"path/filepath"
	"strings"
)

type Lang struct {
	ext          string
	name         string
	aliases      []string
	templateName string
}

var langConfigs = []Lang{
	{ext: "c", name: "c", aliases: []string{"c"}, templateName: "c.tpl"},
	{ext: "java", name: "java", aliases: []string{"java"}, templateName: "java.tpl"},
	{ext: "go", name: "golang", aliases: []string{"go", "golang"}, templateName: "go.tpl"},
	{ext: "js", name: "javascript", aliases: []string{"js"}, templateName: "js.tpl"},
	{ext: "ts", name: "typescript", aliases: []string{"ts"}, templateName: "ts.tpl"},
	{ext: "py", name: "python3", aliases: []string{"py"}, templateName: "py.tpl"},
}

//go:embed tpl/*
var templates embed.FS

func loadEmbeddedTemplate(name string) (string, error) {
	bs, err := templates.ReadFile(filepath.Join("tpl", name))
	return string(bs), err
}

func ValidateLang(s string) (Lang, error) {
	s = strings.ToLower(s)
	available := make([]string, 0, len(langConfigs))

	for _, cfg := range langConfigs {
		available = append(available, cfg.ext)
		for _, alias := range cfg.aliases {
			if s == alias {
				return cfg, nil
			}
		}
	}

	return Lang{}, fmt.Errorf("invalid language, must be one of [%s], got %s", available, s)
}
