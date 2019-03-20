package modelprinter

import (
	"bytes"
	"fmt"

	"github.com/gogolfing/gooptions/src/gooptions/model"
)

type Generator struct {
	buffer *bytes.Buffer

	indent string
}

func NewGenerator() *Generator {
	return &Generator{
		buffer: &bytes.Buffer{},
		indent: "",
	}
}

func (g *Generator) Indent() {
	g.indent += "\t"
}

func (g *Generator) Unindent() {
	if l := len(g.indent); l > 0 {
		g.indent = g.indent[0 : l-1]
	}
}

func (g *Generator) P(f string, args ...interface{}) {
	fmt.Fprintf(g.buffer, g.indent+f+"\n", args...)
}

func (g *Generator) Bytes() []byte {
	return g.buffer.Bytes()
}

//All tols must be same package. Only first is used in output.
func GenTypeOptionLists(g *Generator /*config*/, tols []*model.TypeOptionList) {
	GenDocComments(g, tols)

	GenPackage(g, tols)

	for _, tol := range tols {
		g.P("")
		GenTypeOptionList(g, tol)
	}
}

func GenDocComments(g *Generator, tols []*model.TypeOptionList) {

}

func GenPackage(g *Generator, tols []*model.TypeOptionList) {
	if len(tols) > 0 {
		g.P("package %s", tols[0].PackageName)
	}
}

func GenTypeOptionList(g *Generator, tol *model.TypeOptionList) {
	GenOptionType(g, tol)

	GenTypeOptionFactories(g, tol)
}

func GenOptionType(g *Generator, tol *model.TypeOptionList) {
	g.P("type Option func(*%s)", tol.TypeName)
}

func GenTypeOptionFactories(g *Generator, tol *model.TypeOptionList) {
	g.P("func With%s(%s %s) Option {", tol.Fields[0].Name, tol.Fields[0].Name, tol.Fields[0].Name)
	g.Indent()
	g.P("return func(%s %s) {", "v", tol.Fields[0].Type.TypeString())
	g.Indent()
}
