package Tree;

import DataStructure.TreeNode;

import java.util.Arrays;
import java.util.Collections;
import java.util.PriorityQueue;

@SuppressWarnings({"all"})
public class _106_BulidTree02 {
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


        int[] inorder = {8, 4, 9, 2, 5, 1, 6, 3, 7};
        int[] postorder = {8, 9, 4, 5, 2, 6, 7, 3, 1};

        TreeNode root2 = buildTree(inorder, postorder);
        System.out.println(_100_SameTree.isSameTree(root2, root));
    }

    public static TreeNode buildTree(int[] inorder, int[] postorder) {
        if (inorder == null || postorder == null) {
            return null;
        }
        return helper(inorder, postorder);
    }

    private static TreeNode helper(int[] in, int[] post) {
        if (in.length == 0) {
            return null;
        }
        //根据后序数组的最后一个元素，创建根节点
        TreeNode root = new TreeNode(post[post.length - 1]);
        //在中序数组中查找值等于【后序数组最后一个元素】的下标
        for (int i = 0; i < in.length; ++i) {
            if (in[i] == post[post.length - 1]) {
                //确定这个下标i后，将中序数组分成两部分，后序数组分成两部分
                int[] inLeft = Arrays.copyOfRange(in, 0, i);
                int[] inRight = Arrays.copyOfRange(in, i + 1, in.length);
                int[] postLeft = Arrays.copyOfRange(post, 0, i);
                int[] postRight = Arrays.copyOfRange(post, i, post.length - 1);
                //递归处理中序数组左边，后序数组左边
                root.left = helper(inLeft, postLeft);
                //递归处理中序数组右边，后序数组右边
                root.right = helper(inRight, postRight);
                break;
            }
        }
        return root;
    }

    PriorityQueue<Integer> minheap = new PriorityQueue<>();
    PriorityQueue<Integer> maxheap = new PriorityQueue<>(Collections.reverseOrder());

//    minheap.add();
//    peek() 查看堆顶元素
//    poll() 弹出堆顶元素并删除
//    size() 堆的大小


}
