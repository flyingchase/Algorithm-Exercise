package Tree;

import DataStructure.TreeNode;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;
import java.util.Queue;

public class _103_zigzagLevelOrder {
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


        List<List<Integer>> lists = levelOrderBottom(root);
        System.out.println(lists);
    }

    public static List<List<Integer>> levelOrderBottom(TreeNode root) {
        List<List<Integer>> res = new ArrayList<>();
        if (root==null) {
            return res;
        }
        Queue<TreeNode> queues = new LinkedList<>();
        queues.add(root);
        while (!queues.isEmpty()) {
            int cnt = queues.size();
            List<Integer> list = new ArrayList<>();
            while (cnt-- > 0) {
                TreeNode node = queues.poll();
                if (node==null) {
                    continue;
                }
                list.add(node.val);
                queues.add(node.left);
                queues.add(node.right);
            }
            if (!list.isEmpty()) {
                res.add(list);
            }
        }
//        Collections.reverse(res);
        return res;
    }
}
