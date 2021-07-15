package Tree;

import DataStructure.TreeNode;

public class _100_SameTree {
    public static boolean isSameTree(TreeNode p, TreeNode q) {
        if (p == null || q == null) {
            if (p == null && q == null) {
                return true;
            } else {
                return false;
            }
        }

        if (p.val != q.val) {
            return false;
        }
        return isSameTree(p.left, q.left) && isSameTree(p.right, p.right);
    }
}
