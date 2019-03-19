package gooptions

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"os"
	"path/filepath"
	"reflect"

	"github.com/gogolfing/gooptions/src/gooptions/model"
)

type InvalidPathError string

func (e InvalidPathError) Error() string {
	return fmt.Sprintf("gooptions: invalid path: %s", string(e))
}

var (
	ErrEmptyModel = errors.New("gooptions: generated model is empty")
)

func GenerateOptionsModel(out io.Writer, path string) (*model.Model, error) {
	goFilePaths, err := ListGoFilePaths(path)
	if err != nil {
		return nil, err
	}

	fmt.Fprintln(out, "goFilePaths", goFilePaths)

	goFileASTs, err := ParseGoFileASTs(goFilePaths)
	if err != nil {
		return nil, err
	}

	fmt.Fprintln(out, "goFileASTs", goFileASTs)

	result := model.NewModel()

	if err := PopulateModelFromGoFileASTs(result, goFileASTs); err != nil {
		return nil, err
	}

	return result, nil
}

func PopulateModelFromGoFileASTs(m *model.Model, goFileASTs map[string]*ast.File) error {
	//Loop over all files.
	for filePath, fileAST := range goFileASTs {

		//Loop over all top level declarations in the file.
		for _, decl := range fileAST.Decls {

			//Continue if the decl is a general declaration. See go/ast package.
			if genDecl, ok := decl.(*ast.GenDecl); ok {

				//Continue if the genDecl is a type declaration.
				if genDecl.Tok == token.TYPE {

					//Loop over all type specs in a possibly paranthesis'd type declaration.
					for _, spec := range genDecl.Specs {

						//Continue if spec is a type spec. This should always be the case from above.
						if typeSpec, ok := spec.(*ast.TypeSpec); ok {

							//Continue if the typeSpec is non anonymous.
							if typeSpec.Name != nil {

								//Continue if typeSpec is a struct type.
								if structType, ok := typeSpec.Type.(*ast.StructType); ok {

									//TODO filter from config.
									//Add the found and filtered structType to m.
									addStructTypeToModel(m, filePath, fileAST.Name.Name, typeSpec.Name.Name, structType)
								}
							}
						}
					}
				}
			}
		}
	}

	if m.IsEmpty() {
		return ErrEmptyModel
	}

	return nil
}

func addStructTypeToModel(m *model.Model, filePath, packageName, typeName string, structType *ast.StructType) error {
	fmt.Println("addStructTypeToModel()", typeName, len(structType.Fields.List))
	for _, field := range structType.Fields.List {
		fmt.Println(field.Names, field.Type, reflect.TypeOf(field.Type))

		if ident, ok := field.Type.(*ast.Ident); ok {
			fmt.Println(ident.Name, ident.Obj)
		}
	}

	tol := model.NewTypeOptionList(packageName, typeName, structType.Fields.List)

	return m.AddType(filePath, tol)
}

func CollectModelFieldsFromASTFieldList(fieldList *ast.FieldList) []*model.Field {
	for _, field := range fieldList.List {
		fmt.Println(field)
	}
	return nil
}

func NewModelFieldFromASTField(field *ast.Field) *model.Field {
	switch fieldType := field.Type.(type) {
	case *ast.Ident:
		fmt.Println(fieldType.Name)
	}
	return nil
}

func ListGoFilePaths(path string) ([]string, error) {
	pathFileInfo, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	if !pathFileInfo.IsDir() {
		if !isFileInfoAGoFile(pathFileInfo) {
			return nil, InvalidPathError("path is not a directory nor a Go file")
		}
		return []string{filepath.Clean(path)}, nil
	}

	file, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
	if err != nil {
		return nil, err
	}

	fileInfos, err := file.Readdir(-1)
	if err != nil {
		return nil, err
	}

	result := make([]string, 0, len(fileInfos))
	for _, fileInfo := range fileInfos {
		if isFileInfoAGoFile(fileInfo) {
			result = append(result, filepath.Join(path, fileInfo.Name()))
		}
	}

	if len(result) == 0 {
		return nil, InvalidPathError("directory does not contain any Go files")
	}

	return result, nil
}

func ParseGoFileASTs(filePaths []string) (map[string]*ast.File, error) {
	fileSet := token.NewFileSet()

	result := make(map[string]*ast.File, len(filePaths))
	for _, filePath := range filePaths {
		file, err := parser.ParseFile(fileSet, filePath, nil, 0)
		if err != nil {
			return nil, err
		}

		result[filePath] = file
	}

	return result, nil
}

func isFileInfoAGoFile(fi os.FileInfo) bool {
	ok, err := filepath.Match("*.go", fi.Name())

	return ok && err == nil && fi.Mode().IsRegular()
}
