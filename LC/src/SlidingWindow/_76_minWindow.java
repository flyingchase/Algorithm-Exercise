package SlidingWindow;

import java.util.HashMap;

/*给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。

 */
public class _76_minWindow {
    public static void main(String[] args) {
        String s = "ADOBECODEBANC", t = "ABC";
        System.out.println(new _76_minWindow().minWindow(s, t));
    }
    public String minWindow(String s, String t) {
        if (s==null) {
            return null;
        }
        HashMap<Character, Integer> map = new HashMap<>();
        int left=0,minStart=0,minLen=Integer.MAX_VALUE,count=0;
        for (char c : t.toCharArray()) {
            map.put(c,map.getOrDefault(c,0)+1);
        }
        for (int i = 0; i < s.length(); i++) {
            char c = s.charAt(i);
            if (map.containsKey(c)) {
                if (map.get(c)>0) {
                    count++;
                }
                map.put(c,map.get(c) -1);
            }
            while (t.length() == count) {
                if (minLen>i-left+1) {
                    minLen=i-left + 1;
                    minStart=left;
                }
                char leftChar = s.charAt(left);
                if (map.containsKey(leftChar)) {
                    map.put(leftChar,map.get(leftChar)+1);
                    if (map.get(leftChar)>0) {
                        count--;
                    }
                }
                left++;
            }
        }
        if (minLen==Integer.MAX_VALUE) {
            return "";
        }
        return s.substring(minStart, minStart + minLen);
    }
}
