/*
Copyright © 2024 Ixo
*/
package raise

import (
	"fmt"
	"strings"
)

/*
replaces chars with superscript versions
*/
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

/*
- Get string / strings from args
- Check if user wants all symbols or just numerical
	- if all symbols add extra chars to check slices
*/
