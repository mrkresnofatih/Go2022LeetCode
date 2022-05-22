package easy

import (
	"fmt"
)

func RunCase14() {
	fmt.Println("Case14")

	// strs := []string{"flower", "flow", "flight"}
	strs := []string{"dog", "racecar", "car"}
	res := longestCommonPrefix(strs)

	fmt.Println(res)

}

func longestCommonPrefix(strs []string) string {
	// find lowest length
	minLen := 1000
	minIdx := 1000
	for id, v := range strs {
		if minLen > len(v) {
			minLen = len(v)
			minIdx = id
		}
	}
	if minLen == 0 {
		return ""
	}

	res := ""
	for i := 0; i < minLen; i++ {
		count := 0
		for _, v := range strs {
			if strs[minIdx][i] == v[i] {
				count++
			} else {
				return res
			}
		}

		if count == len(strs) {
			res = res + string(strs[minIdx][i])
		}
	}

	return res
}
