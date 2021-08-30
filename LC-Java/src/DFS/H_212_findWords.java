// package DFS;
//
// import java.util.ArrayList;
// import java.util.HashSet;
// import java.util.List;
// import java.util.PriorityQueue;
//
// //给定一个 m x n 二维字符网格 board 和一个单词（字符串）列表 words，找出所有同时在二维网格和字典中出现的单词。
// //
// // 单词必须按照字母顺序，通过 相邻的单元格 内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母在一个单词中不允许被重复使用。
//
// public class H_212_findWords {
//     int[][] dirs= {{0,1},{0,-1},{-1,0},{1,0}};
//     public List<String> findWords(char[][] board, String[] words) {
//         HashSet<String> set = new HashSet<String>(List.of(words));
//         ArrayList<String> res = new ArrayList<>();
//         boolean[][] used = new boolean[board.length][board[0].length];
//
//         dfs(res,new StringBuilder(),board,set,used);
//         return res;
//     }
//
//     private void dfs(ArrayList<String> res, StringBuilder sb, char[][] board, HashSet<String> set, boolean[][] used) {
//         PriorityQueue<int[]> q = new PriorityQueue<>((a, b) -> (a[2] - b[2]));
//         q.offer(new int[]{})
//     }
// }
