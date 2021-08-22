package DFS;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashSet;
import java.util.List;

//给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，在字符串中增加空格来构建一个句子，使得句子中所有的单词都在词典中。返回所有这些可能的句子。
//
// 说明：
//
// 分隔时可以重复使用字典中的单词。
// 你可以假设字典中没有重复的单词。
public class _H_140_wordBreak {
    public static void main(String[] args) {
        String s = "catsanddog";
        List<String> wordDict = Arrays.asList("cat", "cats", "and", "sand", "dog");
        System.out.println(new _H_140_wordBreak().wordBreak(s, wordDict));
    }

    public List<String> wordBreak(String s, List<String> wordDict) {
        if (s == null || wordDict == null) {
            return null;
        }
        HashSet<String> dicts = new HashSet<>(wordDict);
        ArrayList<String> res = new ArrayList<>();
        dfs(res, new StringBuilder(), s, 0, dicts, new boolean[s.length()]);
        return res;
    }

    private void dfs(ArrayList<String> res, StringBuilder sb, String s, int startIndex, HashSet<String> dicts, boolean[] used) {
        // boolean firstBlank = false;
        //     Arrays.fill(used, false);
        if (startIndex == s.length()) {

            if (sb.charAt(0) == ' ') {
                sb.deleteCharAt(0);
            }
            res.add(sb.toString());
            return;
        } else if (startIndex > s.length()) {
            return;
        } else {
            for (int i = startIndex; i < s.length(); i++) {
                String substring = s.substring(startIndex, i + 1);
                if (!dicts.contains(substring) || used[i]) {
                    continue;
                }
                if (!sb.isEmpty()) {
                    sb.append(" " + substring);
                used[i] = true;
                } else {
                    sb.append(substring);
                used[i] = true;
                }

                dfs(res, sb, s, i + 1, dicts, used);
                used[i] = false;
                // startIndex--;

                res.remove(res.size()-1);

                // sb.delete(0, i);
            }
        }
    }
}

