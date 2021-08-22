package DFS;

import java.util.ArrayList;
import java.util.Comparator;
import java.util.List;

//给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。
//
// 给你一个整数 n ，返回和为 n 的完全平方数的 最少数量 。
public class M_279_numSquares {
    public static void main(String[] args) {
        // System.out.println(new M_279_numSquares().numSquares(13));
    }
    // dfs 超时
    // public int numSquares(int n) {
    //     ArrayList<List<Integer>> res = new ArrayList<>();
    //     // Integer min = Integer.MAX_VALUE;
    //     int min =dfs(res, new ArrayList<Integer>(), n, n, 1);
    //
    //     return min;
    // }
    //
    // private int dfs(ArrayList<List<Integer>> res, ArrayList<Integer> lists, int n, int min, int start) {
    //     if (n < 0) {
    //         return -1;
    //     } else if (n == 0) {
    //         res.add(new ArrayList<>(lists));
    //         res.sort(new Comparator<List<Integer>>() {
    //             @Override
    //             public int compare(List<Integer> o1, List<Integer> o2) {
    //                 if (o1.size() == o2.size()) {
    //                     return 0;
    //                 } else {
    //
    //                     return o1.size() - o2.size();
    //                 }
    //
    //             }
    //         });
    //         min = Math.min(res.get(0).size(), min);
    //         // return min;
    //     } else {
    //         for (int i = start; i <= n; i++) {
    //             lists.add(i * i);
    //             // min++;
    //             dfs(res, lists, n - i * i, min, i);
    //             min = Math.min(res.get(0).size(), min);
    //             lists.remove(lists.size() - 1);
    //         }
    //     }
    //     return min;
    // }

    // bfs test
    // public int numSquares(int n) {
    //
    // }
}
