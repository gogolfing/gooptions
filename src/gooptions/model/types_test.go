package model_test

import (
	"go/ast"
	"reflect"
	"testing"

	. "github.com/gogolfing/gooptions/src/gooptions/model"
)

func TestIdentType_SetPackageNames_DoesNotSetAny(t *testing.T) {
	it := IdentType("value")

	var pns map[string]bool

	it.SetPackageNames(pns)

	if l := len(pns); l != 0 {
		t.Fatal(l)
	}
}

func TestIdentType_TypeString_ReturnsTheStringCastedValues(t *testing.T) {
	it := IdentType("value")

	if ts := it.TypeString(); ts != "value" {
		t.Fatal(ts)
	}
}

func TestChanType_SetPackageNames_DefersToCompositeType(t *testing.T) {
	wantNames := map[string]bool{
		"a": true,
		"b": false,
	}

	ct := &ChanType{
		Type: &StubTargetType{
			PackageNames: wantNames,
		},
	}

	result := map[string]bool{}
	ct.SetPackageNames(result)

	if !reflect.DeepEqual(result, wantNames) {
		t.Fatal()
	}
}

func TestChanType_TypeString_ReturnsCorrectTypeStringAndDirections(t *testing.T) {
	cases := []struct {
		compType string
		dir      ast.ChanDir
		result   string
	}{
		{"int", ast.SEND | ast.RECV, "chan int"},
		{"rune", ast.SEND, "chan<- rune"},
		{"bool", ast.RECV, "<-chan bool"},
		{"float64", 0, "chan float64"}, //Shouldn't happend, but still testing againt it.
	}

	for i, tc := range cases {
		ct := &ChanType{
			ChanDir: tc.dir,
			Type:    IdentType(tc.compType),
		}

		result := ct.TypeString()

		if result != tc.result {
			t.Errorf("%d: %q WANT %q", i, result, tc.result)
		}
	}
}

func TestPointerType_SetPackageNames_DefersToCompositeType(t *testing.T) {
	wantNames := map[string]bool{
		"a": true,
		"b": false,
	}

	pt := &PointerType{
		Type: &StubTargetType{
			PackageNames: wantNames,
		},
	}

	result := map[string]bool{}
	pt.SetPackageNames(result)

	if !reflect.DeepEqual(result, wantNames) {
		t.Fatal()
	}
}

func TestPointerType_TypeString_ReturnsStarPlusCompositeType(t *testing.T) {
	pt := &PointerType{
		Type: &StubTargetType{
			TypeString_: "foobar",
		},
	}

	if result := pt.TypeString(); result != "*foobar" {
		t.Fatal(result)
	}
}

func TestArrayType_SetPackageNames_DefersToCompositeType(t *testing.T) {
	wantNames := map[string]bool{
		"a": true,
		"b": false,
	}

	at := &ArrayType{
		Type: &StubTargetType{
			PackageNames: wantNames,
		},
	}

	result := map[string]bool{}
	at.SetPackageNames(result)

	if !reflect.DeepEqual(result, wantNames) {
		t.Fatal()
	}
}

func TestArrayType_TypeString_ReturnsSliceWithLengthPlusCompositeType(t *testing.T) {
	at := &ArrayType{
		Len: "len",
		Type: &StubTargetType{
			TypeString_: "foobar",
		},
	}

	if result := at.TypeString(); result != "[len]foobar" {
		t.Fatal(result)
	}
}

type StubTargetType struct {
	PackageNames map[string]bool

	TypeString_ string
}

func (stt *StubTargetType) SetPackageNames(pns map[string]bool) {
	for k, v := range stt.PackageNames {
		pns[k] = v
	}
}

func (stt *StubTargetType) TypeString() string {
	return stt.TypeString_
}
