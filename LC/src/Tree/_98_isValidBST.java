package Tree;

import DataStructure.TreeNode;
import com.sun.source.tree.Tree;

import java.util.ArrayList;
import java.util.Stack;

@SuppressWarnings({"all"})
public class _98_isValidBST {
    public static void main(String[] args) {
        TreeNode root = new TreeNode(4);
        root.left=new TreeNode(2);
        root.right=new TreeNode(6);
        root.left.left=new TreeNode(1);
        root.left.right=new TreeNode(3);
        root.right.right=new TreeNode(7);
        root.right.left=new TreeNode(5);

        if (isVaildBST(root)) {
            System.out.println("yes");
        }else {
            System.out.println("No");
        }
    }

    public static boolean isVaildBST (TreeNode root) {
        boolean flag = true;
        if (root==null) {
            return flag;
        }
        Stack<TreeNode> stack = new Stack<>();
        TreeNode cur = root;
        Integer preNodeValue = null;

        while (cur != null || !stack.isEmpty()) {
            while (cur != null) {
                stack.push(cur);
                cur=cur.left;
            }
            TreeNode node = stack.pop();
            if (preNodeValue!=null&&preNodeValue.intValue()>=node.val) {
                flag=false;
                break;
            } else {
                preNodeValue=(Integer) node.val;
            }
            cur=node.right;
        }
        return flag;
    }


}
