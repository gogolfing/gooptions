package model

import (
	"fmt"
	"go/ast"
)

type Field struct {
	Name string
	Type TargetType
}

func (f *Field) GoString() string {
	return fmt.Sprintf("&model.Field{%q %q}", f.Name, f.Type.TypeString())
}

type TargetType interface {
	SetPackageNames(map[string]bool)

	TypeString() string
}

var (
	//Ensure that all required types implement TargetType.

	_ TargetType = IdentType("")
	_ *ChanType  = &ChanType{}
)

type IdentType string

func (t IdentType) SetPackageNames(_ map[string]bool) {}

func (t IdentType) TypeString() string {
	return string(t)
}

type ChanType struct {
	ChanDir ast.ChanDir
	Type    TargetType
}

const (
	ChanDirBoth = ast.SEND | ast.RECV
)

func (t *ChanType) SetPackageNames(pns map[string]bool) {
	t.Type.SetPackageNames(pns)
}

func (t *ChanType) TypeString() string {
	result := "chan"

	if t.ChanDir&ChanDirBoth != ChanDirBoth {
		if t.ChanDir&ast.SEND == ast.SEND {
			result = "<-" + result
		} else if t.ChanDir&ast.RECV == ast.RECV {
			result += "<-"
		}
	}

	result += " " + t.Type.TypeString()

	return result
}
