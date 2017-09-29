package leetcode

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

var testQuestion = NewQuestion("https://leetcode.com/problems/two-sum/description/")

func TestQuestion_Parse(t *testing.T) {
	assert.NoError(t, testQuestion.Parse())
	assert.NotEmpty(t, testQuestion.String())
}
