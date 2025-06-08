// https://leetcode.com/problems/valid-parentheses
package solution

// Time: O(n)
// Space: O(n)
func isValid(str string) bool {
	pairs := [128]rune{
		')':'(',
		']':'[',
		'}':'{',
	}
	stack := []rune{}
	for _, bracket := range str {
		if pairs[bracket] == 0 { // left bracket
			stack = append(stack, bracket)
		} else { // right bracket
			if len(stack) == 0 || stack[len(stack)-1] != pairs[bracket] {
				return false
			}
			stack = stack[:len(stack)-1] // stack pop
		}
	}
	return len(stack) == 0 // no remaining left brackets
}
