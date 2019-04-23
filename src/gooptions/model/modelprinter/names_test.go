package modelprinter_test

import (
	"fmt"
	"testing"

	. "github.com/gogolfing/gooptions/src/gooptions/model/modelprinter"
)

func ExampleParamName() {
	printTypeAndParamName := func(typeName string) {
		fmt.Printf("%q -> %q\n", typeName, ParamNameFromType(typeName))
	}

	printTypeAndParamName("int")
	printTypeAndParamName("<-chan bool")
	printTypeAndParamName("chan<- *os.File")
	printTypeAndParamName("[2]byte")
	printTypeAndParamName("[]string")
	printTypeAndParamName("map[string]int")

	//Output:
	//"int" -> "i"
	//"<-chan bool" -> "c"
	//"chan<- *os.File" -> "f"
	//"[2]byte" -> "b"
	//"[]string" -> "s"
	//"map[string]int" -> "m"
}

func TestParamNameFromType(t *testing.T) {
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

		result = ParamNameFromType("packageName." + tc.typeName)
		if result != tc.result {
			t.Errorf("%d: packageName.result = %q WANT %q", i, result, tc.result)
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

		{"[]", ""},
		{"[2]_", "_"},
		{"[0xA]F", "F"},
		{"[...]****", ""},
		{"[02314]<-*<-<---<<", ""},
		{"[]*<- ", " "},
		{"[]*int", "int"},
	}

	for i, tc := range cases {
		result := TrimArrayPointerChanPrefix(tc.typeName)

		if result != tc.result {
			t.Errorf("%d: result = %q WANT %q", i, result, tc.result)
		}
	}
}
