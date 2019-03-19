package main

import (
	"os"

	"github.com/gogolfing/gooptions/src/cli"
)

func main() {
	c := &cli.Command{}

	exitCode := c.Do(os.Args[0], os.Args[1:], os.Stdout, os.Stderr)

	os.Exit(exitCode)
}
