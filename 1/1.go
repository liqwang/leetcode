// https://leetcode.com/problems/two-sum
package solution

// Time: O(n)
// Space: O(n)
func twoSum(nums []int, target int) []int {
	indexes := make(map[int]int)
	for i, num := range nums {
		if index, found := indexes[num]; found {
			return []int{index, i}
		}
		indexes[target-num] = i
	}
	return nil; // never reach
}
