// https://leetcode.com/problems/valid-parentheses

import java.util.ArrayDeque;

public class Q20 {

    // Time: O(n)
    // Space: O(n)
    public boolean isValid(String str) {
        char[] pairs = new char[128];
        pairs[')'] = '(';
        pairs[']'] = '[';
        pairs['}'] = '{';

        ArrayDeque<Character> stack = new ArrayDeque<>();
        for(int i=0; i<str.length(); i++) {
            char bracket = str.charAt(i);
            if (pairs[bracket] == 0) { // left bracket
                stack.push(bracket);
            } else if (stack.isEmpty() || pairs[bracket] != stack.pop()) {
                return false;
            }
        }
        return stack.isEmpty(); // no remaining left brackets
    }
}
