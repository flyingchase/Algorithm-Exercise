package Tree;

import DataStructure.TreeNode;

public class MirrorTree26 {
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
