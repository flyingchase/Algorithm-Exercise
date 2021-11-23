package Tree;

import DataStructure.TreeNode;

public class HasSubTree26 {
    public static boolean hasSubTre(TreeNode root1, TreeNode root2) {
        if (root2==null||root1==null) {
            return false;
        }
        return process(root1, root2);
    }

    public static boolean process(TreeNode root1, TreeNode root2) {
        if (root2==null) {
            return true;
        }
        if (root1==null && root2!=null) {
            return false;
        }
        if (root1.val==root2.val) {
            if (process(root1.left, root2.left) && process(root1.right,root2.right)) {
                return true;
            }
        }

        return process(root1.left, root2) || process(root2.right, root1);

    }

    public static void main(String[] args) {

    }
}




