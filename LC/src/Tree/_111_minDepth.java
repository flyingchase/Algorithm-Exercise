package Tree;

import DataStructure.TreeNode;

@SuppressWarnings({"all"})
public class _111_minDepth {
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

        System.out.println(new _111_minDepth().minDepth(root));
    }

    public int minDepth(TreeNode root) {
        if (root == null) {
            return 0;
        }
        if (root.left == null || root.right == null) {
            if (root.left == null && root.right == null) {
                return 1;
            } else {
                return 1 + (root.left == null ? minDepth(root.right) : minDepth(root.left));
            }
        }
        return 1 + Math.min(minDepth(root.right), minDepth(root.left));

    }
}
