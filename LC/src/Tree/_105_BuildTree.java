package Tree;

import DataStructure.TreeNode;

import java.util.Arrays;

import static Tree._100_SameTree.isSameTree;

public class _105_BuildTree {
    /*
     * 从前序和后序遍历结果构造二叉树
     * */
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

        int[] preorder = {1, 2, 4, 8, 9, 5, 3, 6, 7};
        int[] inorder = {8, 4, 9, 2, 5, 1, 6, 3, 7};

        TreeNode root2 = buildTree(preorder, inorder);

        System.out.println(isSameTree(root2, root));

    }

    public static TreeNode buildTree(int[] preorder, int[] inorder) {
        if (preorder.length == 0) {
            return null;
        }
        assert (preorder.length == inorder.length);

        TreeNode root = new TreeNode(preorder[0]);

        for (int i = 0; i < preorder.length; i++) {
            // 找到root在中序遍历数组中的位置
            if (preorder[0] == inorder[i]) {
                /*将前 中序遍历数组各分为两半*/
                int[] preLeft = Arrays.copyOfRange(preorder, 1, i + 1);
                int[] preRight = Arrays.copyOfRange(preorder, i + 1, preorder.length);
                int[] inLeft = Arrays.copyOfRange(inorder, 0, i);
                int[] inRight = Arrays.copyOfRange(inorder, i + 1, inorder.length);
                root.left = buildTree(preLeft, inLeft);
                root.right = buildTree(preRight, inRight);
                break;
            }

        }
        return root;

    }
}
