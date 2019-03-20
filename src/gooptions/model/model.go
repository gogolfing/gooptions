package model

import "errors"

var (
	ErrEmptyTypeName = errors.New("model: empty type name")
)

type Model struct {
	//Slice of ordered filePath string indexed by order added to the Model.
	orderedSourceFilePaths []string

	//Map from filePath strings to all the TypeOptionLists found in the file.
	options map[string][]*TypeOptionList
}

func NewModel() *Model {
	return &Model{
		orderedSourceFilePaths: []string{},
		options:                map[string][]*TypeOptionList{},
	}
}

func (m *Model) IsEmpty() bool {
	return len(m.orderedSourceFilePaths) == 0 || len(m.options) == 0
}

func (m *Model) AddType(filePath string, tol *TypeOptionList) error {
	if tol.TypeName == "" {
		return ErrEmptyTypeName
	}

	tolSlice, ok := m.options[filePath]
	if !ok {
		m.orderedSourceFilePaths = append(m.orderedSourceFilePaths, filePath)
	}
	m.options[filePath] = append(tolSlice, tol)

	return nil
}

func (m *Model) VisitSourceFilePathTypeOptionLists(visit func(sourceFilePath string, tol *TypeOptionList) error) error {
	for _, sourceFilePath := range m.orderedSourceFilePaths {
		for _, tol := range m.options[sourceFilePath] {
			if err := visit(sourceFilePath, tol); err != nil {
				return err
			}
		}
	}

	return nil
}
