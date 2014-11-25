package main

import (
	"fmt"
	"os"

	flags "github.com/jessevdk/go-flags"
)

type options struct {
	Profile string `short:"p" long:"profile" description:"Name of profile in ~/.aws/config" required:"true"`
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
