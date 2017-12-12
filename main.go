package main

import (
	"flag"

	"github.com/WindomZ/leetcode-init/leetcode"
)

const (
	usageTitle    = "the title of leetcode problem, without the number."
	usageURL      = "the url of leetcode problem."
	usageMarkdown = "load and rendering markdown template, and save to TEMPLATE.md."
	usageHelp     = "prints a usage message documenting all defined command-line flags."
)

var (
	titleFlag    string
	urlFlag      string
	markdownFlag string
	helpFlag     bool
)

func main() {
	flag.StringVar(&titleFlag, "t", "", usageTitle)
	flag.StringVar(&urlFlag, "u", "", usageURL)
	flag.StringVar(&markdownFlag, "m", "", usageMarkdown)
	flag.BoolVar(&helpFlag, "h", false, usageHelp)

	flag.Parse()

	if helpFlag {
		flag.Usage()
		return
	}

	var problem *leetcode.Problem
	if urlFlag != "" {
		problem = leetcode.NewProblem(leetcode.LanguageGo,
			urlFlag, markdownFlag)
	} else if titleFlag != "" {
		problem = leetcode.NewProblemByTitle(leetcode.LanguageGo,
			titleFlag, markdownFlag)
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
	if err := problem.OutputCode(); err != nil {
		panic(err)
	}
	if err := problem.OutputTestCode(); err != nil {
		panic(err)
	}
	if err := problem.OutputMarkdown(); err != nil {
		panic(err)
	}
}
