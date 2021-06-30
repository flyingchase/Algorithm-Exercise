package Tree;

import DataStructure.TreeNode;

public class isSymmentrical28 {
    public static boolean IsSymmentrical28(TreeNode pRoot) {
        if (pRoot==null) {
            return true;
        }
    return IsSymmentrical28(pRoot.left, pRoot.right);
    }

    public static boolean IsSymmentrical28(TreeNode t1, TreeNode t2) {
        if (t1==null&&t2==null) {
            return true;
        }
        if (t1==null||t2==null) {
            return false;
        }
        if (t1.val == t2.val) {
            return false;
        }

            return IsSymmentrical28(t1.left, t2.right) && IsSymmentrical28(t1.right,t2.left);

    }





}
