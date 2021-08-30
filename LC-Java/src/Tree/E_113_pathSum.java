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
