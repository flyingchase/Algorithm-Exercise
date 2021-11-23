package Tree;

import DataStructure.TreeNode;

import java.util.Stack;

public class _36_ConvertToLinkedList {
    public static TreeNode Convert01(TreeNode root) {
        if (root==null) {
            return null;
        }
        Stack<TreeNode> stack = new Stack<>();
        TreeNode cur = root;
        TreeNode res = null;
        TreeNode pre = null;
        while (cur != null || !stack.isEmpty()) {
            if (cur!=null) {
                stack.push(cur);
                cur=cur.left;
            }else {
                cur= stack.pop();
                if (pre==null) {
                    pre=cur;
                    res=pre;
                }else {
                    pre.right=cur;
                    cur.left=pre;
                    pre=cur;
                }
                cur=cur.right;
            }
        }
        return res;
    }


    public static TreeNode Convert02 (TreeNode root) {
        if (root==null) {
            return null;
        }
        TreeNode HeadList = root;
        while (HeadList.left != null) {
            HeadList=HeadList.left;
        }
        transform(root);
        return HeadList;
    }

    public static void transform(TreeNode root) {
        if (root==null) {
            return;
        }
        TreeNode preRoot = root.left,nextRoot = root.right;
        while (preRoot != null && preRoot.right != null) {
            preRoot=preRoot.right;
        }
        while (nextRoot != null && nextRoot.left != null) {
            nextRoot=nextRoot.left;
        }
        // recursively
        transform(root.left);
        transform(root.right);
        root.left=preRoot;
        if (preRoot!=null) {
            preRoot.right=root;
        }
        root.right=nextRoot;
        if (nextRoot!=null) {
            nextRoot.left=root;
        }
    }

    public static void main(String[] args) {
        TreeNode root = new TreeNode(4);
        root.left=new TreeNode(2);
        root.right=new TreeNode(6);
        root.left.left=new TreeNode(1);
        root.left.right=new TreeNode(3);
        root.right.right=new TreeNode(7);
        root.right.left=new TreeNode(5);
//        root.left.left.left=new TreeNode(8);
//        root.left.left.right=new TreeNode(9);
        TreeNode headList = Convert02(root);
        while (headList != null) {
            System.out.println(headList.val);
            headList=headList.right;
        }
    }
}



















