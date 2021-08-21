package Tree;

import DataStructure.ListNode;
import DataStructure.TreeNode;
import com.sun.source.tree.Tree;

import java.net.HttpURLConnection;
import java.sql.ConnectionBuilder;
import java.util.Deque;
import java.util.LinkedList;
import java.util.List;
import java.util.Queue;

/*镜像二叉树*/
public class E_101_isSymmetric {

    public static void main(String[] args) {
        // TreeNode root = new TreeNode(1);
        /*root.left=new TreeNode(2);
        root.right=new TreeNode(3);
        root.left.left=new TreeNode(4);
        root.left.right=new TreeNode(5);
        root.right.right=new TreeNode(7);
        root.right.left=new TreeNode(6);
        root.left.left.left=new TreeNode(8);
        root.left.left.right=new TreeNode(9);
*/
// preOrder: 1 2 4 8 9 5 3 6 7
// inOrder: 8 4 9 2 5 1 6 3 7
// postOrder: 8 9 4 5 2 6 7 3 1


        System.out.println(new E_101_isSymmetric().isSymmetric(new TreeNode()));
    }

    // Divide 分治
    // public boolean isSymmetric (TreeNode root) {
    //     if (root==null) {
    //         return true;
    //     }
    //     TreeNode cur = root;
    //     return isSymmetric(cur.left,cur.right);
    // }
    //
    // private boolean isSymmetric(TreeNode lRoot, TreeNode rRoot) {
    //
    //     // 上述两个 if 可以替换成
    //     if(lRoot==null || rRoot==null) {
    //         return lRoot==rRoot;
    //     }
    //     if (lRoot.val!=rRoot.val) {
    //         return false;
    //     }
    //
    //     return isSymmetric(lRoot.left,rRoot.right)&&isSymmetric(rRoot.left,lRoot.right);
    // }


    public boolean isSymmetric (TreeNode root) {
        if (root== null) {
            return true;
        }

        Deque<TreeNode> queue = new LinkedList<>();
        queue.add(root.left);
        queue.add(root.right);
        while (!queue.isEmpty()) {
            TreeNode left = queue.poll();
            TreeNode right = queue.poll();
            if (left==null&&right==null) {
                continue;
            }
            if (left==null||right==null) {
                return false;
            }
            if (left.val!=right.val) {
                return  false;
            }
            queue.add(left.left);
            queue.add(right.right);
            queue.add(left.right);
            queue.add(right.left);
        }

        return true;
    }


    // public List<List<Integer>> levelTravelBT (TreeNode root) {
    //     if (root==null) {
    //         return null;
    //     }
    //
    //     List<List<Integer>> res = new LinkedList<>();
    //     Queue<TreeNode> queue = new LinkedList<>();
    //     TreeNode cur = root;
    //     queue.add(cur);
    //
    //     while (!queue.isEmpty()) {
    //         List<Integer> list = new LinkedList<>();
    //         TreeNode node = queue.poll();
    //         if (node==null) {
    //             continue;
    //         }
    //         list.add();
    //
    //     }
    //
    //
    // }



}
