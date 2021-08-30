package DFS;

import java.util.ArrayList;
import java.util.List;

//给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//
//
public class E_53_maxSubArray {

    int max = Integer.MIN_VALUE;

    public static void main(String[] args) {
        int[] nums = {-2, 1, -3, 4, -1, 2, 1, -5, 4};
        System.out.println(new E_53_maxSubArray().maxSubArray(nums));
    }

    public int maxSubArray(int[] nums) {
        ArrayList<List<Integer>> res = new ArrayList<>();
        // Arrays.sort(nums);
        dfs(res, new ArrayList<Integer>(), nums, 0);
        int max=0;
        for (List<Integer> list : res) {
            int sum=0;
            for (Integer i : list) {
                sum+=i;
            }
            max=Math.max(max, sum);
        }

        return max;
    }

    private void dfs(ArrayList<List<Integer>> res, ArrayList<Integer> lists, int[] nums, int start) {

        res.add(new ArrayList<Integer>(lists));
        for (int i = start; i < nums.length; i++) {
            // if (i>start&&lists.indexOf(lists.size()-1)!=nums[i-1]) {
            //     continue;
            // }
            lists.add(nums[i]);
            dfs(res, lists, nums, i + 1);
            lists.remove(lists.size() - 1);
        }
    }
}
