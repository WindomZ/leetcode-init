package leetcode

import (
	"errors"
	"strings"
)

// ICode the Code interface
type ICode interface {
	outputCode(dirName, packageName string) error
	outputTestCode(dirName, packageName string) error
}

// Code the struct of leetcode codes.
type Code struct {
	Value       string `json:"value"`
	Text        string `json:"text"`
	DefaultCode string `json:"defaultCode"`
}

func (c Code) match(key string) bool {
	key = strings.ToLower(key)
	return strings.ToLower(c.Value) == key ||
		strings.ToLower(c.Text) == key
}

// Codes the slice of Code
type Codes []*Code

// Code get a Code by key string
func (c Codes) Code(lang LanguageType) (ICode, error) {
	for _, code := range c {
		if code.match(lang.String()) {
			switch lang {
			case LanguageGo:
				return &CodeGo{Code: *code}, nil
			}
		}
	}
	return nil, errors.New("not support the language: " + lang.String())
}
