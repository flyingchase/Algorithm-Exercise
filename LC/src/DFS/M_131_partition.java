package DFS;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

/*给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。

回文串 是正着读和反着读都一样的字符串。

*/
public class M_131_partition {

    public static void main(String[] args) {
        String s = "aab";
        System.out.println(new M_131_partition().partition(s));
    }

    public List<List<String>> partition(String s) {
        if (s == null) {
            return null;
        } else if (s.length() == 1) {
            new ArrayList<>(Arrays.asList(s));
        }
        List<List<String>> res = new ArrayList<>();
        dfs(res, new ArrayList<String>(), s, 0);

        return res;
    }

    private void dfs(List<List<String>> res, List<String> lists, String s, int start) {
        if (start == s.length()) {
            res.add(new ArrayList<>(lists));
            return;
        } else {
            for (int i = start; i < s.length(); i++) {
                String substring = s.substring(start, i +1);
                if (!isPalindrome(substring)) {
                    continue;
                }
                lists.add(substring);
                dfs(res,lists,s,i+1);
                lists.remove(lists.size() - 1);
            }
        }
    }

    private boolean isPalindrome(String s) {

        int len = s.length();
        for (int i = 0, j = len - 1; i <= j; i++, j--) {
            if (s.charAt(i) != s.charAt(j)) {
                return false;
            }

        }
        return true;
    }
}
