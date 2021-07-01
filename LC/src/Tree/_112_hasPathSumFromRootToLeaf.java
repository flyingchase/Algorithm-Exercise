package Tree;

import DataStructure.TreeNode;

public class _112_hasPathSumFromRootToLeaf {
    public static boolean hasPathSumFromRootToLeaf(TreeNode root, int sum) {
        if (root==null) {
            return false;
        }
        if (root.val==sum&&root.left==null&&root.right==null) {
            return true;
        }

        return hasPathSumFromRootToLeaf(root.left, sum - root.val) || hasPathSumFromRootToLeaf(root.right, sum - root.val);
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

        if (hasPathSumFromRootToLeaf(root,15)) {
            System.out.println("Yes, has The path that the sum is equal to sum");
        }else {
            System.out.println("No");
        }
    }
}
