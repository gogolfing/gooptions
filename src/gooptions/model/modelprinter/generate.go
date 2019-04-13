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

func (g *Generator) Fpln(f string, args ...interface{}) {
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
		g.Fpln("")
		GenTypeOptionList(g, tol)
	}
}

func GenDocComments(g *Generator, tols []*model.TypeOptionList) {

}

func GenPackage(g *Generator, tols []*model.TypeOptionList) {
	if len(tols) > 0 {
		g.Fpln("package %s", tols[0].PackageName)
	}
}

func GenTypeOptionList(g *Generator, tol *model.TypeOptionList) {
	GenOptionType(g, tol)

	g.Fpln("")

	GenTypeOptionFieldFactories(g, tol)
}

func GenOptionType(g *Generator, tol *model.TypeOptionList) {
	g.Fpln("type Option func(*%s)", tol.TypeName)
}

func GenTypeOptionFieldFactories(g *Generator, tol *model.TypeOptionList) {
	// g.Fpln("func With%s(%s %s) Option {", tol.Fields[0].Name, tol.Fields[0].Name, tol.Fields[0].Name)
	// g.Indent()
	// g.Fpln("return func(%s %s) {", "v", tol.Fields[0].Type.TypeString())
	// g.Indent()

	for _, field := range tol.Fields {
		GenTypeOptionFieldFactory(g, field)
		g.Fpln("")
	}
}

func GenTypeOptionFieldFactory(g *Generator, field *model.Field) {
	g.Fpln("func With%s(%s %s) Option {", field.Name, ParamNameFromType(field.Type.TypeString()), field.Name)
}
