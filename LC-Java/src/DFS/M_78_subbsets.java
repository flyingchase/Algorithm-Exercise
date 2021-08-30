package DFS;
//给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
//
// 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。

import java.util.ArrayList;
import java.util.List;

public class M_78_subbsets {
    // public List<List<Integer>> subsets(int[] nums) {
    //     List<List<Integer>> res = new ArrayList<>();
    //     backtrack(res,new ArrayList<Integer>(),nums,0);
    //     return res;
    // }
    //
    // private void backtrack(List<List<Integer>> res, ArrayList<Integer> temp, int[] nums, int start) {
    //     res.add(new ArrayList<>(temp));
    //     for (int i = 0; i < nums.length; i++) {
    //         temp.add(nums[i]);
    //         backtrack(res,temp,nums,i+1);
    //         temp.remove(temp.size()-1);
    //     }
    // }
    public List<List<Integer>> subsets(int[] nums) {

        int totalNumber = 1 << nums.length;
        List<List<Integer>> res = new ArrayList<>();
        for (int mask = 0; mask < totalNumber; mask++) {
            List<Integer> set = new ArrayList<>();
            for (int i = 0; i < nums.length; i++) {
                if ((mask & (1 << i)) != 0) {
                    set.add(nums[i]);
                }
            }
            res.add(set);
        }
        return res;
    }
}
