package leetcode

import "errors"

// ICode the Code interface
type ICode interface {
	outputCode(dirName, packageName string) error
	outputTestCode(dirName, packageName string) error
}

// Code get a Code by key string
func (q Question) Code(lang LanguageType) (ICode, error) {
	if c := q.Codes.Code(lang.String()); c != nil {
		switch lang {
		case LanguageGo:
			return &CodeGo{Code: *c}, nil
		}
	}
	return nil, errors.New("not support the language: " + lang.String())
}
