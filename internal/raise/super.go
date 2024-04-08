/*
Copyright © 2024 Ixo
*/
package raise

import (
	"fmt"
	"strings"
)

// Returns normal and superscript set of characters to be found.
//
// Depending on the math_flag and ext_flag the slice is extended with non-numerical superscript characters.
func mapper(math_flag bool, ext_flag bool) ([]string, []string) {
	normal := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	normal_math := []string{"+", "-", "=", "(", ")"}
	normal_ext := []string{"i", "n"}

	superscript := []string{"⁰", "¹", "²", "³", "⁴", "⁵", "⁶", "⁷", "⁸", "⁹"}
	superscript_math := []string{"⁺", "⁻", "⁼", "⁽", "⁾"}
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
func SuperLoop(original_string string, math_flag bool, ext_flag bool) {
	normal, superscript := mapper(math_flag, ext_flag)

	for i, char := range normal {

		original_string = strings.Replace(original_string, char, superscript[i], -1)
	}

	fmt.Println(original_string)
}
