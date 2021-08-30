package Tree;

import DataStructure.TreeNode;

public class _889_constructFromPrePost {
    public static void main(String[] args) {
        TreeNode root = new TreeNode(1);
        root.left = new TreeNode(2);
        root.right = new TreeNode(3);
        root.left.left = new TreeNode(4);
        root.left.right = new TreeNode(5);
        root.right.right = new TreeNode(7);
        root.right.left = new TreeNode(6);
        root.left.left.left = new TreeNode(8);
        root.left.left.right = new TreeNode(9);


    }

    public static TreeNode constructFromPrePost (int[] pre, int[] post) {
        if (pre==null||post==null||post.length!=pre.length) {
            return null;
        }

        TreeNode root = new TreeNode(pre[0]);

        return null;
    }
}
