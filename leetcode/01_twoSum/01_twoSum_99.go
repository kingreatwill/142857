package main

// 优秀时间.
func twoSum(nums []int, target int) []int {
	result := []int{}
	m := make(map[int]int)
	for i, v := range nums {
		if value, exist := m[target-v]; exist {
			result = append(result, value)
			result = append(result, i)
		}
		m[v] = i
	}
	return result
}

// 优秀内存.
func twoSum(nums []int, target int) []int {
	step := 1
	for i, v1 := range nums[0 : len(nums)-1] {
		for k, v2 := range nums[step:] {
			if v1+v2 == target {
				return []int{i, step + k}
			}
		}
		step++
	}

	return []int{}
}
