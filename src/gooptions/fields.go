package gooptions

import (
	"fmt"
	"go/ast"
	"log"

	"github.com/gogolfing/gooptions/src/gooptions/model"
)

func addStructTypeToModel(m *model.Model, filePath, packageName, typeName string, structType *ast.StructType) error {
	fmt.Println("addStructTypeToModel()", typeName, len(structType.Fields.List))

	modelFields := CollectModelFieldsFromASTFieldList(structType.Fields)

	fmt.Printf("%#v\n", modelFields)

	tol := model.NewTypeOptionList(
		filePath,
		packageName,
		typeName,
		CollectModelFieldsFromASTFieldList(structType.Fields),
	)

	return m.AddType(filePath, tol)
}

func CollectModelFieldsFromASTFieldList(fieldList *ast.FieldList) []*model.Field {
	result := make([]*model.Field, 0, len(fieldList.List))

	for _, field := range fieldList.List {
		modelField := NewModelFieldFromASTField(field)
		if modelField != nil {
			result = append(result, modelField)
		}
	}

	return result
}

func NewModelFieldFromASTField(field *ast.Field) *model.Field {
	var modelFT model.FieldType

	switch astFT := field.Type.(type) {
	case *ast.Ident:
		modelFT = NewModelIdentType(astFT)

	default:
		log.Printf("unsupported *ast.Field.Type %T", astFT)
	}

	return &model.Field{
		Name: NameOfField(field),
		Type: modelFT,
	}
}

func NameOfField(field *ast.Field) string {
	if len(field.Names) == 0 {
		return ""
	}

	return field.Names[0].Name
}

func NewModelIdentType(ident *ast.Ident) model.IdentType {
	return model.IdentType(ident.Name)
}
