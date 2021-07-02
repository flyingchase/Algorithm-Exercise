package Tree;

import DataStructure.TreeNode;

public class _54_KthNodeInBST {

    private static int count =0 ;
    public static TreeNode KthNodeInBST(TreeNode root, int k) {
        if (k==0||root==null) {
            return null;
        }

        TreeNode retNode = KthNodeInBST(root.left, k);
        if (retNode!=null) {
            return retNode;
        }

        count++;
        if (count==k) {
            return root;
        }

        retNode=KthNodeInBST(root.right,k);
        if (retNode!=null) {
            return retNode;
        }

        return null;
    }


    public static void main(String[] args) {
        TreeNode root = new TreeNode(4);
        root.left=new TreeNode(2);
        root.right=new TreeNode(6);
        root.left.left=new TreeNode(1);
        root.left.right=new TreeNode(3);
        root.right.right=new TreeNode(7);
        root.right.left=new TreeNode(5);

        TreeNode KthNode = KthNodeInBST(root, 4);
        System.out.println(KthNode.val);


    }
}
