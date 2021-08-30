package Tree;

import DataStructure.TreeNode;

/*给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 高度平衡 二叉搜索树。

高度平衡 二叉树是一棵满足「每个节点的左右两个子树的高度差的绝对值不超过 1 」的二叉树。*/
public class E_108_sortedArrayToBST {
    public TreeNode sortedArrayToBST(int[] nums) {
        if (nums.length == 0 || nums == null) {
            return null;
        }
        return helper(nums, 0, nums.length - 1);
    }

    private TreeNode helper(int[] nums, int l, int r) {
        if (l > r) {
            return null;
        }
        int mid = l + ((r - l) >> 1);
        TreeNode node = new TreeNode(nums[mid]);
        node.left = helper(nums, l, mid-1 );
        node.right = helper(nums, mid + 1, r);
        return node;
    }


}
