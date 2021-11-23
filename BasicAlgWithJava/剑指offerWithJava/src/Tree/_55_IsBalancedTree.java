package Tree;

import DataStructure.TreeNode;

public class _55_IsBalancedTree {
    private static boolean isBalanced;

    public static boolean isBalancedTree(TreeNode root) {
        if (root==null) {
            return true;
        }

        isBalanced=true;
        treeDepth(root);
        return isBalanced;
    }

    private static int treeDepth (TreeNode  root) {
        if (root==null||!isBalanced) {
            return 0;
        }
        int lDepth = treeDepth(root.left);
        int rDepth = treeDepth(root.right);
        if (Math.abs(lDepth-rDepth)>1) {
            isBalanced=false;
        }
        return 1+Math.max(lDepth,rDepth);
    }

    public static void main(String[] args) {
        TreeNode root = new TreeNode(1);
        root.left=new TreeNode(2);
        root.right=new TreeNode(3);
        root.left.left=new TreeNode(4);
        root.left.right=new TreeNode(5);
        root.right.right=new TreeNode(7);
        root.right.left=new TreeNode(6);
        root.left.left.left=new TreeNode(8);
        root.left.left.right=new TreeNode(9);
        root.left.left.right.left=new TreeNode(9);

        if (isBalancedTree(root)) {
            System.out.println("yes It is balanced");
        }else {
            System.out.println("No It is unbalanced");
        }
    }
}


