package main

import (
	"flag"

	"github.com/WindomZ/leetcode-init/leetcode"
)

const (
	defaultTitle string = ""
	defaultURL          = ""
)

const (
	usageTitle string = ""
	usageURL          = ""
)

var (
	titleFlag = flag.String("title", defaultTitle, usageTitle)
	urlFlag   = flag.String("url", defaultURL, usageURL)
)

func main() {
	flag.StringVar(titleFlag, "t", defaultTitle, usageTitle)
	flag.StringVar(urlFlag, "u", defaultURL, usageURL)

	flag.Parse()

	var problem *leetcode.Problem
	if *urlFlag != "" {
		problem = leetcode.NewProblem(*urlFlag)
	} else if *titleFlag != "" {
		problem = leetcode.NewProblemByTitle(*titleFlag)
	} else {
		flag.Usage()
		return
	}

	if err := problem.Parse(); err != nil {
		panic(err)
	}

	if err := problem.OutputReadMe(); err != nil {
		panic(err)
	}
	if err := problem.OutputCode("go"); err != nil {
		panic(err)
	}
	if err := problem.OutputTestCode("go"); err != nil {
		panic(err)
	}
}
