package main

import (
	"os"

	"github.com/vishen/go-chromecast/cmd"
)

var (
	// These are build-time variables that get set by goreleaser.
	version = "dev"
	commit  = "master"
	date    = ""
)

func main() {
	os.Exit(main1())
}

func main1() int {
	return cmd.Execute(version, commit, date)
}