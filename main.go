package main

import "os"

func main() {
	err := mainWithArgs(os.Args[0], os.Args[1:])

	exitCode := 0
	if err != nil {
		exitCode = 1
	}

	os.Exit(exitCode)
}

func mainWithArgs(command string, args []string) error {
	return nil
}
