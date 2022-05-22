package easy

import (
	"fmt"
)

func RunCase1() {
	fmt.Println("Case1")

	nums := []int{3, 3}
	// nums := []int{3, 2, 4}
	// nums := []int{2, 7, 11, 15}
	target := 6
	// target := 6
	// target := 9
	res := twoSum(nums, target)

	fmt.Println(res)
}

func twoSum(nums []int, target int) []int {
	result := []int{}

	for index, value := range nums {
		if index+1 == len(nums) {
			continue
		}

		for subIndex, subValue := range nums[index+1:] {
			if value+subValue == target {
				result = append(result, index)
				result = append(result, index+1+subIndex)
			}
		}
	}

	return result
}
