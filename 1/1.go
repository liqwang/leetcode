package solution

func twoSum(nums []int, target int) []int {
	indexes := make(map[int]int)
	for i, num := range nums {
		need := target - num
		if index, found := indexes[need]; found {
			return []int{index, i}
		}
		indexes[num] = i
	}
	return nil; // never reach
}
