// https://leetcode.com/problems/add-two-numbers
package solution

import . "github.com/liqwang/leetcode/lib"

// Time: O(max(m,n))
// Space: O(1), not includes the necessary cost of the returned linked list
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dummyHead := &ListNode{Val: 0}
    cur := dummyHead
    carry := 0
    for l1 != nil || l2 != nil || carry != 0 {
        x := 0
        if l1 != nil {
            x = l1.Val
        }
        y := 0
        if l2 != nil {
            y = l2.Val
        }
        sum := x + y + carry
        carry = sum / 10
        cur.Next = &ListNode{Val: sum % 10}

        cur = cur.Next
        if l1 != nil {
            l1 = l1.Next
        }
        if l2 != nil {
            l2 = l2.Next
        }
    }
    return dummyHead.Next
}
