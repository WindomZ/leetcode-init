package leetcode

import (
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"
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
	// golang language
	LanguageGo LanguageType = "go"
)

// Problem the struct of leetcode problem.
type Problem struct {
	Question
	Language    LanguageType
	Description string `json:"description"`
	Difficulty  string `json:"difficulty"`
}

// Parse parses URL and constructs.
func (p *Problem) Parse() error {
	if p.URL == "" && p.TitleSlug == "" {
		return errors.New("can't find the problem")
	}

	// URL
	if p.URL == "" {
		p.URL = fmt.Sprintf(
			"https://leetcode.com/problems/%s/description/",
			p.TitleSlug,
		)
	}

	doc, err := goquery.NewDocument(p.URL)
	if err != nil {
		return err
	}
	return p.parseDoc(doc)
}

func (p *Problem) parseDoc(doc *goquery.Document) (err error) {
	if err = p.Question.parseDoc(doc); err != nil {
		return
	}

	// TitleSlug
	if p.TitleSlug == "" {
		if strings.HasPrefix(p.URL, "https://leetcode.com/problems/") {
			p.TitleSlug = p.URL[30:]
			p.TitleSlug = strings.TrimSpace(p.TitleSlug[:strings.Index(p.TitleSlug, "/")])
		}
	}

	// Id & Title
	if p.Title == "" {
		p.Title = doc.Find(".question-title .row h3").First().Text()
		if p.Title != "" {
			idx := strings.Index(p.Title, ".")
			p.ID = strings.TrimSpace(p.Title[:idx])
			p.Title = strings.TrimSpace(p.Title[idx+1:])
		}
	}

	// Description
	p.Description, err = doc.Find("div.question-description").First().Html()
	if err != nil {
		return
	}
	p.Description = strings.TrimSpace(p.Description)

	// Difficulty
	p.Difficulty = doc.Find("span.difficulty-label").First().Text()
	p.Difficulty = strings.TrimSpace(p.Difficulty)

	return
}

// ReadMe convert description to markdown.
func (p Problem) ReadMe() string {
	return html2md.Convert(p.Description)
}

func (p Problem) dirName() string {
	return strings.ToLower(strings.Replace(p.TitleSlug, "-", "_", -1))
}

func (p Problem) packageName() string {
	return strings.ToLower(strings.Replace(p.TitleSlug, "-", "", -1))
}

func (p Problem) ensureDir() error {
	return path.Ensure(filepath.Join(".", p.dirName()), true)
}

// OutputReadMe save to README.md.
func (p Problem) OutputReadMe() error {
	if err := p.ensureDir(); err != nil {
		return err
	}
	if p.ID == "" || p.Title == "" {
		return errors.New("not found the language description")
	}
	return path.OverwriteFile(
		filepath.Join(".", p.dirName(), "README.md"),
		fmt.Sprintf("# %s. %s", p.ID, p.Title), "",
		"## Description", "",
		p.ReadMe(),
	)
}

// OutputCode save to src code file with language.
func (p Problem) OutputCode() error {
	code := p.Codes.Code(p.Language.String())
	if code == nil {
		return errors.New("not found the language code")
	}
	if err := p.ensureDir(); err != nil {
		return err
	}
	return code.outputCode(p.dirName(), p.packageName(), p.Language)
}

// OutputTestCode save to test code file with language.
func (p Problem) OutputTestCode() error {
	code := p.Codes.Code(p.Language.String())
	if code == nil {
		return errors.New("not found the language code")
	}
	if err := p.ensureDir(); err != nil {
		return err
	}
	return code.outputTestCode(p.dirName(), p.packageName(), p.Language)
}

// String returns a string.
func (p Problem) String() string {
	b, _ := json.Marshal(p)
	return string(b)
}

// NewProblem returns new Problem instance with a url string.
func NewProblem(lang LanguageType, uri string) *Problem {
	return &Problem{
		Question: Question{
			URL: fmt.Sprintf("https://leetcode.com/problems/%s/description/",
				mustFindFirstStringSubmatch("leetcode.com/problems/([^/]+)", uri)),
		},
		Language: lang,
	}
}

// NewProblemByTitle returns new Problem instance with a title string.
func NewProblemByTitle(lang LanguageType, title string) *Problem {
	title = strings.Replace(strings.TrimSpace(strings.ToLower(title)),
		" ", "-", -1)
	return &Problem{
		Question: Question{
			TitleSlug: title,
		},
		Language: lang,
	}
}
