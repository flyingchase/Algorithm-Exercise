package Tree;

import DataStructure.TreeNode;

@SuppressWarnings({"all"})

public class reConstructBinaryTree07 {
    public static TreeNode reConstructBinaryTree01(int[] pre, int[] in) {
        if (pre==null||in==null||pre.length == 0||in.length==0||pre.length!=in.length) {
            return null;
        }
        return new TreeNode(-1);
    }

    public static TreeNode reBuild(int[] pre, int i, int j, int[] in, int m, int n) {
        int rootVal=pre[i],index=findIndex(rootVal,in,m,n);
        if(index<0) {
            return null;
        }

        int leftNodes=index-m,rightNodes = n-index;
        TreeNode root = new TreeNode(rootVal);
        if (leftNodes==0) {
            root.left=null;
        }else {
            root.left=reBuild(pre,i+1,i+leftNodes,in,m,m+leftNodes-1);
        }
        if (rightNodes==0) {
            root.right=null;
        }else {
            root.right=reBuild(pre, i+1, i+leftNodes+1, in,n-rightNodes+1,n);
        }

        return root;

    }

    public static int findIndex(int target, int[] arr, int from, int to) {
        for (int i=from;i<=to;i++) {
            if (arr[i] == target) {
                return i;
            }
        }
        return - 1;
    }

    public static void main(String[] args) {

    }
}

