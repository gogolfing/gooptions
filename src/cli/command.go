package cli

import (
	"flag"
	"fmt"
	"io"

	"github.com/gogolfing/gooptions/src/gooptions"
	"github.com/gogolfing/gooptions/src/gooptions/model/modelprinter"
)

type Command struct {
	path string
}

func (c *Command) Do(_ string, args []string, out, outErr io.Writer) int {
	err := c.do(args, out, outErr)
	defer maybePrintError(outErr, err)

	return ExitCodeFromError(err)
}

func (c *Command) do(args []string, out, outErr io.Writer) error {
	if err := c.setAndParseFlags(args, outErr); err != nil {
		return err
	}

	optionsModel, err := gooptions.GenerateOptionsModel(out, c.path)
	if err != nil {
		return err
	}

	printerConfig := modelprinter.Config{}
	printer := modelprinter.New(printerConfig, optionsModel)

	err = printer.Print()

	return err
}

func (c *Command) setAndParseFlags(args []string, outErr io.Writer) error {
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.SetOutput(outErr)

	fs.StringVar(&c.path, "path", ".", "`path` to directory or file that contains the type for which to create options")

	err := fs.Parse(args)
	if err != nil {
		err = &FlagParseError{Err: err}
	}
	return err
}

func maybePrintError(w io.Writer, err error) {
	if err != nil {
		fmt.Fprintln(w, err)
	}
}
