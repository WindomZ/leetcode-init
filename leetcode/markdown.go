package leetcode

import (
	"bytes"
	"path/filepath"
	"text/template"

	"github.com/WindomZ/go-develop-kit/path"
)

// Markdown the struct of leetcode markdown template manager.
type Markdown struct {
	File    string  `json:"-"`
	Problem Problem `json:"-"`
}

// NewMarkdown returns new Markdown impl.
func NewMarkdown(file string, p Problem) *Markdown {
	return &Markdown{
		File:    file,
		Problem: p,
	}
}

// MarkdownFragments the struct of leetcode markdown template.
type MarkdownFragments struct {
	ID          string
	TitleSlug   string
	Title       string
	Language    string
	Difficulty  string
	DirName     string
	PackageName string
}

func (m Markdown) outputMarkdown() error {
	t, err := template.ParseFiles(m.File)
	if err != nil {
		return err
	}

	var buff bytes.Buffer
	if err := t.Execute(&buff, &MarkdownFragments{
		ID:          m.Problem.ID,
		TitleSlug:   m.Problem.TitleSlug,
		Title:       m.Problem.Title,
		Language:    m.Problem.Language.String(),
		Difficulty:  m.Problem.Difficulty,
		DirName:     m.Problem.dirName(),
		PackageName: m.Problem.packageName(),
	}); err != nil {
		return err
	}

	return path.OverwriteFile(
		filepath.Join(".", m.Problem.dirName(), "TEMPLATE.md"),
		buff.String(),
	)
}
