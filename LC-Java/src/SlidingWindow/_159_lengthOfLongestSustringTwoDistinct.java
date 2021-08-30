package SlidingWindow;

import java.util.HashMap;

/*
* 给定一个字符串，找出最长子串
T
T的长度，它最多包含2个不同的字符。*/
public class _159_lengthOfLongestSustringTwoDistinct {
    public static void main(String[] args) {
        String s = "eceba";
        System.out.println(new _159_lengthOfLongestSustringTwoDistinct().lengthofLongestSubstringTwoDistinct(s, 2));

    }


    public int lengthofLongestSubstringTwoDistinct(String s, int k) {
        HashMap<Character, Integer> map = new HashMap<>();
        int left = 0, res = 0;
        for (int i = 0; i < s.length(); i++) {
            char cur = s.charAt(i);
            // 遍历值入窗口
            map.put(cur, map.getOrDefault(cur, 0) + 1);
            // 窗口不符合条件 left 持续退出
            while (map.size() > k) {
                char c = s.charAt(left);
                map.put(c, map.get(c) - 1);
                if (map.get(c) == 0) {
                    map.remove(c);
                }
                // size超过窗口大小即left 左移
                left++;
            }
            // 计算子串最值
            res = Math.max(res, i - left + 1);
        }
        return res;
    }
}
