package DFS;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

//给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
public class M_46_permute {
    public List<List<Integer>> permute(int[] nums) {
        List<List<Integer>> res = new ArrayList<>();
        Arrays.sort(nums);
        dfs(res, new ArrayList<Integer>(), nums);
        return res;
    }

    private void dfs(List<List<Integer>> res, List<Integer> list, int[] nums) {
        if (list.size() == nums.length) {
            res.add(new ArrayList<>(list));
        } else {
            for (int i = 0; i < nums.length; i++) {
                if (list.contains(nums[i])) {
                    continue;
                }
                list.add(nums[i]);
                dfs(res,list,nums);
                list.remove(list.size()-1);
            }
        }
    }
}
