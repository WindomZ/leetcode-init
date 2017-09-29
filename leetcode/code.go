package leetcode

import (
	"errors"
	"path/filepath"
	"strings"

	"github.com/WindomZ/go-develop-kit/path"
)

// Code the struct of leetcode codes.
type Code struct {
	Value       string `json:"value"`
	Text        string `json:"text"`
	DefaultCode string `json:"defaultCode"`
}

type Codes []*Code

func (c Codes) Code(key string) *Code {
	for _, code := range c {
		if code.match(key) {
			return code
		}
	}
	return nil
}

func (c Code) match(key string) bool {
	key = strings.ToLower(key)
	return strings.ToLower(c.Value) == key ||
		strings.ToLower(c.Text) == key
}

func (c Code) outputCode(packageName, lang string) error {
	var fileName string
	var head string
	switch strings.ToLower(lang) {
	case "golang", "go":
		fileName = packageName + ".go"
		head = "package " + packageName
	default:
		return errors.New("not support the language")
	}
	return path.OverwriteFile(
		filepath.Join(".", packageName, fileName),
		head, "", c.DefaultCode,
	)
}

func (c Code) outputTestCode(packageName, lang string) error {
	var fileName string
	var head string
	switch strings.ToLower(lang) {
	case "golang", "go":
		fileName = packageName + "_test.go"
		head = "package " + packageName
	default:
		return errors.New("not support the language")
	}
	return path.OverwriteFile(
		filepath.Join(".", packageName, fileName),
		head, "",
	)
}