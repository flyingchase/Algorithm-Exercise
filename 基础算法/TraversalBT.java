import DataStructure.TreeNode;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
import java.util.Stack;

public class TraversalBT {
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

        List<Integer> resPreOrder = preOrderTraversalBT(root);
        System.out.println(resPreOrder);

        List<Integer> resInOrder = inOrderTraversalBT(root);
        System.out.println(resInOrder);

        List<Integer> resPostOrder = postOrderTraversalBT(root);
        System.out.println(resPostOrder);





    }


    public static List<Integer> inOrderTraversalBT(TreeNode root) {
        List<Integer> res = new ArrayList<>();
        //空树则直接返回null的res
        if (root == null) {
            return res;
        }
        Stack<TreeNode> stack = new Stack<>();
        // 保存root
        TreeNode cur = root;
        while (cur != null || !stack.isEmpty()) {
            // 向树的最左侧走 并不断压栈
            while (cur != null) {
                stack.push(cur);
                cur = cur.left;
            }
            // 走完最左侧后弹栈顶
            TreeNode node = stack.pop();
            res.add(node.val);
            // 往上层的右边去再重复找最左侧
            cur = node.right;
        }
        return res;
    }


    public static List<Integer> preOrderTraversalBT(TreeNode root) {
        List<Integer> res = new ArrayList<>();
        //空树则直接返回null的res
        if (root == null) {
            return res;
        }
        Stack<TreeNode> stack = new Stack<>();
        // 保存root
        TreeNode cur = root;
        while (cur != null || !stack.isEmpty()) {
            // 向树的最左侧走 并不断压栈入数组
            while (cur != null) {
                res.add(cur.val);
                stack.push(cur);
                cur = cur.left;
            }
            // 走完最左侧后弹栈顶 无须入组 cur循环内已经入组
            TreeNode node = stack.pop();
            // 往上层的右边去 再重复找最左侧
            cur = node.right;
        }

        return res;
    }


    public static List<Integer> postOrderTraversalBT(TreeNode root) {
        List<Integer> res = new ArrayList<>();
        //空树则直接返回null的res
        if (root == null) {
            return res;
        }
        Stack<TreeNode> stack = new Stack<>();
        // 保存root
        TreeNode cur = root;
        while (cur != null || !stack.isEmpty()) {
            // 向树的最右侧走 并不断压栈入数组
            while (cur != null) {
                res.add(cur.val);
                stack.push(cur);
                cur = cur.right; // 使得reverse后正常左在前
            }
            // 走完最右侧后弹栈顶 无须入组 cur循环内已经入组
            TreeNode node = stack.pop();
            // 往上层的左边去 再重复找最右侧
            cur=node.left;
        }
        Collections.reverse(res);
        return res;
    }
}