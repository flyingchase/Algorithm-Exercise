package DFS;

import java.util.ArrayList;
import java.util.List;
//找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字。
//
// 说明：
//
// 所有数字都是正整数。
// 解集不能包含重复的组合。
public class M_216_combinationSum3 {

    public static void main(String[] args) {
        System.out.println(new M_216_combinationSum3().combinationSum3(3, 9));
    }
    int[] nums = {1,2,3,4,5,6,7,8,9};
    public List<List<Integer>> combinationSum3(int k, int n) {
        if (k<=0||n<=0||n>55) {
            return null;
        }
        List<List<Integer>> res = new ArrayList<>();
        dfs(res,new ArrayList<Integer>(),nums,k,0,n);
        return res;
    }

    private void dfs(List<List<Integer>> res, List<Integer> lists, int[] nums, int k, int startIndex, int n) {
        if (n<0) {
            return;
        }else if (n==0&&lists.size()==k){
            res.add(new ArrayList<>(lists));
            return ;
        }else {
            for (int i = startIndex; i < nums.length; i++) {
                lists.add(nums[i]);
                dfs(res,lists,nums,k,i+1,n-nums[i]);
                lists.remove(lists.size() - 1);
            }
        }
    }
}
