package Array;
// 给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。

import java.util.Arrays;

public class _189_rotateArray {
    public static void main(String[] args) {
        int[] nums = {1, 2, 3, 4, 5, 6};
        new _189_rotateArray().rotate(nums, 7);
        System.out.println(Arrays.toString(nums));
    }

    public void rotate(int[] nums, int k) {

        if (k == 0 || nums == null || nums.length < 2) {
            return;
        }
        if (k > nums.length) {
            k %= nums.length;
        }

        reverseBlock(nums, 0, nums.length - 1 - k);
        reverseBlock(nums, nums.length - k, nums.length - 1);
        reverseBlock(nums, 0, nums.length - 1);
    }

    private void reverseBlock(int[] nums, int i, int j) {
        for (; i < j; i++, j--) {
            int temp = nums[i];
            nums[i] = nums[j];
            nums[j] = temp;
        }
    }
}
