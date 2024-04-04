// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// You can return the answer in any order.

// Example 1:

// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
// Example 2:

// Input: nums = [3,2,4], target = 6
// Output: [1,2]
// Example 3:

// Input: nums = [3,3], target = 6
// Output: [0,1]

// Constraints:

// 2 <= nums.length <= 104
// -109 <= nums[i] <= 109
// -109 <= target <= 109
// Only one valid answer exists.

package Two_Sum

func twoSum(nums []int, target int) []int {
	var s []int
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if (nums[i] + nums[j]) == target {
				s = append(s, i, j)
				return s
			}
		}
	}
	return s
}

// func twoSum(nums []int, target int) []int {
// 	hashMap := make(map[int]int)
// 	for i, num := range nums {
// 	  difference := target - num
// 	  if index, ok := hashMap[difference]; ok {
// 		return []int{index, i}
// 	  }
// 	  hashMap[num] = i
// 	}
// 	return []int{}
//   }
