package Tree;

import DataStructure.TreeNode;

public class _543_DiameterOfBinarytree {
    private  static int max=0;

    public static int DiameterOfBinaryTree (TreeNode root) {
        if (root==null) {
            return 0;
        }
        depth(root);
        return max;
    }

    private static int depth (TreeNode root) {
        int lDepth = depth(root.left);
        int rDepth = depth(root.right);
        max = Math.max(rDepth+lDepth,max);
        return Math.max(rDepth, lDepth) + 1;
    }
}
