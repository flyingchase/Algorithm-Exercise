package Tree;

import DataStructure.TreeNode;

import java.util.Stack;

/*
* 找到BST中给定结点的后继节点 在中序遍历中最接近的node
*
*
* */
public class _285_inOrderSuccessor {

    public static void main(String[] args) {
        TreeNode root = new TreeNode(4);
        root.left=new TreeNode(2);
        root.right=new TreeNode(6);
        root.left.left=new TreeNode(1);
        root.left.right=new TreeNode(3);
        root.right.right=new TreeNode(7);
        root.right.left=new TreeNode(5);
    }
    public static TreeNode inOrderSuccessorBST(TreeNode root, TreeNode p) {
        TreeNode find = null;
        if (root==null||p==null) {
            return find;
        }
        Stack<TreeNode> stack = new Stack<>();
        TreeNode cur = root;
        TreeNode smallestTopsmal = new TreeNode(Integer.MAX_VALUE);
        while (cur != null || !stack.isEmpty()) {
            while (cur != null) {
                stack.add(cur);
                cur=cur.left;
            }
            TreeNode node = stack.pop();

            cur=node.right;
        }
        return find;

    }

}
