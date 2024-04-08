/*
Copyright © 2024 Ixo
*/
package raise

import (
	"fmt"
	"regexp"
	"strings"
)

// Returns normal and superscript set of characters to be found.
//
// Depending on the math_flag and ext_flag the slice is extended with non-numerical superscript characters.
func mapper(math_flag bool, ext_flag bool) ([]string, []string) {
	normal := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "-"}
	normal_math := []string{"+", "=", "(", ")"}
	normal_ext := []string{"i", "n"}

	superscript := []string{"⁰", "¹", "²", "³", "⁴", "⁵", "⁶", "⁷", "⁸", "⁹", "⁻"}
	superscript_math := []string{"⁺", "⁼", "⁽", "⁾"}
	superscript_ext := []string{"ⁱ", "ⁿ"}

	if math_flag {
		normal = append(normal, normal_math...)
		superscript = append(superscript, superscript_math...)
	}

	if ext_flag {
		normal = append(normal, normal_ext...)
		superscript = append(superscript, superscript_ext...)
	}
	return normal, superscript
}

// Replaces chars in a string with the superscript versions if they exist.
//
// It loops through the normal characters in a slice and replaces them in the string one at a time.
func superreplace(original_string string, normal_chars []string, super_chars []string) string {
	for i, char := range normal_chars {
		original_string = strings.Replace(original_string, char, super_chars[i], -1)
	}
	return original_string
}

// uses superreplace function raw to replace any possible characters with superscript version
func SuperSimple(original_string string, math_flag bool, ext_flag bool) {
	normal, superscript := mapper(math_flag, ext_flag)
	fmt.Println(superreplace(original_string, normal, superscript))
}

// Replace matching characters after a "^" with their superscript equivalent
//
// TODO add a flag to make it so each number needs a "^" rather than catching all matching ones
func SuperUpReg(original_string string, math_flag bool, ext_flag bool) {
	normal, superscript := mapper(math_flag, ext_flag)

	var search_string string

	// careful with linking here or changing flags as the mapper func can cause some issues
	if math_flag && ext_flag {
		search_string = `\^\^(-*[\d+-=()in]+)`
	} else if math_flag {
		search_string = `\^(-*[\d+-=()]+)`
	} else if ext_flag {
		search_string = `\^(-*[\din]+)`
	} else {
		search_string = `\^(-*\d+)`
	}

	var search = regexp.MustCompile(search_string)
	matches := search.FindAllStringSubmatch(original_string, -1)

	for _, match := range matches {
		updated := superreplace(match[1], normal, superscript)

		original_string = strings.Replace(original_string, match[0], updated, 1)
	}

	fmt.Println(original_string)
}
