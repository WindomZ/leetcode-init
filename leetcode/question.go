package leetcode

import (
	"encoding/json"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Question the struct of leetcode question.
type Question struct {
	URL            string `json:"url"`
	Id             string `json:"id"`
	TitleSlug      string `json:"title_slug"`
	Title          string `json:"title"`
	CodeDefinition string `json:"-"`
	Codes          Codes  `json:"codes"`
}

// Parse parses URL and constructs.
func (q *Question) Parse() error {
	doc, err := goquery.NewDocument(q.URL)
	if err != nil {
		return err
	}
	return q.parseDoc(doc)
}

func (q *Question) parseDoc(doc *goquery.Document) error {
	var pageData string
	doc.Find("script").EachWithBreak(func(i int, selection *goquery.Selection) bool {
		pageData = selection.Text()
		if strings.Contains(pageData, "var pageData") {
			pageData = strings.TrimSpace(selection.Text())
			return false
		}
		return true
	})

	// Id
	q.Id = regexp.MustCompile(`questionId: '(\d+)',`).
		FindStringSubmatch(pageData)[1]

	// TitleSlug
	q.TitleSlug = regexp.MustCompile(`questionTitleSlug: '(.+)',`).
		FindStringSubmatch(pageData)[1]

	// Title
	q.Title = regexp.MustCompile(`questionTitle: '(.+)',`).
		FindStringSubmatch(pageData)[1]

	// CodeDefinition
	q.CodeDefinition = regexp.MustCompile(`codeDefinition: (.+),`).
		FindStringSubmatch(pageData)[1]

	return q.parseCode()
}

func (q *Question) parseCode() error {
	if q.CodeDefinition == "" {
		return nil
	}

	q.CodeDefinition = strings.Replace(q.CodeDefinition,
		`'`, `"`, -1)
	q.CodeDefinition = q.CodeDefinition[:len(q.CodeDefinition)-2] + "]"

	if err := json.Unmarshal([]byte(q.CodeDefinition), &q.Codes); err != nil {
		return err
	}

	return nil
}

// String returns a string.
func (q Question) String() string {
	b, _ := json.Marshal(q)
	return string(b)
}

// NewQuestion returns new Question impl.
func NewQuestion(uri string) *Question {
	return &Question{
		URL: uri,
	}
}
