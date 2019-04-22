package gooptions

import (
	"errors"
	"go/ast"
	"log"

	"github.com/gogolfing/gooptions/src/gooptions/model"
)

var (
	errUnsupportedASTExpr = errors.New("unsupported ast.Expr type")
)

func addStructTypeToModel(m *model.Model, filePath, packageName, typeName string, structType *ast.StructType) error {
	log.Println("addStructTypeToModel()", typeName, len(structType.Fields.List))

	modelFields := CollectModelFieldsFromASTFieldList(structType.Fields)

	log.Printf("%#v\n", modelFields)

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
		if modelField, err := NewModelFieldFromASTField(field); err != nil {
			log.Println(err)
		} else {
			result = append(result, modelField)
		}
	}

	return result
}

func NewModelFieldFromASTField(field *ast.Field) (*model.Field, error) {
	modelTargetType, err := NewModelTargetType(field.Type)
	if err != nil {
		return nil, err
	}

	return &model.Field{
		Name: NameOfField(field),
		Type: modelTargetType,
	}, nil
}

func NewModelTargetType(expr ast.Expr) (model.TargetType, error) {
	var result model.TargetType
	var err error

	switch astType := expr.(type) {
	case *ast.ArrayType:
		result, err = NewModelArrayType(astType)

	case *ast.ChanType:
		result, err = NewModelChanType(astType)

	case *ast.Ident:
		result = NewModelIdentType(astType)

	case *ast.StarExpr:
		result, err = NewModelPointerType(astType)

	default:
		return nil, errUnsupportedASTExpr
	}

	return result, err
}

func NameOfField(field *ast.Field) string {
	if len(field.Names) == 0 {
		return ""
	}

	return field.Names[0].Name
}

func NewModelChanType(c *ast.ChanType) (*model.ChanType, error) {
	t, err := NewModelTargetType(c.Value)
	if err != nil {
		return nil, err
	}

	return &model.ChanType{
		ChanDir: c.Dir,
		Type:    t,
	}, nil
}

func NewModelIdentType(ident *ast.Ident) model.IdentType {
	return model.IdentType(ident.Name)
}

func NewModelPointerType(se *ast.StarExpr) (*model.PointerType, error) {
	t, err := NewModelTargetType(se.X)
	if err != nil {
		return nil, err
	}

	return &model.PointerType{
		Type: t,
	}, nil
}

func NewModelArrayType(at *ast.ArrayType) (*model.ArrayType, error) {
	t, err := NewModelTargetType(at.Elt)
	if err != nil {
		return nil, err
	}

	lenString := ""
	if at.Len != nil {
		if basic, ok := at.Len.(*ast.BasicLit); ok {
			lenString = basic.Value
		}
	}

	return &model.ArrayType{
		Len:  lenString,
		Type: t,
	}, nil
}
