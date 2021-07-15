package Tree;

import DataStructure.TreeNode;

public class _26_MirrorTree {
    public static void MirrorTree(TreeNode root ) {
        if (root==null) {
            return;
        }

        TreeNode tmp =root.left;
        root.left=root.right;
        root.right=tmp;
        MirrorTree(root.left);
        MirrorTree(root.right);
    }


}
