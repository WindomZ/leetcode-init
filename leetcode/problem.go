package leetcode

import (
	"errors"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/WindomZ/go-develop-kit/path"
	"github.com/lunny/html2md"
)

type Problem struct {
	Index       int
	URL         string
	Short       string
	Name        string
	Description string
	Difficulty  string
}

func (p *Problem) Parse() error {
	if p.URL == "" && p.Short == "" {
		return errors.New("can't find the problem")
	}
	if p.URL == "" {
		p.URL = fmt.Sprintf(
			"https://leetcode.com/problems/%s/description/",
			p.Short,
		)
	}

	if p.Short == "" {
		if strings.HasPrefix(p.URL, "https://leetcode.com/problems/") {
			p.Short = p.URL[30:]
			p.Short = strings.TrimSpace(p.Short[:strings.Index(p.Short, "/")])
		}
	}

	doc, err := goquery.NewDocument(p.URL)
	if err != nil {
		return err
	}

	p.Name = doc.Find(".question-title .row h3").First().Text()
	if p.Name != "" {
		idx := strings.Index(p.Name, ".")
		p.Index, _ = strconv.Atoi(strings.TrimSpace(p.Name[:idx]))
		p.Name = strings.TrimSpace(p.Name[idx+1:])
	}

	p.Description, err = doc.Find("div.question-description").First().Html()
	if err != nil {
		return err
	}
	p.Description = strings.TrimSpace(p.Description)

	p.Difficulty = doc.Find("span.difficulty-label").First().Text()
	p.Difficulty = strings.TrimSpace(p.Difficulty)

	return nil
}

func (p *Problem) ReadMe() string {
	return html2md.Convert(p.Description)
}

func (p *Problem) OutputReadMe(dir string) error {
	return path.OverwriteFile(
		filepath.Join(dir, "README.md"),
		fmt.Sprintf("# %d. %s", p.Index, p.Name),
		"## description",
		p.ReadMe(),
	)
}
