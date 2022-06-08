package easy

import (
	"fmt"
	"strings"
)

func RunCase20() {
	fmt.Println("Case20")
	s := "{}[][][](){[]()}"
	valid := isValidUsingStack(s)
	fmt.Println(valid)
}

// first try
func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	keys := []string{"{}", "[]", "()"}
	maxReps := len(s) / 2
	ss := s
	for len(ss) != 0 && maxReps != 0 {
		for _, v := range keys {
			if strings.Contains(ss, v) {
				ss = strings.ReplaceAll(ss, v, "")
			}
		}
		maxReps--
	}

	return len(ss) == 0
}

// second try
func isValidUsingStack(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	stack := []rune{}
	var sup rune
	pair := map[string]string{
		"{": "}",
		"(": ")",
		"[": "]",
	}

	revpair := map[string]string{
		"}": "{",
		"]": "[",
		")": "(",
	}

	for _, v := range s {
		_, ok := pair[string(v)]
		if ok {
			stack = append(stack, v)
		} else {
			if len(stack) == 0 {
				return false
			}

			sup, stack = stack[len(stack)-1], stack[:len(stack)-1]
			_, ok := revpair[string(v)]
			if !ok || string(sup) != revpair[string(v)] {
				return false
			}
		}
	}

	return len(stack) == 0

}
