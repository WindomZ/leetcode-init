package leetcode

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

var testProblem = NewProblem(LanguageGo,
	"https://leetcode.com/problems/two-sum/description/")

func TestProblem_Parse(t *testing.T) {
	assert.NoError(t, testProblem.Parse())
}

func TestProblem_ReadMe(t *testing.T) {
	assert.NotEmpty(t, testProblem.ReadMe())
}

func TestProblem_OutputReadMe(t *testing.T) {
	assert.NoError(t, testProblem.OutputReadMe())
}

func TestProblem_OutputCode(t *testing.T) {
	assert.NoError(t, testProblem.OutputCode())
}

func TestProblem_OutputTestCode(t *testing.T) {
	assert.NoError(t, testProblem.OutputTestCode())
}
