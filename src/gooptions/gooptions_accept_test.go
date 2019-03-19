package gooptions_test

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"

	. "github.com/gogolfing/gooptions/src/gooptions"
)

func TestListGoFilePaths_ReturnsPathErrorOnNonExistentFile(t *testing.T) {
	_, err := ListGoFilePaths("./does_not_exist")

	if _, ok := err.(*os.PathError); !ok {
		t.Fatal(err)
	}
}

func TestListGoFilePaths_ReturnsInvalidPathErrorWhenNormalFileIsNotAGoFile(t *testing.T) {
	_, err := ListGoFilePaths("./testdata/single.notgo")

	if _, ok := err.(InvalidPathError); !ok {
		t.Fatal(err)
	}
	if !strings.Contains(err.Error(), "nor a Go file") {
		t.Fatal(err)
	}
}

func TestListGoFilePaths_ReturnsInvalidPathErrorWhenDirectoryHasNoGoFiles(t *testing.T) {
	_, err := ListGoFilePaths("./testdata/dir_without_go_files")

	if _, ok := err.(InvalidPathError); !ok {
		t.Fatal(err)
	}
	if !strings.Contains(err.Error(), "does not contain any Go files") {
		t.Fatal(err)
	}
}

func TestListGoFilePaths_ReturnsCorrectPathsWhenNormalFileIsAGoFile(t *testing.T) {
	paths, err := ListGoFilePaths("./testdata/single.go")

	if err != nil {
		t.Fatal(err)
	}

	want := []string{"testdata/single.go"}

	if !reflect.DeepEqual(paths, want) {
		t.Fatalf("%v WANT %v", paths, want)
	}
}

func TestListGoFilePaths_ReturnsCorrectPathsForADirectoryWithGoFiles(t *testing.T) {
	paths, err := ListGoFilePaths("./testdata/dir")

	if err != nil {
		t.Fatal(err)
	}

	want := []string{
		"testdata/dir/a.go",
		"testdata/dir/a_test.go",
	}

	sort.Strings(paths)

	if !reflect.DeepEqual(paths, want) {
		t.Fatalf("%v WANT %v", paths, want)
	}
}

func TestParseGoFileASTs_ReturnsParseErrorWhenFileIsBadSyntax(t *testing.T) {
	paths := []string{"testdata/bad_syntax/imports.go"}

	_, err := ParseGoFileASTs(paths)

	if err == nil {
		t.Fatal(err)
	}
}

func TestParseGoFileASTs_ReturnsPopulatedMapWithFilePathsKeyedToNonNilFiles(t *testing.T) {
	paths := []string{
		"testdata/good_syntax/a.go",
		"testdata/good_syntax/b.go",
	}

	fileASTs, err := ParseGoFileASTs(paths)

	if err != nil {
		t.Fatal(err)
	}

	for _, path := range paths {
		fileAST, ok := fileASTs[path]
		if !ok || fileAST == nil {
			t.Errorf("want non-nil ast for path %q", path)
			continue
		}

		base := strings.TrimSuffix(filepath.Base(path), ".go")

		if want := fmt.Sprintf("goodsyntax_%s", base); fileAST.Name.Name != want {
			t.Errorf("package name = %q WANT %q", fileAST.Name.Name, want)
			continue
		}
	}
}
