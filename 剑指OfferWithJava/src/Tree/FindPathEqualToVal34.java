package Tree;

import DataStructure.TreeNode;

import java.util.ArrayList;
import java.util.Iterator;
import java.util.Stack;

public class FindPathEqualToVal34 {
    private static ArrayList<ArrayList<Integer>> ret = new ArrayList<>();
    public static ArrayList<ArrayList<Integer>> FindPath02(TreeNode root, int target) {
        backtracking(root, target,new ArrayList<>());
        return ret;
    }

    private static void backtracking(TreeNode node, int target, ArrayList<Integer> path) {
        if (node==null) {
            return;
        }
        path.add(node.val);
        target-= node.val;
        if (target==0&&node.left==null&&node.right==null) {
            ret.add(new ArrayList<>(path));

        }else {
            backtracking(node.left,target,path);
            backtracking(node.right,target,path);
        }
        path.remove(path.size()-1);
    }

    public static ArrayList<ArrayList<Stack<Integer>>> FindPath01 (TreeNode root, int target) {
        ArrayList<ArrayList<Stack<Integer>>> res = new ArrayList();
        if (root==null) {
            return res;
        }
        Stack<Integer> stack = new Stack<>();
        preOrder(root, stack,0,target,res);
        return res;
    }

    public static  void preOrder (TreeNode root, Stack<Integer> stack, int sum, int target, ArrayList<ArrayList<Stack<Integer>>> res) {
        if (root==null) {
            return;
        }
        stack.push(root.val);
        sum+= root.val;

        if (root.left==null&&root.right==null&&sum==target) {
            ArrayList<Stack<Integer>> one = new ArrayList<>();
            one.add(stack);
            res.add(one);
        }

        preOrder(root.left,stack,sum,target,res);
        preOrder(root.right,stack,sum,target,res);
        sum-=stack.pop();
    }
    public static void main(String[] args) {
        TreeNode root = new TreeNode(1);
        root.left=new TreeNode(2);
        root.right=new TreeNode(3);
        root.left.left=new TreeNode(4);
        root.left.right=new TreeNode(5);
        root.right.right=new TreeNode(7);
        root.right.left=new TreeNode(6);
        root.left.left.left=new TreeNode(8);
        root.left.left.right=new TreeNode(9);

        ArrayList<ArrayList<Stack<Integer>>> arrayLists = FindPath01(root, 4);
        Iterator<ArrayList<Stack<Integer>>> iterator = arrayLists.iterator();
        while (iterator.hasNext()) {
            ArrayList<Stack<Integer>> next = iterator.next();
            System.out.println(next);
        }

    }
}
