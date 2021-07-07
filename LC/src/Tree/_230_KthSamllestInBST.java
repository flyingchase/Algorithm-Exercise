package Tree;

import DataStructure.TreeNode;

import java.util.Stack;

public class _230_KthSamllestInBST {
    public static void main(String[] args) {
        TreeNode root = new TreeNode(4);
        root.left = new TreeNode(2);
        root.right = new TreeNode(6);
        root.left.left = new TreeNode(1);
        root.left.right = new TreeNode(3);
        root.right.right = new TreeNode(7);
        root.right.left = new TreeNode(5);

        System.out.println(KthSmallestBST(root, 2));
    }

    public static int KthSmallestBST(TreeNode root, int k) {
        if (root == null || k <= 0) {
            return Integer.parseInt(null);
        }
        int Kth=0;
        Stack<TreeNode> stack = new Stack<>();
        TreeNode cur = root;
        while (cur != null || !stack.isEmpty()) {
            while (cur != null) {
                stack.push(cur);
                cur = cur.left;
            }
            TreeNode node = stack.pop();
            if (--k <= 0) {
                Kth = node.val;
//                break;
                return Kth;
            }
            cur = node.right;
        }
        return Kth;
    }
}
