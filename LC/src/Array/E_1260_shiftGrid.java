package Array;

import java.util.Arrays;
import java.util.LinkedList;
import java.util.List;

/*给你一个 m 行 n 列的二维网格 grid 和一个整数 k。你需要将 grid 迁移 k 次。

 */
public class _1260_shiftGrid {
    public List<List<Integer>> shiftGrid(int[][] grid, int k) {
        int row = grid.length, col = grid[0].length;
        int newK = k % (row * col);
        List<List<Integer>> ret = new LinkedList<List<Integer>>();
        int endPoint = newK == 0 ? 0 : row * col - newK;
        int boundry = row * col;
        Integer newGrid[][] = new Integer[row][col];
        for (int i = 0, j = endPoint; i < boundry; i++, j++) {
            if (j == boundry) {
                j = 0;
            }
            newGrid[i / col][i % col] = grid[j / col][j % col];
        }
        for (Integer[] a : newGrid) {
            ret.add(Arrays.asList(a));
        }
        return ret;
    }
}
