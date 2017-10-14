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
	usageTitle string = "the title of leetcode problem, without the number."
	usageURL          = "the url of leetcode problem."
	usageHelp         = "prints a usage message documenting all defined command-line flags."
)

var (
	titleFlag = flag.String("title", defaultTitle, usageTitle)
	urlFlag   = flag.String("url", defaultURL, usageURL)
	helpFlag  = flag.Bool("help", false, usageHelp)
)

func main() {
	flag.StringVar(titleFlag, "t", defaultTitle, usageTitle)
	flag.StringVar(urlFlag, "u", defaultURL, usageURL)
	flag.BoolVar(helpFlag, "h", false, usageHelp)

	flag.Parse()

	if *helpFlag {
		flag.Usage()
		return
	}

	var problem *leetcode.Problem
	if *urlFlag != "" {
		problem = leetcode.NewProblem(leetcode.LanguageGo, *urlFlag)
	} else if *titleFlag != "" {
		problem = leetcode.NewProblemByTitle(leetcode.LanguageGo, *titleFlag)
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
}
