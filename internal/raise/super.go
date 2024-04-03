/*
Copyright © 2024 Ixo
*/
package raise

import (
	"fmt"
	"strings"
)

// SuperLoop replaces chars in a string with the superscript versions.
//
// It loops through the normal characters in a slice and replaces them in the string one at a time.
// Depending on the bools given to math_flag and ext_flag the slice is extended with non-numerical superscript characters.
func SuperLoop(original_string string, math_flag bool, ext_flag bool) {
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

	for i, char := range normal {

		original_string = strings.Replace(original_string, char, superscript[i], -1)
	}

	fmt.Println(original_string)
}
