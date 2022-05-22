package easy

import (
	"fmt"
	"strings"
)

func RunCase13() {
	fmt.Println("Case13")

	s := "MCMXCIV"

	res := romanToInt(s)
	fmt.Println(res)
}

func romanToInt(s string) int {
	temp := s
	res := 0
	specialDict := map[string]int{
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}

	normalDict := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	for key, value := range specialDict {
		if strings.Contains(temp, key) {
			temp = strings.ReplaceAll(temp, key, "")
			res = res + value
		}
	}

	for _, ch := range temp {
		for k, v := range normalDict {
			if string(ch) == k {
				res = res + v
			}
		}
	}

	return res
}
