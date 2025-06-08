// https://leetcode.com/problems/longest-substring-without-repeating-characters
package solution

// Time: O(n)
// Space: O(∣Σ∣), the size of the character set
func lengthOfLongestSubstring(str string) int {
    chars := [128]bool{}
	left := 0
	longest := 0
	
	// why this works: each iteration finds the longest string ending at `right` pointer
	for right, char := range str {
		if chars[char] {
			for chars[char] {
				chars[str[left]] = false
				left++
			}
		} else {
			longest = max(longest, right-left+1)
		}
		chars[char] = true
	}
	return longest
}
