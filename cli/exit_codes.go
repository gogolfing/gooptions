package cli

import "flag"

const (
	ExitCodeFlagParseErrHelp      = 2
	ExitCodeFlagParseNonHelpError = 3

	ExitCodeDefaultError = 127
)

type FlagParseError struct {
	Err error
}

func (e *FlagParseError) Error() string {
	if e.Err != flag.ErrHelp {
		return e.Err.Error()
	}
	return ""
}

func ExitCodeFromError(err error) int {
	if err == nil {
		return 0
	}

	if e, ok := err.(*FlagParseError); ok {
		if e.Err == flag.ErrHelp {
			return ExitCodeFlagParseErrHelp
		}
		return ExitCodeFlagParseNonHelpError
	}

	return ExitCodeDefaultError
}
