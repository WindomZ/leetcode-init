package leetcode

import (
	"encoding/json"

	"github.com/WindomZ/leetcode-graphql"
)

// Question the struct of leetcode question.
type Question struct {
	leetcodegraphql.BaseQuestion
	TitleSlug string `json:"-"`
}

// Parse parses the response to constructs.
func (q *Question) Parse() error {
	return q.Do(q.TitleSlug)
}

// String returns a string.
func (q Question) String() string {
	b, _ := json.Marshal(q)
	return string(b)
}
