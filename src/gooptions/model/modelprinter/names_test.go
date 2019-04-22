package modelprinter_test

import (
	"testing"

	. "github.com/gogolfing/gooptions/src/gooptions/model/modelprinter"
)

func TestParamName(t *testing.T) {
	cases := []struct {
		typeName string
		result   string
	}{
		{"", "_"},
		{"_", "__"},
		{"__", "__"},
		{"___", "__"},
		{" ", " _"},
		{"_x9", "__"},
		{"\n", "\n_"},
		{"$", "$_"},
		{"f", "f_"},
		{"F", "f"},
		{"1", "1_"},
		{"1234", "1"},
		{"αβ", "α"},
		{"Ëllo", "ë"},
		{"foo", "f"},
		{"Foo", "f"},
		{"ThisVariableIsExported", "t"},
	}

	for i, tc := range cases {
		result := ParamNameFromType(tc.typeName)
		if result != tc.result {
			t.Errorf("%d: result = %q WANT %q", i, result, tc.result)
		}

		result = ParamNameFromType("*<-" + tc.typeName)
		if result != tc.result {
			t.Errorf("%d: result*<- = %q WANT %q", i, result, tc.result)
		}
	}
}

func TestTrimStarAndChanRunesPrefix(t *testing.T) {
	cases := []struct {
		typeName string
		result   string
	}{
		{"", ""},
		{"_", "_"},
		{"F", "F"},
		{"****", ""},
		{"<-*<-<---<<", ""},
		{"*<- ", " "},
		{"*int", "int"},
	}

	for i, tc := range cases {
		result := TrimStarAndChanRunesPrefix(tc.typeName)

		if result != tc.result {
			t.Errorf("%d: result = %q WANT %q", i, result, tc.result)
		}
	}
}
