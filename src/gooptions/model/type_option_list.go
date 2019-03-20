package model

type TypeOptionList struct {
	SourceFilePath string
	PackageName    string
	TypeName       string
	Fields         []*Field
}

func NewTypeOptionList(sourceFilePath, packageName, typeName string, fields []*Field) *TypeOptionList {
	return &TypeOptionList{
		SourceFilePath: sourceFilePath,
		PackageName:    packageName,
		TypeName:       typeName,
		Fields:         fields,
	}
}
