package modelprinter

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

var (
	//Normally - is the range character in the character class. That is why it is at the end of the class.
	arrayPointerChanPrefix = regexp.MustCompile(`^(\[[^\]]*\])?([<*-]+)?`)
)

//ParamNameFromType returns a parameter name from a type's name.
//It returns the lower-cased first rune of typeName.
//As a special case, if typeName is empty, then this rune value is set to '_'.
//If this value is equal to typeName, then an underscore is appended to the result.
//
//Any slice or array type expression at the beginning of typeName, along with any
//of the following < - * runes will be removed.
//This is to account for slice and array types as well as channel types in any direction.
//
//If there is a period in typeName, then the value after the period is then treated
//as typeName. This is to account for package references like os.File.
//
//Note that an invalid type name will likely result in an invalid parameter name.
func ParamNameFromType(typeName string) string {
	typeName = TrimArrayPointerChanPrefix(typeName)

	//Remove the package name, SelectorType, prefix if present.
	if index := strings.Index(typeName, "."); index >= 0 {
		typeName = typeName[index+1:]
	}

	firstRune, _ := utf8.DecodeRune([]byte(typeName))
	firstRuneString := string(firstRune)

	if firstRune == utf8.RuneError {
		firstRuneString = ""
	}

	result := strings.ToLower(string(firstRuneString))
	if result == typeName || result == "_" {
		result += "_"
	}

	return result
}

//TrimArrayPointerChanPrefix returns typeName with a possible array or slice type
//square bracket length expression followed by all * < - runes at the beginning
//of typeString replaced with the empty string..
func TrimArrayPointerChanPrefix(typeName string) string {
	return arrayPointerChanPrefix.ReplaceAllString(typeName, "")
}
