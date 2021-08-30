package Tree;

import DataStructure.TreeNode;
import com.sun.source.tree.Tree;
import org.w3c.dom.Node;

import java.util.*;

public class TraversalBinaryTree {
    // non-recursive
    public static List<Integer> preOrderTraversal (TreeNode root) {
        List<Integer> res = new ArrayList<>();
        Stack<TreeNode> stack = new Stack<>();
        stack.push(root);
        while (!stack.isEmpty()) {
            TreeNode node = stack.pop();
            if (node==null) {
                continue;
            }
            res.add(node.val);
            stack.push(node.right);
            stack.push(node.left);
        }
        return res;
    }

    public static List<Integer> inOrderTraversal (TreeNode root) {
        List<Integer> res = new ArrayList<>();
        if (root==null) {
            return res;
        }
        Stack<TreeNode> stack = new Stack<>();
        TreeNode cur= root;
        while (cur != null || !stack.isEmpty()) {
            while (cur != null) {
                stack.push(cur);
                cur=cur.left;
            }
            TreeNode node = stack.pop();
            res.add(node.val);
            cur=node.right;
        }
        return res;
    }

    public static List<Integer> postOrderTraveral(TreeNode root) {
        List<Integer> res = new ArrayList<>();
        if (root==null) {
            return res;
        }
        Stack<TreeNode> stack = new Stack<>();
        stack.push(root);
        while (!stack.isEmpty()) {
            TreeNode node = stack.pop();
            if (node==null) {
                continue;
            }
            res.add(node.val);
            stack.push(node.left);
            stack.push(node.right);

        }
        Collections.reverse(res);
        return res;
    }


    public static List<List<Integer>> levelOrder (TreeNode root) {
        List<List<Integer>> res = new ArrayList<>();
        if (root==null) {
            return res;
        }
        Queue<TreeNode> queue = new LinkedList<>();
        queue.add(root);
        while (!queue.isEmpty()) {
            int cnt = queue.size();
            List<Integer> list = new ArrayList<>();
            while (cnt-- > 0) {
                TreeNode node = queue.poll();
                if (node==null) {
                    continue;
                }
                list.add(node.val);
                queue.add(node.left);
                queue.add(node.right);
            }
            if (!list.isEmpty()) {
                res.add(list);
            }
        }
        return res;
    }

}


