package model

import "fmt"

type Field struct {
	Name string
	Type FieldType
}

func (f *Field) GoString() string {
	return fmt.Sprintf("&model.Field{%q %q}", f.Name, f.Type.TypeString())
}

type FieldType interface {
	Packages() PackageSet

	TypeString() string
}

type PackageSet map[string]bool

func (ps PackageSet) Add(name string) {
	ps[name] = true
}

func (ps PackageSet) With(other PackageSet) PackageSet {
	result := make(map[string]bool, len(ps)+len(other))

	for key, value := range ps {
		if value {
			result[key] = true
		}
	}
	for key, value := range other {
		if value {
			result[key] = true
		}
	}

	return result
}

type IdentType string

func (t IdentType) Packages() PackageSet {
	return nil
}

func (t IdentType) TypeString() string {
	return string(t)
}
