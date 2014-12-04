package main

import (
	"fmt"
	"os"

	flags "github.com/thbishop/aws-creds/Godeps/_workspace/src/github.com/jessevdk/go-flags"
)

type options struct {
	Export  bool   `short:"e" long:"export" description:"Output as export statements"`
	Profile string `short:"p" long:"profile" description:"Name of profile in ~/.aws/credentials" required:"true" default:"default"`
	Version func() `short:"v" long:"version" description:"Display the version of aws-creds"`
}

func parseCliArgs() *options {
	opts := &options{}

	opts.Version = func() {
		fmt.Println(version)
		os.Exit(0)
	}

	parser := flags.NewParser(opts, flags.Default)

	args, err := parser.Parse()
	if err != nil {
		helpDisplayed := false

		for _, i := range args {
			if i == "-h" || i == "--help" {
				helpDisplayed = true
				break
			}
		}

		if !helpDisplayed {
			parser.WriteHelp(os.Stderr)
		}
		os.Exit(1)
	}

	return opts
}
