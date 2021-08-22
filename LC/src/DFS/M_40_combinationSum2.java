package DFS;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

//给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
//
// candidates 中的每个数字在每个组合中只能使用一次。
//
// 注意：解集不能包含重复的组合。
public class M_40_combinationSum2 {
    public static void main(String[] args) {
        int[] nums = {10,1,2,7,6,1,5};
        System.out.println(new M_40_combinationSum2().combinationSum2(nums, 8));
    }
    public List<List<Integer>> combinationSum2(int[] candidates, int target) {
        List<List<Integer>> res = new ArrayList<>();
        Arrays.sort(candidates);
        dfs(res, new ArrayList<Integer>(), candidates, target, 0, new boolean[candidates.length]);
        return res;
    }

    private void dfs(List<List<Integer>> res, List<Integer> lists, int[] nums, int remain, int index, boolean[] used) {
        if (remain < 0) {
            return;
        } else if (remain == 0) {
            res.add(new ArrayList<>(lists));
        } else {
            for (int i = index; i < nums.length; i++) {
                if (used[i] || (i>0&&nums[i]==nums[i-1]&&!used[i-1])) {
                    continue;
                }
                lists.add(nums[i]);
                used[i]=true;
                dfs(res,lists,nums,remain-nums[i],i+1,used);
                used[i]=false;
                lists.remove(lists.size()-1);
            }
        }
    }
}
