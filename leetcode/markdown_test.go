package leetcode

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func TestNewMarkdown(t *testing.T) {
	assert.NoError(t, testProblem.Parse())
	assert.NoError(t, NewMarkdown("../TEMPLATE.md", *testProblem).outputMarkdown())
}
