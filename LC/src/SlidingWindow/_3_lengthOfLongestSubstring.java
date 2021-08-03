package SlidingWindow;

import java.util.HashMap;

/*给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

 */
public class _3_lengthOfLongestSubstring {
    public static void main(String[] args) {
        String s = "pwwkew";
        System.out.println(new _3_lengthOfLongestSubstring().lengthOfLongestSubstring(s));
    }

    public int lengthOfLongestSubstring(String s) {
        int left = 0, right = 0, res = 0;
        HashMap<Character, Integer> map = new HashMap<>();
        while (right < s.length()) {
            while (right < s.length() && map.getOrDefault(s.charAt(right), 0) < 1) {
                map.put(s.charAt(right), map.getOrDefault(s.charAt(right), 0) + 1);
                right++;
            }
            res = Math.max(res, right - left);
            while (right < s.length() && map.getOrDefault(s.charAt(right), 0) >= 1) {
                char c = s.charAt(left);
                map.put(c, map.getOrDefault(c, 0) - 1);
                left++;
            }
        }
        return res;
    }
}
