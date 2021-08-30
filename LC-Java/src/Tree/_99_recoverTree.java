package Tree;

import DataStructure.TreeNode;
import org.w3c.dom.Node;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.Queue;
import java.util.Stack;

public class _99_recoverTree {
    public static void main(String[] args) {
        TreeNode root = new TreeNode(4);
        root.left = new TreeNode(2);
        root.right = new TreeNode(7);
        root.left.left = new TreeNode(1);
        root.left.right = new TreeNode(3);
        root.right.right = new TreeNode(6);
        root.right.left = new TreeNode(5);

        System.out.println(PrintLevelBST(root));
        recoverBST(root);
        System.out.println(PrintLevelBST(root));

    }

    public static void recoverBST(TreeNode root) {

        if (root == null) {
            return;
        }
        Stack<TreeNode> stack = new Stack<>();
        TreeNode cur = root;
        TreeNode firstNode = null;
        TreeNode secondNode = null;
        TreeNode preNode = new TreeNode(Integer.MIN_VALUE);


        while (cur != null || !stack.isEmpty()) {
            while (cur != null) {
                stack.push(cur);
                cur = cur.left;
            }
            TreeNode node = stack.pop();
            if (firstNode == null && preNode.val > node.val) {
                firstNode = preNode;
            }
            if (firstNode != null && preNode.val > node.val) {
                secondNode = node;
            }
            preNode = node;
            cur = node.right;
        }
        if (firstNode != null && secondNode != null) {

            int temp = firstNode.val;
            firstNode.val = secondNode.val;
            secondNode.val = temp;
        }
    }

    public static ArrayList<ArrayList<Integer>> PrintLevelBST(TreeNode root) {
        ArrayList<ArrayList<Integer>> res = new ArrayList<>();
        if (root == null) {
            return res;
        }
        Queue<TreeNode> queues = new LinkedList<>();
        queues.add(root);
        TreeNode cur =root;
        while (!queues.isEmpty()) {
            ArrayList<Integer> list = new ArrayList<>();
            int cnt = queues.size();
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
        return res;
    }
}

