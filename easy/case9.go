package easy

import (
	"fmt"
	"strconv"
)

func RunCase9() {
	fmt.Println("Case9")
	x := 102301

	res := isPalindrome(x)
	fmt.Println(res)
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	stringifiedInt := strconv.Itoa(x)
	totalLen := len(stringifiedInt)
	limit := 0

	if len(stringifiedInt)%2 == 0 {
		limit = totalLen / 2
	} else {
		limit = (totalLen - 1) / 2
	}

	for i := 0; i < limit; i++ {
		if stringifiedInt[i] != stringifiedInt[totalLen-1-i] {
			return false
		}
	}

	return true
}
