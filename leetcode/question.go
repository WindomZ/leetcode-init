package leetcode

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Question the struct of leetcode question.
type Question struct {
	URL            string `json:"url"`
	ID             string `json:"id"`
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

	if pageData == "" {
		return fmt.Errorf("not found in '%s'", doc.Url.String())
	}

	// ID
	q.ID = mustFindFirstStringSubmatch(`questionId: '(\d+)',`, pageData)

	// TitleSlug
	q.TitleSlug = mustFindFirstStringSubmatch(`questionTitleSlug: '(.+)',`, pageData)

	// Title
	q.Title = mustFindFirstStringSubmatch(`questionTitle: '(.+)',`, pageData)

	// CodeDefinition
	q.CodeDefinition = mustFindFirstStringSubmatch(`codeDefinition: (.+),`, pageData)

	return q.parseCode()
}

func (q *Question) parseCode() error {
	if q.CodeDefinition == "" {
		return nil
	}

	q.CodeDefinition = strings.Replace(q.CodeDefinition,
		`'`, `"`, -1)
	q.CodeDefinition = q.CodeDefinition[:len(q.CodeDefinition)-2] + "]"

	return json.Unmarshal([]byte(q.CodeDefinition), &q.Codes)
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
