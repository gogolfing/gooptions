package modelprinter

import (
	"fmt"

	"github.com/gogolfing/gooptions/src/gooptions/model"
)

type FilePrinter struct {
	//some write closer factory.

	//TODO some sort of individaul tol config.

	tols []*model.TypeOptionList
}

func NewFilePrinter(destFilePath string, tols []*model.TypeOptionList) *FilePrinter {
	return &FilePrinter{
		tols: tols,
	}
}

func (fp *FilePrinter) Print() error {
	g := NewGenerator()

	GenTypeOptionLists(g, fp.tols)

	fmt.Println(string(g.Bytes()))

	return nil
}
