package BFS;

import java.util.HashSet;
import java.util.List;
import java.util.Set;

// 字典 wordList 中从单词 beginWord 和 endWord 的 转换序列 是一个按下述规格形成的序列：
//
//         序列中第一个单词是 beginWord 。
//         序列中最后一个单词是 endWord 。
//         每次转换只能改变一个字母。
//         转换过程中的中间单词必须是字典 wordList 中的单词。
//         给你两个单词 beginWord 和 endWord 和一个字典 wordList ，找到从 beginWord 到 endWord 的 最短转换序列 中的 单词数目 。如果不存在这样的转换序列，返回 0。
public class H_127_wordList {
    // traditional bfs
    // public int ladderLength(String beginWord, String endWord, List<String> wordList) {
    //     Set<String> set = new HashSet<>(wordList);
    //     Queue<String> queue = new LinkedList<>();
    //     queue.offer(beginWord);
    //
    //     int step = 1, len = beginWord.length();
    //     while (!queue.isEmpty()) {
    //         int size = queue.size();
    //         while (size-- > 0) {
    //             String cur = queue.poll();
    //             if (cur.equals(endWord)) {
    //                 return step;
    //             }
    //             for (int i = 0; i < len; i++) {
    //                 // 单词的每个位置均替换
    //                 for (char ch = 'a'; ch < ='z'; ch++) {
    //                     StringBuilder next = new StringBuilder(cur);
    //                     next.setCharAt(i, ch);
    //                     String nextWord = next.toString();
    //                     if (set.contains(nextWord)) {
    //                         if (nextWord.equals(endWord)) {
    //                             return step + 1;
    //                         }
    //                         // or deadwhile
    //                         set.remove(nextWord);
    //                         queue.offer(nextWord);
    //                     }
    //                 }
    //             }
    //         }
    //         step++;
    //     }
    //     return 0;
    // }

    // two-end bfs
    public int ladderLength(String beginWord, String endWord, List<String> wordList) {
        Set<String> beginSet = new HashSet<>();
        Set<String> endSet = new HashSet<>();
        HashSet<String> vosited = new HashSet<>();

        Set<String> dict = new HashSet<>(wordList);

        if (!dict.contains(endWord)) {
            return 0;
        }
        int step = 1, len = beginWord.length();
        beginSet.add(beginWord);
        endSet.add(endWord);
        while (!beginSet.isEmpty() && !endSet.isEmpty()) {
            Set<String> nextSet = new HashSet<>();
            for (String word : beginSet) {
                char[] chs = word.toCharArray();
                for (int i = 0; i < len; i++) {
                    for (char c = 'a'; c <= 'z'; c++) {
                        char prev = chs[i];
                        chs[i] = c;
                        String nextWord = new String(chs);
                        if (endSet.contains(nextWord)) {
                            return step + 1;
                        }
                        if (vosited.add(nextWord) && dict.contains(nextWord)) {
                            nextSet.add(nextWord);
                        }
                        chs[i] = prev;
                    }
                }
            }
            // 判断endset 和 nextset 谁小 走谁 遍历的是 beginset 故需要替换
            if (endSet.size() < nextSet.size()) {
                beginSet = endSet;
                endSet = nextSet;
            } else {
                beginSet = nextSet;
            }
            step++;
        }
        return 0;
    }

}
