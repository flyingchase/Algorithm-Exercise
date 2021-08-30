package SlidingWindow;

import java.util.HashMap;

/*给定字符串S，找到最多有k个不同字符的最长子串T的长度。

 */
public class _340_lengthOfLongestSubstringDistinct {
    public static void main(String[] args) {

    }


    public int lengthOfLongestSubstringDistinct(String str, int k) {
        if (str==null||k>str.length()) {
            return Integer.parseInt(null);
        }
        HashMap<Character, Integer> map = new HashMap<>();
        int left =0,res=0;
        int len = str.length();
        for (int i = 0; i < len; i++) {
            char cur = str.charAt(i);
            map.put(cur,map.getOrDefault(cur,0)+1);
            while (map.size() > k) {
                char c = str.charAt(left);
                map.put(c,map.get(c)-1);
                if (map.get(c)==0) {
                    map.remove(c);
                }
                left++;

            }
            res=Math.max(res,i-left+1);
        }
        return res;
    }

}
