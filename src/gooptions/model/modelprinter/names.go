package modelprinter

import (
	"strings"
	"unicode/utf8"
)

//ParamNameFromType returns a parameter name from a type's name.
//It returns the lower-cased first rune of typeName.
//As a special case, if typeName is empty, then this rune value is set to '_'.
//If this value is equal to typeName, then an underscore is appended to the result.
//
//Note that an invalid type name will likely result in an invalid parameter name.
//
//TODO something about a qualifier package name.
func ParamNameFromType(typeName string) string {
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
