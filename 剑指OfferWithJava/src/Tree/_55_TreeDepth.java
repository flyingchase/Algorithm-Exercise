package Tree;

import DataStructure.TreeNode;

public class _55_TreeDepth {
    public static int TreeDepth (TreeNode root) {
        if (root==null) {
            return 0;
        }
        int lDepth = TreeDepth(root.left);
        int rDepth = TreeDepth(root.right);
        return 1 + Math.max(lDepth, rDepth);
    }

    public static int TreeDepth02 (TreeNode root) {
        return root == null ? 0 : 1 + Math.max(TreeDepth02(root.left), TreeDepth02(root.right));
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

//        System.out.println(TreeDepth(root));

        System.out.println(TreeDepth02(root));
    }
}
