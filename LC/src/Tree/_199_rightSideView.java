package Tree;

import DataStructure.TreeNode;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;
import java.util.Queue;

public class _199_rightSideView {
    public static void main(String[] args) {
        TreeNode root = new TreeNode(1);
        root.left = new TreeNode(2);
        root.right = new TreeNode(3);
        root.left.left = new TreeNode(4);
        root.left.right = new TreeNode(5);
        root.right.right = new TreeNode(7);
        root.right.left = new TreeNode(6);
        root.left.left.left = new TreeNode(8);
        root.left.left.right = new TreeNode(9);

        System.out.println(new _199_rightSideView().rightSideView(root));
    }

    public List<Integer> rightSideView(TreeNode root) {
        List<Integer> res = new ArrayList<>();
        if (root == null) {
            return res;
        }

        Queue<TreeNode> queues = new LinkedList<>();
        TreeNode cur = root;
        queues.add(cur);
        while (!queues.isEmpty()) {
            int size = queues.size();
            List<Integer> integers = new LinkedList<>();
            while (size-- > 0) {
                TreeNode node = queues.poll();
                if (node == null) {
                    continue;
                }
                integers.add(node.val);
                queues.add(node.left);
                queues.add(node.right);
            }
            if (!integers.isEmpty()) {

                Object[] objects = integers.toArray();
                res.add((Integer) objects[objects.length-1]);
            }
        }

        return res;
    }

}
