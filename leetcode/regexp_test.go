package leetcode

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func Test_mustFindFirstStringSubmatch(t *testing.T) {
	assert.Equal(t, "",
		mustFindFirstStringSubmatch(`abc(def)ghi`, `abcefghij`))
	assert.Equal(t, "def",
		mustFindFirstStringSubmatch(`abc(def)ghi`, `abcdefghi`))
}
