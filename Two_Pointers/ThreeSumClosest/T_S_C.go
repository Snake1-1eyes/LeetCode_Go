package threesumclosest

import "sort"

func threeSumCosest(nums []int, target int) int {
	closest := nums[0] + nums[1] + nums[2]
	if len(nums) <= 3 {
		return closest
	}

	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1
		for left < right {
			currentSum := nums[i] + nums[left] + nums[right]
			if abs(currentSum-target) < abs(closest-target) {
				closest = currentSum
			}
			if currentSum == target {
				return currentSum
			} else if currentSum < target {
				left++
			} else {
				right--
			}
		}
	}
	return closest
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
