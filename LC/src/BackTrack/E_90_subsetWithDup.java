package BackTrack;

import java.util.*;

/*给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。

解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。

*/
public class E_90_subsetWithDup {
/*    public List<List<Integer>> subsetWithDup(int[] nums) {
        Arrays.sort(nums);
        Set<List<Integer>> ans = new HashSet<>();
        List<Integer> cur = new ArrayList<>();
        dfs(nums, 0, cur, ans);
        return new ArrayList<List<Integer>>(ans);
    }*/

    public static void main(String[] args) {
        int[] nums = {11, 0, 9, 2, 6, 3, 8, 4, 7, 5, 16, 1, 14, 12, 13, 15};

        System.out.println(Arrays.toString(new E_90_subsetWithDup().quickSort(nums, 0, nums.length - 1)));
    }
    /*
    private void dfs(int[] nums, int u, List<Integer> cur, Set<List<Integer>> ans) {
        if (nums.length == u) {
            ans.add(new ArrayList<>(cur));
            return;
        }
        cur.add(nums[u]);
        dfs(nums, u + 1, cur, ans);

        cur.remove(cur.size() - 1);
        dfs(nums, u + 1, cur, ans);


    }
    */



    public int[] quickSort(int[] nums ,int l, int r) {
        if (l<r) {
            swap(nums, l, r);
            int mid = l + ((r - l) >> 1);
            int[] p = paratition(nums, l, r);
            quickSort(nums,p[1]+1,r);
            quickSort(nums,l,p[0]-1);
        }

        return nums;
    }

    public int[] paratition (int[] nums, int l, int r) {
        int less= l-1,more=r;
        while (l<more) {
            if (nums[l]<nums[r]) {
                swap(nums, ++less,l++);
            }else if (nums[l]> nums[r]){
                swap(nums,--more,l);
            }else {
                l++;
            }
        }
        swap(nums,more,r);
        return new int[]{less + 1, more};
    }

    private void swap(int[] nums, int i, int j) {
        int temp = nums[i];
        nums[i]=nums[j];
        nums[j] = temp;
    }
}

