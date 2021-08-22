package DFS;

import java.util.ArrayList;
import java.util.List;

//给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
//
// 你可以按 任何顺序 返回答案。
public class M_77_combine {
    public List<List<Integer>> combine(int n, int k) {
        List<List<Integer>> res = new ArrayList<>();
        dfs(res, new ArrayList<Integer>(), n, k, 1);
        return res;
    }

    private void dfs(List<List<Integer>> res, List<Integer> lists, int n, int k, int index) {
        if (lists.size() == k) {
            res.add(new ArrayList<>(lists));
        } else {
            for (int i = index; i <= n; i++) {
                lists.add(i);
                dfs(res,lists,n,k,i+1);
                lists.remove(lists.size()-1);
            }
        }
    }
}
