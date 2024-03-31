// Given an integer array nums, return true if any value appears at least twice in the array,
// and return false if every element is distinct.

// Example 1:

// Input: nums = [1,2,3,1]
// Output: true
// Example 2:

// Input: nums = [1,2,3,4]
// Output: false

package Con_Dupl

func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool, len(nums))

	for _, num := range nums {
		if seen[num] {
			return true // If the number is already seen, return true
		}
		seen[num] = true
	}

	return false
}
