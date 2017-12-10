package leetcode

import (
	"path/filepath"

	"github.com/WindomZ/go-develop-kit/path"
)

// CodeGo the struct of leetcode codes.
type CodeGo struct {
	Code
}

func (c CodeGo) outputCode(dirName, packageName string) error {
	return path.OverwriteFile(
		filepath.Join(".", dirName, packageName+".go"),
		"package "+packageName, "", c.DefaultCode,
	)
}

func (c CodeGo) outputTestCode(dirName, packageName string) error {
	return path.OverwriteFile(
		filepath.Join(".", dirName, packageName+"_test.go"),
		"package "+packageName, "", `import "testing"`, "",
		"func Test_"+mustFindFirstStringSubmatch(`func (.+)\(`,
			c.DefaultCode)+"(t *testing.T) {", "}", "",
		"func Benchmark_"+mustFindFirstStringSubmatch(`func (.+)\(`,
			c.DefaultCode)+"(b *testing.B) {", "}",
	)
}
