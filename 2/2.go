// https://leetcode.com/problems/add-two-numbers
package solution

import . "github.com/liqwang/leetcode/lib"

// Time: O(max(m,n))
// Space: O(1), not includes the necessary cost of the returned linked list
func addTwoNumbers(l1, l2 *ListNode) *ListNode {
    dummy := &ListNode{}
    cur := dummy
    carry := 0
    for l1 != nil || l2 != nil || carry != 0 {
        sum := carry
        if l1 != nil {
            sum += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            sum += l2.Val
            l2 = l2.Next
        }
        carry = sum / 10
        cur.Next = &ListNode{Val: sum % 10}
        cur = cur.Next
    }
    return dummy.Next
}
