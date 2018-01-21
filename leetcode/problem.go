package leetcode

import (
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/WindomZ/go-develop-kit/path"
	"github.com/lunny/html2md"
)

// LanguageType type of language
type LanguageType string

// String returns a string
func (l LanguageType) String() string {
	return string(l)
}

const (
	// LanguageGo golang language
	LanguageGo LanguageType = "golang"
)

// Problem the struct of leetcode problem.
type Problem struct {
	Question Question
	Language LanguageType `json:"language"`
	Markdown string       `json:"markdown"`
}

// Parse parses URL and constructs.
func (p *Problem) Parse() error {
	if p.Question.TitleSlug == "" {
		return errors.New("can't find the problem")
	}

	if err := p.Question.Parse(); err != nil {
		return err
	}

	if !p.Question.Valid() {
		return errors.New("can't parse to a question")
	}

	return p.ensureDir()
}

// ReadMe convert description to markdown.
func (p Problem) ReadMe() string {
	return html2md.Convert(p.Question.Content)
}

func (p Problem) dirName() string {
	return strings.ToLower(strings.Replace(p.Question.TitleSlug,
		"-", "_", -1))
}

func (p Problem) packageName() string {
	return strings.ToLower(strings.Replace(p.Question.TitleSlug,
		"-", "", -1))
}

func (p Problem) ensureDir() error {
	return path.Ensure(filepath.Join(".", p.dirName()), true)
}

// OutputReadMe save to README.md.
func (p Problem) OutputReadMe() error {
	if !p.Question.Valid() {
		return errors.New("not found the language description")
	}
	return path.OverwriteFile(
		filepath.Join(".", p.dirName(), "README.md"),
		fmt.Sprintf("# [%s. %s](%s)",
			p.Question.QuestionID, p.Question.QuestionTitle, p.Question.Referer), "",
		"## Description", "",
		p.ReadMe(), "",
		"## Solution",
		fmt.Sprintf("- [Code](%s.go)", p.packageName()),
		fmt.Sprintf("- [Testing](%s_test.go)", p.packageName()), "",
		"## Note", "- [English](NOTE.md)", "- [中文](NOTE_Ch-zh.md)",
	)
}

// OutputCode save to src code file with language.
func (p Problem) OutputCode() error {
	code, err := p.Question.Code(p.Language)
	if err != nil {
		return err
	}
	return code.outputCode(p.dirName(), p.packageName())
}

// OutputTestCode save to test code file with language.
func (p Problem) OutputTestCode() error {
	code, err := p.Question.Code(p.Language)
	if err != nil {
		return err
	}
	return code.outputTestCode(p.dirName(), p.packageName())
}

// OutputMarkdown prints markdown template.
func (p Problem) OutputMarkdown() error {
	if p.Markdown == "" {
		return nil
	}
	return NewMarkdown(p.Markdown, p).outputMarkdown()
}

// String returns a string.
func (p Problem) String() string {
	b, _ := json.Marshal(p)
	return string(b)
}

// NewProblemByURI returns new Problem instance with a url string.
func NewProblemByURI(lang LanguageType, uri, markdown string) *Problem {
	return &Problem{
		Question: Question{
			TitleSlug: mustFindFirstStringSubmatch("leetcode.com/problems/([^/]+)", uri),
		},
		Language: lang,
		Markdown: markdown,
	}
}

// NewProblem returns new Problem instance with a title string.
func NewProblem(lang LanguageType, key, markdown string) *Problem {
	return &Problem{
		Question: Question{
			TitleSlug: key,
		},
		Language: lang,
		Markdown: markdown,
	}
}
