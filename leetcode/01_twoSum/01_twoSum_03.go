package main

import "fmt"

func main() {
	fmt.Printf("%v\n", twoSum([]int{2, 7, 11, 15}, 9))
}

func twoSum(nums []int, target int) []int {
	result := []int{}
	source := map[int]int{}
	for i, v := range nums {
		if value, exist := source[target-v]; exist {
			result = append(result, value)
			result = append(result, i)
		}
		source[v] = i
	}
	return result
}
