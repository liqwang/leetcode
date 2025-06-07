// https://leetcode.com/problems/longest-substring-without-repeating-characters/description
import java.util.HashMap;

public class Q3 {

    // Time: O(n)
    // Space: O(∣Σ∣), the size of the character set
    public int lengthOfLongestSubstring(String str) {
        boolean[] chars = new boolean[128]; // the hashset of current window
        int left = 0;
        int longest = 0;

        // why this works: each iteration finds the longest string ending at `right` pointer
        for(int right = 0; right < str.length(); right++) {
            char c = str.charAt(right);
            if(chars[c]){ // if current window contains the right character
                while(chars[c]){ // move the left pointer until the window clears the duplicate character
                    chars[str.charAt(left++)] = false;
                }
            }else{ // in this case, the length of current window may be the longest
                longest = Math.max(longest, right-left+1);
            }
            chars[c] = true; // put the right character into the hashset
        }
        return longest;
    }
}
