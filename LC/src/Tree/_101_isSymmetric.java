package Tree;

import DataStructure.TreeNode;

public class _101_isSymmetric {
    public static void main(String[] args) {
        TreeNode root = new TreeNode(1);
        root.left = new TreeNode(2);
        root.right = new TreeNode(2);
        root.left.left = new TreeNode(4);
        root.left.right = new TreeNode(5);
        root.right.right = new TreeNode(4);
        root.right.left = new TreeNode(5);
//        root.left.left.left = new TreeNode(8);
//        root.left.left.right = new TreeNode(9);

        System.out.println(new _101_isSymmetric().isSymmetric(root));
    }

    public boolean isSymmetric(TreeNode root) {
        if (root == null) {
            return true;
        }

        if (root.left.val != root.right.val) {
            return false;
        }
        return isSymmetric(root.left) && isSymmetric(root.right);
    }

    public boolean isSymmetric(TreeNode node1, TreeNode node2) {
        if (node1 == null || node2 == null) {
            if (node1 != null || node2 != null) {
                return false;
            }
            return true;
        }

        if (node1.val != node2.val) {
            return false;
        }
        return isSymmetric(node1.left, node2.right) && isSymmetric(node1.right, node2.left);
    }
}
