package model

import "go/ast"

type TypeOptionList struct {
	PackageName string
	TypeName    string
	Fields      []*ast.Field
}

func NewTypeOptionList(packageName, typeName string, fields []*ast.Field) *TypeOptionList {
	return &TypeOptionList{
		PackageName: packageName,
		TypeName:    typeName,
		Fields:      fields,
	}
}

type Field struct{}
