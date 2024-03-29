////package offerWithJava
//public class ReConstructBinaryTree {
//    //二叉树的定义
//
//    // LinkNode n = new LinkNode();
//    // LinkNode h
//    private TreeNode constructBinaryTree(int[] pre, int startPre, int endPre, int[] in, int startIn, int endIn) {
//        TreeNode node = new ReConstructBinaryTree().new TreeNode(pre[startPre]);
//        if (startPre == endPre) {
//            if(startIn == endIn) return node;
//            throw new IllegalArgumentException("Invalid Input!");
//        }
//        int inOrder = startIn;
//        while (in[inOrder]!=pre[startPre]) {
//            ++inOrder;
//            if(inOrder>endIn) new IllegalArgumentException("Invalid Input!");
//
//        }
//        int len = inOrder-startIn;
//        if(len>0){
//            //递归构建左侧树
//            node.left =  constructBinaryTree(pre, startPre, endPre, in, startIn, endIn);
//        }
//        if(inOrder<endIn){
//            //递归构建右侧树
//            node.right =  constructBinaryTree(pre, startPre+len+1, endPre, in, inOrder+1, endIn);
//        }
//        return node;
//
//    }
//    private TreeNode reConstructBinaryTree(int[] pre, int[] in) {
//        if(pre==null||in==null||pre.length!=in.length)   return null;
//        int n=pre.length;
//        return constructBinaryTree(pre, 0, n-1, in, 0, n-1);
//
//    }
//
//    public static void main(String[] args) {
//        int[] pre={1,2,4,7,3,5,6,8};
//        int[] in={4,7,2,1,5,3,8,6};
//        ReConstructBinaryTree reConstructBinaryTree1=new ReConstructBinaryTree();
//        TreeNode tree1= reConstructBinaryTree1.reConstructBinaryTree(pre, in);
//
//    }
//
//
//
//
//}
