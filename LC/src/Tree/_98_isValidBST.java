package Tree;

import DataStructure.TreeNode;

public class _98_isValidBST {
    public static void main(String[] args) {
        TreeNode root = new TreeNode(1);


        root.left=new TreeNode(2);
//        root.right=new TreeNode(3);
        root.left.left=new TreeNode(4);
//        root.left.right=new TreeNode(5);
//        root.right.right=new TreeNode(7);
//        root.right.left=new TreeNode(6);
//        root.left.left.left=new TreeNode(8);
//        root.left.left.right=new TreeNode(9);

        if (isValidBST(root)) {
            System.out.println("yes");
        }else {
            System.out.println("No");
        }
    }

    public static boolean isValidBST (TreeNode root) {
return true;
    }

//    public static ArrayList<Integer> inOrderTraversalBST(TreeNode root) {
//        ArrayList<Integer> res = new ArrayList<>();
//        if (root==null) {
//            return res;
//        }
//        Stack<TreeNode> stack = new Stack<>();
//        TreeNode cur=root;
//        while (cur != null || !stack.isEmpty()) {
//            while (cur != null) {
//                stack.push(cur);
//                cur=cur.left;
//            }
//            TreeNode node= stack.pop();
//            int min = node.val;
//
//            res.add(node.val);
//            cur=node.right;
//
//        }
//
////        while (root != null || !stack.isEmpty()) {
////            while (root != null) {
////                res.add(root.val);
////                stack.push(root);
////                root=root.left;
////            }
////            TreeNode cur = stack.pop();
////            root=cur.right;
////        }
//    }


}
