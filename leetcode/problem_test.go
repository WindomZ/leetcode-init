package leetcode

import (
	"testing"

	"github.com/WindomZ/go-develop-kit/path"
	"github.com/WindomZ/testify/assert"
)

var testProblem = &Problem{URL: "https://leetcode.com/problems/two-sum/description/"}

func TestProblem_Parse(t *testing.T) {
	assert.NoError(t, testProblem.Parse())
}

func TestProblem_ReadMe(t *testing.T) {
	assert.NotEmpty(t, testProblem.ReadMe())
}

func TestProblem_OutputReadMe(t *testing.T) {
	assert.NoError(t, testProblem.OutputReadMe("."))
	path.RemoveFile("README.md", false)
}
