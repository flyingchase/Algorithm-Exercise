package Tree;

import DataStructure.TreeNode;

import java.util.*;

@SuppressWarnings({"all"})
public abstract class DailyExercise0705 {
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

        int[] nums = {1, 0, 9, 2, 6, 3, 8, 4, 7, 5, 16, 10, 14, 12, 13, 15};
        ArrayList<ArrayList<Integer>> arrayLists = zigziPrintBtziginal(root);
        System.out.println(arrayLists);
        int[] nums1 = Arrays.copyOf(nums, nums.length);
        heapSort(nums1);
        PrimitiveIterator.OfInt iterator = Arrays.stream(nums1).iterator();
        while (iterator.hasNext()) {
            Integer next = iterator.next();
            System.out.print(next + " ");
        }
        System.out.println();


        int[] nums2 = Arrays.copyOf(nums, nums.length);
        quickSort(nums2);
        System.out.println(Arrays.toString(nums2));


        int[] nums3 = Arrays.copyOf(nums, nums.length);
        BubbleSort(nums3);
        System.out.println(Arrays.toString(nums3));


    }

    public static ArrayList<Integer> preOrderBT(TreeNode root) {
        ArrayList<Integer> res = new ArrayList<>();
        if (root == null) {
            return res;
        }

        Stack<TreeNode> stack = new Stack<>();
        TreeNode cur = root;
        while (cur != null || !stack.isEmpty()) {
            while (cur != null) {
                res.add(cur.val);
                stack.push(cur);
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
            cur = node.left;


        }
        Collections.reverse(res);
        return res;
    }


    public static ArrayList<ArrayList<Integer>> zigziPrintBtziginal(TreeNode root) {
        ArrayList<ArrayList<Integer>> res = new ArrayList<>();

        if (root == null) {
            return res;
        }
        boolean reverseFalg = false;
        Queue<TreeNode> queue = new LinkedList<>();
        queue.add(root);
        TreeNode cur = root;
        while (!queue.isEmpty()) {
            ArrayList<Integer> list = new ArrayList<>();

            int size = queue.size();
            while (size-- > 0) {
                TreeNode node = queue.poll();
                if (node == null) {
                    continue;
                }
                list.add(node.val);
                queue.add(node.left);
                queue.add(node.right);
            }
            if (reverseFalg) {
                Collections.reverse(list);
            }
            reverseFalg = !reverseFalg;
            if (!list.isEmpty()) {
                res.add(list);
            }

        }

        return res;
    }

    public static void heapInsert(int[] nums, int index) {
        while (nums[index] > nums[(index - 1) / 2]) {
            swap(nums, index, (index - 1) / 2);
            index = (index - 1) / 2;
        }

    }

    public static void heapify(int[] nums, int index, int size) {
        int left = index * 2 + 1, right = index * 2 + 2;

        while (left < size) {
            int largest = right < size && nums[left] < nums[right] ? right : left;
            largest = nums[index] > nums[largest] ? index : largest;
            if (largest == index) {
                break;
            }
            swap(nums, index, largest);
            index = largest;
            left = 2 * index + 1;
            right = 2 * index + 2;
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
            swap(nums, 0, --size);
            heapify(nums, 0, size);
        }
    }

    public static void swap(int[] nums, int i, int j) {
        nums[i] = nums[i] ^ nums[j];
        nums[j] = nums[i] ^ nums[j];
        nums[i] = nums[i] ^ nums[j];
    }


    public static void mergeSort(int[] nums) {
        if (nums == null) {
            return;
        }
        mergeSort(nums, 0, nums.length - 1);
    }

    public static void mergeSort(int[] nums, int l, int r) {

        if (l == r) return;
        int mid = ((r - l) >> 1) + l;
        mergeSort(nums, l, mid);
        mergeSort(nums, mid + 1, r);
        merge(nums, l, mid, r);
    }

    public static void merge(int[] nums, int l, int mid, int r) {

        int p1 = l, p2 = mid + 1, i = 0;
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
        for (i = 0; i < nums.length; i++) {
            nums[l + i] = helper[i];
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
            swap(nums, l + (int) (Math.random() * (r - l + 1)), r);
            int[] p = partition(nums, l, r);
            quickSort(nums, l, p[0] - 1);
            quickSort(nums, p[1] + 1, r);
        }
    }

    public static int[] partition(int[] nums, int l, int r) {
        int less = l - 1;
        int more = r;
        while (l < more) {
            if (nums[l] < nums[r]) {
                swap(nums, ++less, l++);

            } else if (nums[l] > nums[r]) {
                swap(nums, --more, l);
            } else {
                l++;
            }
        }
        swap(nums, r, more);
        return new int[]{less + 1, more};
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


}
