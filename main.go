package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/WindomZ/leetcode-init/leetcode"
)

const (
	usageKey      = "the keyword of leetcode problem, such as: '1', 'two-sum', 'Two Sum'"
	usageTitle    = "the title of leetcode problem, such as: 'two-sum', 'Two Sum'"
	usageURL      = "the url of leetcode problem."
	usageMarkdown = "load and rendering markdown template, and save to TEMPLATE.md."
	usageHelp     = "prints a usage message documenting all defined command-line flags."
	usageVersion  = "prints version information."
)

var (
	keyFlag      string
	urlFlag      string
	markdownFlag string
	helpFlag     bool
	versionFlag  bool
)

func main() {
	flag.StringVar(&keyFlag, "k", "", usageKey)
	flag.StringVar(&keyFlag, "t", "", usageTitle)
	flag.StringVar(&urlFlag, "u", "", usageURL)
	flag.StringVar(&markdownFlag, "m", "", usageMarkdown)
	flag.BoolVar(&helpFlag, "h", false, usageHelp)
	flag.BoolVar(&versionFlag, "v", false, usageVersion)

	flag.Parse()

	if helpFlag {
		flag.Usage()
		return
	}

	if versionFlag {
		fmt.Fprintf(os.Stderr, "%s version: v1.1.0\n", os.Args[0])
		return
	}

	var problem *leetcode.Problem
	if urlFlag != "" {
		problem = leetcode.NewProblemByURI(leetcode.LanguageGo,
			urlFlag, markdownFlag)
	} else if keyFlag != "" {
		problem = leetcode.NewProblem(leetcode.LanguageGo,
			keyFlag, markdownFlag)
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
