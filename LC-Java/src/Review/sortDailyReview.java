package Review;

import DataStructure.TreeNode;

import java.util.LinkedList;
import java.util.List;
import java.util.Queue;

import static DataStructure.swap.swap;

/**
 * @author :qlzhou;
 * @date: :创建于2021/9/39:27 下午
 */
public class sortDailyReview {
    public void quicksort(int[] nums) {
        if (nums == null) {
            return;
        }
        quicksort(nums, 0, nums.length - 1);

    }

    private void quicksort(int[] nums, int l, int r) {
        if (l > r) {
            return;
        }
        swap(nums, (int) (l + Math.random() * (r - l + 1)), r);
        int[] p = paratitions(nums, l, r);
        quicksort(nums, l, p[0] - 1);
        quicksort(nums, p[1] + 1, r);


    }

    private int[] paratitions(int[] nums, int l, int r) {
        int less = l - 1, more = r;
        while (l < more) {
            if (nums[l] > nums[r]) {
                swap(nums, ++less, l++);
            } else if (nums[l] < nums[r]) {
                swap(nums, l, --more);
            } else {
                l++;
            }
        }
        swap(nums, r, more);
        return new int[]{
                less + 1, more
        };
    }


    public void heapsort(int[] nums) {
        if (nums == null) {
            return;
        }
        for (int i = 0; i < nums.length; i++) {
            heapinsert(nums, i);
        }

        int size = nums.length;
        while (size-- > 0) {
            swap(nums, 0, --size);
            heapIfy(nums, 0, size);
        }
    }

    private void heapIfy(int[] nums, int index, int size) {
        int left = 2 * index + 1;
        while (left < size) {
            int largest = left + 1 < size && nums[left + 1] > nums[left] ? left + 1 : left;
            largest = nums[largest] > nums[index] ? largest : index;
            if (largest == index) {
                break;
            }
            swap(nums, largest, index);
            index = largest;
            left = index * 2 + 1;
        }
    }

    private void heapinsert(int[] nums, int index) {
        while (nums[index] < nums[2 * index - 1]) {
            swap(nums, index, 2 * index - 1);
            index = 2 * index - 1;
        }
    }



    public List<List<Integer>> zigzalevelTravelBT (TreeNode root) {
        if (root==null) {
            return null;
        }
        List<List<Integer>> res = new LinkedList<>();

        Queue<TreeNode> queue = new LinkedList<>();

        TreeNode cur = root;
        queue.add(cur);

        while (!queue.isEmpty()) {
            int size = queue.size();
            LinkedList<Integer> lists = new LinkedList<>();
            while (size-- > 0) {
                TreeNode node = queue.poll();
                if (node==null) {
                    continue;
                }
                lists.add(node.val);
                queue.add(node.right);
                queue.add(node.left);
            }
        }

        return res;


    }
}
