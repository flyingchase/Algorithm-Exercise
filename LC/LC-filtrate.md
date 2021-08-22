# LC 重点





LC 881







## 101 对称二叉树

### 题意

​	给定一个二叉树，检查它是否是镜像对称的。



### 思路：

- 递归 要求左右两侧子树均对称 

  - `return isSymmetric(leftRoot.left,rightRoot.right)&& isSymmetric(leftRoot.right,rightRoot.left)`

  - 左侧子树的右子树与右侧子树的左子树比较

  - ```java
    public boolean isSymmetric (TreeNode root) {
        if (root==null) {
            return true;
        }
        TreeNode cur = root;
        return isSymmetric(cur.left,cur.right);
    }
    
    private boolean isSymmetric(TreeNode lRoot, TreeNode rRoot) {
        if (lRoot==null&&rRoot==null) {
            return true;
        }
        if (lRoot==null||rRoot==null) {
            return false;
        }
        // 上述两个 if 可以替换成
            if(lRoot==null || rRoot==null)
            return lRoot==rRoot;
        if (lRoot.val!=rRoot.val) {
            return false;
        }
    
        return isSymmetric(lRoot.left,rRoot.right)&&isSymmetric(rRoot.left,lRoot.right);
    }
    ```

- 迭代 使用 queue 保证先进先出

  - 将 root 的 left 和 right 依次入队  并 poll 出来比较
  - 再将 left.right right.left 和left.left right.right 依次入堆  循环弹出
  - 注意：`left==right==null`时 contine ！queue.isEmpty()的循环

  

- BFS







## 113 二叉树路径和

### 题意

给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。

叶子节点 是指没有子节点的节点。





### 思路：

- stack 实现 DFS
  - 保存上个节点 prev 做到回溯
- 回溯 recursive



### codes

```java
package Tree;

import DataStructure.TreeNode;

import java.sql.Array;
import java.util.ArrayList;
import java.util.List;
import java.util.Stack;

/*给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。

叶子节点 是指没有子节点的节点。*/
public class E_113_pathSum {

    // use satck to implement dfs
    //
    // public List<List<Integer>> pathSum(TreeNode root, int targetSum) {
    //     if (root==null) {
    //         return null;
    //     }
    //     TreeNode cur = root;
    //     TreeNode prev = null;
    //     Stack<TreeNode> stack = new Stack<>();
    //     List<Integer> path = new ArrayList<>();
    //     List<List<Integer>> res = new ArrayList<>();
    //     int sum=0;
    //     while (cur != null || !stack.isEmpty()) {
    //         while (cur != null) {
    //             stack.push(cur);
    //             path.add(cur.val);
    //             sum+=cur.val;
    //             cur=cur.left;
    //         }
    //         cur = stack.peek();
    //         if (cur.right!=null&&cur.right!=prev) {
    //             cur=cur.right;
    //             continue;
    //         }
    //         if (cur.left==null&&cur.right==null&&sum==targetSum) {
    //             res.add(new ArrayList<>(path));
    //         }
    //
    //         prev=cur;
    //         stack.pop();
    //         path.remove(path.size()-1);
    //         sum-=cur.val;
    //         cur=null;
    //     }
    //
    //     return res;
    // }

    public List<List<Integer>> pathSum(TreeNode root, int targetSum) {
        if (root==null) {
            return null;
        }
        List<Integer> path = new ArrayList<>();
        List<List<Integer>> res = new ArrayList<>();
        int sum=0;
        TreeNode cur = root;
        dfs(cur,sum,res,path);
        return res;

    }

    private void dfs(TreeNode root, int sum, List<List<Integer>> res, List<Integer> path) {
        if (root==null) {
            return;
        }
        path.add(root.val);
        if (root.left==null&&root.right==null) {
            if (root.val==sum) {
                res.add(new ArrayList<>(path));
            }
            return;
        }

        if (root.left==null) {
            dfs(root.left,sum-root.val,res,path);
            path.remove(path.size()-1);
        }
        if (root.right==null) {
            dfs(root.right,sum-root.val,res,path);
        }
    }
}
```

## 114 二叉树转换为链表

### 题意

将一棵二叉树按照前序遍历拆解成为一个 `假链表`。所谓的假链表是说，用二叉树的 *right* 指针，来表示链表中的 *next* 指针。

### codes

``` java
public void flatten(TreeNode root) {
    if (root==null) {
        return;
    }
    while (root != null) {
        if (root.left!=null) {
            TreeNode right = root.right;
            root.right=root.left;
            root.left=null;
            TreeNode node = root.right;
            while (node != null && node.right != null) {
                node=node.right;
            }
            node.right=right;

        }
        root=root.right;
    }
}
```





### 注意事项

- 非递归：
  - 不断将左子树拆下来接到右边去。然后将左子树与右子树连接。







## 127 单词接龙-BFS

### 题意

字典 wordList 中从单词 beginWord 和 endWord 的 转换序列 是一个按下述规格形成的序列：

​	序列中第一个单词是 beginWord 。
​	序列中最后一个单词是 endWord 。
​	每次转换只能改变一个字母。
​	转换过程中的中间单词必须是字典 wordList 中的单词。
给你两个单词 beginWord 和 endWord 和一个字典 wordList ，找到从 beginWord 到 endWord 的 最短转换序列 中的 单词数目 。如果不存在这样的转换序列，返回 0。

### 思路：

- 传统 BFS

  - 构造 set 字典存储 wordList queue 作为遍历路径

  - 双层遍历 确保 str 每个位置被替换 26 次

    - 替换中 set.contains 存在则 set.remove queue.offer

  - ```java
    public int ladderLength(String beginWord, String endWord, List<String> wordList) {
        Set<String> set = new HashSet<>(wordList);
        Queue<String> queue = new LinkedList<>();
        queue.offer(beginWord);
        int step = 1, len = beginWord.length();
        while (!queue.isEmpty()) {
            int size = queue.size();
            while (size-- > 0) {
                String cur = queue.poll();
                if (cur.equals(endWord)) {
                    return step;
                }
                for (int i = 0; i < len; i++) {
                    // 单词的每个位置均替换
                    for (char ch = 'a'; ch <= 'z'; ch++) {
                        StringBuilder next = new StringBuilder(cur);
                        next.setCharAt(i, ch);
                        String nextWord = next.toString();
                        if (set.contains(nextWord)) {
                            if (nextWord.equals(endWord)) {
                                return step + 1;
                            }
                            // or deadwhile
                            set.remove(nextWord);
                            queue.offer(nextWord);
                        }
                    }
                }
            }
            step++;
        }
        return 0;
    }
    ```



- 双向 BFS
  - 同时从起点终点开始扩展 直到交集出现
    - <img src="https://cdn.jsdelivr.net/gh/flyingchase/Private-Img@master/uPic/image-20210821162748064.png" alt="image-20210821162748064" style="zoom:50%;" />











## 490 505迷宫-二维 BFS

``` java
int[][] dirs = {{0,1},{0,-1},{-1,0},{1,0}};
    public int shortestDistance(int[][] maze, int[] start, int[] destination) {
        int[][] distance = new int[maze.length][maze[0].length];
        for (int[] row : distance) {
            Arrays.fill(row,Integer.MAX_VALUE);
        }
        distance[start[0]][start[1]]=0;
        dijkstra(maze,start,distance);
        return distance[destination[0]][destination[1]] == Integer.MAX_VALUE ? -1 : distance[destination[0]][destination[1]];
    }

    private void dijkstra(int[][] maze, int[] start, int[][] distance) {
        PriorityQueue<int[]> q = new PriorityQueue<>((a, b) -> (a[2] - b[2]));
        q.offer(new int[]{start[0],start[1]});
        while (!q.isEmpty()) {
            int[] cur = q.poll();
            for (int[] dir : dirs) {
                int x = cur[0] + dir[0];
                int y = cur[1] + dir[1];
                int count =0;
                while (x >= 0 && y >= 0 && x < maze.length && y < maze[0].length && maze[x][y] == 0) {
                    x+=dir[0];
                    y+=dir[1];
                    count++;
                }
                x-=dir[0];
                y-=dir[1];
                if (distance[cur[0]][cur[1]]+count< distance[x][y]) {
                    distance[x][y]=distance[cur[0]][cur[1]]+count;
                    q.add(new int[]{x,y,distance[x][y]});
                }
            }
        }
    }
```











