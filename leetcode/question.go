package leetcode

import (
	"encoding/json"
	"strings"

	"github.com/WindomZ/leetcode-graphql"
)

// Question the struct of leetcode question.
type Question struct {
	leetcodegraphql.BaseQuestion
	TitleSlug string `json:"-"`
}

// Parse parses the response to constructs.
func (q *Question) Parse() error {
	if err := q.Do(q.TitleSlug); err != nil {
		return err
	}
	q.TitleSlug = strings.Replace(strings.
		TrimSpace(strings.ToLower(q.QuestionTitle)),
		" ", "-", -1)
	return nil
}

// String returns a string.
func (q Question) String() string {
	b, _ := json.Marshal(q)
	return string(b)
}
