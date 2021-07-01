package Tree;

import DataStructure.TreeNode;

import java.lang.invoke.VarHandle;
import java.util.*;

public class _637_AverageOfLevelsInBinaryTree {

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

        List<Double> res = AverageOfLevelBinaryTree(root);
        System.out.println(res);
    }

    public static List<Double> AverageOfLevelBinaryTree (TreeNode root) {
        List<Double> res = new ArrayList<>();
        Queue<TreeNode> queue = new LinkedList<>();
        if (root==null) {
            return res;
        }

        queue.add(root);
        while (!queue.isEmpty()) {
            int count = queue.size();
            float list = 0.0F;
            int cnt =count;
            while (count-- > 0) {
                TreeNode node = queue.poll();
                if (node==null) {
                    continue;
                }
                list+=node.val;
                queue.add(node.left);
                queue.add(node.right);
            }
            if (list!=0) {
                res.add(((double)list/cnt));
            }
        }
        return res;
    }


}




































