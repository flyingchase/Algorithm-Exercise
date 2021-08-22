package DFS;

import java.util.ArrayList;
import java.util.List;

//给定一个无重复元素的正整数数组 candidates 和一个正整数 target ，找出 candidates 中所有可以使数字和为目标数 target 的唯一组合。
//
// candidates 中的数字可以无限制重复被选取。如果至少一个所选数字数量不同，则两种组合是唯一的。
//
// 对于给定的输入，保证和为 target 的唯一组合数少于 150 个。
public class M_39_combinationSum {

    public static void main(String[] args) {
        int[] candidates = {2, 3, 6, 7};
        System.out.println(new M_39_combinationSum().combinationSum(candidates, 7));
    }
    // public List<List<Integer>> combinationSum(int[] candidates, int target) {
    //     List<List<Integer>> res = new ArrayList<>();
    //     dfs(res, new ArrayList<Integer>(), candidates, target, 0);
    //     return res;
    // }
    //
    // private void dfs(List<List<Integer>> res, ArrayList<Integer> level, int[] candidates, int remain, int index) {
    //     if (remain<0) {
    //         return;
    //     } else if (remain==0){
    //         res.add(new ArrayList<>(level));
    //     }else {
    //         for (int i = index; i < candidates.length; i++) {
    //             level.add(candidates[i]);
    //             dfs(res,level,candidates,remain-candidates[i],i);
    //             level.remove(level.size()-1);
    //         }
    //     }
    // }

    public List<List<Integer>> combinationSum(int[] candidates, int target) {
        List<List<Integer>> res = new ArrayList<>();
        dfs(res, new ArrayList<Integer>(), candidates, target, 0);
        return res;
    }

    private void dfs(List<List<Integer>> res, List<Integer> lists, int[] candidates, int remain, int index) {
        if (remain == 0) {
            res.add(new ArrayList<>(lists));
        } else if (remain < 0) {
            return;
        } else {
            for (int i = index; i < candidates.length; i++) {
                lists.add(candidates[i]);
                dfs(res, lists, candidates, remain - candidates[i], i);
                lists.remove(lists.size() - 1);
            }
        }
    }
}
