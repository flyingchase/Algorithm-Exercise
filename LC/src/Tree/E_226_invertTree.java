package Tree;

import DataStructure.TreeNode;

import java.awt.image.ComponentSampleModel;
import java.util.ArrayList;
import java.util.List;
import java.util.Stack;

/*翻转一棵二叉树。

 */
public class E_226_invertTree {

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

// preOrder: 1 2 4 8 9 5 3 6 7
// inOrder: 8 4 9 2 5 1 6 3 7
// postOrder: 8 9 4 5 2 6 7 3 1


        TreeNode res = new E_226_invertTree().invertTree(root);

        System.out.println(new E_226_invertTree().inOrderTraversalBT(res).toString());

    }

    public TreeNode invertTree(TreeNode root) {

        if (root == null) {
            return null;
        }

        final TreeNode left = root.left,
                right = root.right;
        root.left = invertTree(right);
        root.right = invertTree(left);
        return root;
    }


    public void swap(int i, int j) {
        int temp = i;
        i = j;
        j = temp;
    }

    public List<Integer> inOrderTraversalBT(TreeNode root) {
        List<Integer> res = new ArrayList<>();
        //空树则直接返回null的res
        if (root == null) {
            return res;
        }
        Stack<TreeNode> stack = new Stack<>();
        // 保存root
        TreeNode cur = root;
        while (cur != null || !stack.isEmpty()) {
            while (cur != null) {
                stack.add(cur);
                cur = cur.left;
            }
            TreeNode node = stack.pop();
            res.add(node.val);
            cur = node.right;

        }

        return res;
    }
}
