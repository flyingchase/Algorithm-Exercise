package Tree;

import DataStructure.TreeNode;

public class _110_IsBalancedTree {
    private static boolean flags = true;
    public static boolean isBalancedTree(TreeNode root) {
        if (root==null) {
            return true;
        }
        return flags;
    }

    public static int maxDepth (TreeNode root) {
        if (root==null) {
            return 0;
        }

        int lDepth = maxDepth(root.left);
        int rDepth = maxDepth(root.right);
        if (Math.abs(lDepth-rDepth)>1) {
            flags=false;
        }
        return 1+Math.max(lDepth,rDepth);
    }
}
