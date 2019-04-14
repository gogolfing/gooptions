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

func (g *Generator) Outdent() {
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
	optionRecvType := tol.TypeName

	GenOptionType(g, optionRecvType)

	g.Fpln("")

	GenTypeOptionFieldFactories(g, optionRecvType, tol.Fields)
}

func GenOptionType(g *Generator, optionRecvType string) {
	g.Fpln("type Option func(*%s)", optionRecvType)
}

func GenTypeOptionFieldFactories(g *Generator, optionRecvType string, fields []*model.Field) {
	for _, field := range fields {
		GenTypeOptionFieldFactory(g, optionRecvType, field)
		g.Fpln("")
	}
}

func GenTypeOptionFieldFactory(g *Generator, optionRecvType string, field *model.Field) {
	fieldTypeString := field.Type.TypeString()
	paramName := ParamNameFromType(fieldTypeString)

	g.Fpln("func With%s(%s %s) Option {", field.Name, paramName, fieldTypeString)
	g.Indent()
	g.Fpln("return Option(func(%s *%s) {", "v", optionRecvType)
	g.Indent()
	g.Fpln("%s.%s = %s", "v", field.Name, paramName)
	g.Outdent()
	g.Fpln("})")
	g.Outdent()
	g.Fpln("}")
}
