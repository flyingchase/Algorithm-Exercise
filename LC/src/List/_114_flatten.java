package List;

import DataStructure.TreeNode;

import java.util.ArrayList;
import java.util.Stack;

public class _114_flatten {
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

// preOrder: 1 2 4 8 9 5 3 6 7
// inOrder: 8 4 9 2 5 1 6 3 7
// postOrder: 8 9 4 5 2 6 7 3 1

        new _114_flatten().flatten(root);
        while (root != null) {
            System.out.println(root.val);
            root = root.right;
        }
    }

    public void flatten(TreeNode root) {
        if (root==null) {
            return;
        }
        while (root != null) {
            if (root.left!=null) {
                TreeNode right = root.right;
                root.right=root.left;
                root.left=null;
                TreeNode node = root.right;
                while (node != null && node.right != null) {
                    node=node.right;
                }
                node.right=right;

            }
            root=root.right;
        }
    }
}
