package DFS;

import java.util.*;

//126. 单词接龙 II
public class _H_126_findLadders {
    public List<List<String>> findLadders(String beginWord, String endWord, List<String> wordList) {
        HashSet<String> set = new HashSet<>(wordList);
        Queue<String> queue = new LinkedList<>();
        ArrayList<List<String>> res = new ArrayList<>();
        queue.add(beginWord);
        int step = 1,len=beginWord.length();
        while (!queue.isEmpty()) {
            int size = queue.size();
            List<String> lists = new ArrayList<>();
            while (size-- > 0) {
                String cur = queue.poll();
                if (cur.equals(endWord)) {
                    return res;
                }
                for (int i = 0; i < len; i++) {
                    for (char ch = 'a';ch<='z';ch++) {
                        StringBuilder sb = new StringBuilder(cur);
                        sb.setCharAt(i,ch);
                        String newWord = sb.toString();
                        if (set.contains(newWord)) {
                            lists.add(newWord);
                            step++;
                            res.add(new ArrayList<>(lists));
                            return res;
                        }
                        set.remove(newWord);
                        queue.add(newWord);
                    }
                }
            }
            step++;
        }
        return res;
    }
}
