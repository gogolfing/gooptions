package modelprinter

import (
	"log"

	"github.com/gogolfing/gooptions/src/gooptions/model"
)

type Printer struct {
	config Config

	model *model.Model
}

func New(c Config, m *model.Model) *Printer {
	return &Printer{
		config: c,
		model:  m,
	}
}

//w may be used if there is a sinlge destination package and
func (p *Printer) Print() error {
	log.Println("Printing model ...")

	err := p.model.VisitSourceFilePathTypeOptionLists(
		func(sourceFilePath string, tol *model.TypeOptionList) error {
			//TODO do some joining of source to output files.
			//TODO do some printing.

			fp := NewFilePrinter(sourceFilePath+"_dest", []*model.TypeOptionList{tol})

			err := fp.Print()

			return err
		},
	)

	return err
}
