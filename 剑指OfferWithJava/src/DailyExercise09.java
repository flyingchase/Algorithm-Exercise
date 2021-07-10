import DataStructure.TreeNode;

import java.util.*;

import static DataStructure.swap.swap;

@SuppressWarnings({"all"})
public class DailyExercise09 {
    public static ArrayList<ArrayList<Integer>> ZigziaPrint(TreeNode root) {
        ArrayList<ArrayList<Integer>> res = new ArrayList<>();
        if (root == null) {
            return res;
        }
        Queue<TreeNode> queue = new Queue<>();
        queue.add(root);
        TreeNode cur = root;
        boolean reverse = false;
        while (!queue.isEmpty()) {
            int size = queue.size();
            ArrayList<Integer> list = new ArrayList<>();
            while (size-- > 0) {
                TreeNode node = queue.poll();
                if (node == null) {
                    continue;
                }
                list.add(node.val);
                queue.add(node.left);
                queue.add(node.right);

            }
            if (reverse) {
                Collections.reverse(list);
            }
            reverse = !reverse;
            if (!list.isEmpty()) {
                res.add(list);
            }
        }

        return res;
    }

    public static List<Integer> PreOrderTraversalBT(TreeNode root) {
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
                res.add(cur.val);
                stack.add(cur);
                cur = cur.right;
            }
            TreeNode node = stack.pop();
            cur = node.left;


        }

        return res;
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

    public static List<Integer> PostOrderTraversalBT(TreeNode root) {
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
                res.add(cur.val);
                stack.push(cur);
                cur = cur.left;
            }
            TreeNode node = stack.pop();
            cur = node.right;


        }
        Collections.reverse(res);
        return res;
    }

    public static void mergeSort(int[] nums) {
        if (nums == null || nums.length < 2) {
            return;
        }
        mergeSort(nums, 0, nums.length - 1);
    }

    public static void mergeSort(int[] nums, int l, int r) {
        if (l == r) {
            return;
        }
        int mid = l + ((r - l) >> 1);
        mergeSort(nums, l, mid);
        mergeSort(nums, mid + 1, r);
        merge(nums, l, mid, r);
    }

    public static void merge(int[] nums, int l, int mid, int r) {
        int i = 0, p1 = l, p2 = mid + 1;
        int[] helper = new int[r - l + 1];
        while (p1 <= mid && p2 <= r) {
            helper[i++] = nums[p1] < nums[p2] ? nums[p1++] : nums[p2++];
        }
        while (p1 <= mid) {
            helper[i++] = nums[p1++];
        }
        while (p2 <= r) {
            helper[i++] = nums[p2++];
        }
        for (i = 0; i < helper.length; i++) {
            nums[i + l] = helper[i];
        }
    }

    public static void heapSort(int[] nums) {
        if (nums == null || nums.length < 2) {
            return;
        }
        for (int i = 0; i < nums.length; i++) {
            heapInsert(nums, i);
        }
        int size = nums.length;
        while (size > 0) {
            swap(nums, i, --size);
            heapify(nums, 0, size);
        }

    }

    public static void heapInsert(int[] nums, int index) {
        while (nums[index] > nums[(index - 1) / 2]) {
            swap(nums, index, (index - 1) / 2);
            index = (index - 1) / 2;
        }
    }

    public static void heapify(int[] nums, int index, int size) {
        int left = 2 * index + 1;
        int right = left + 1;
        while (left < size) {
            int largest = right < size && nums[left] < nums[right] ? right : left;
            largest = nums[index] < nums[largest] ? largest : index;
            if (largest == index) {
                break;
            }
            swap(nums, index, largest);
            index = largest;
            left = index * 2 + 1;
            right = left + 1;
        }
    }


    public static void BubbleSort(int[] nums) {
        if (nums == null || nums.length < 2) {
            return;
        }
        for (int i = nums.length - 1; i > 0; i--) {
            for (int j = 0; j < i; j++) {
                if (nums[j] > nums[j + 1]) {
                    swap(nums, j, j + 1);
                }
            }
        }
    }

    public static void quickSort(int[] nums) {
        if (nums == null || nums.length < 2) {
            return;
        }
        quickSort(nums, 0, nums.length - 1);
    }

    public static void quickSort(int[] nums, int l, int r) {
        if (l < r) {
            swap(nums, l, r);
            int[] partition = partition(nums, l, r);
            quickSort(nums, l, partition[0] - 1);
            quickSort(nums, partition[1] + 1, r);
        }
    }

    public static int[] partition(int[] nums, int l, int r) {
        int less = l - 1, more = r;
        while (l < more) {
            if (nums[l] < nums[r]) {
                swap(nums, ++less, l++);
            } else if (nums[l] > nums[r]) {
                swap(nums, --more, l);
            } else {
                l++;
            }
        }
        swap(nums, more, r);
        return new int[]{less + 1, more};
    }


    public static void SelectionSort(int[] nums) {
        if (nums == null || nums.length < 2) {
            return;
        }
        for (int i = 0; i < nums.length - 1; i++) {
            int minIndex = i;
            for (int j = i + 1; j < nums.length; j++) {
                minIndex = nums[j] < nums[minIndex] ? j : minIndex;
            }
            swap(nums, i, minIndex);
        }
    }

    public static void InsertionSort(int[] nums) {
        if (nums==null||nums.length<2) {
            return;
        }
        for (int i = 1; i < nums.length; i++) {
            for (int j = i-1; j >0; j--) {
                if (nums[j]>nums[j+1]) {
                    swap(nums,j,j+1);
                }
            }
        }
    }
}
