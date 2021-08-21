package BFS;

import java.util.LinkedList;
import java.util.Queue;

//由空地（用 0 表示）和墙（用 1 表示）组成的迷宫 maze 中有一个球。球可以途经空地向 上、下、左、右 四个方向滚动，且在遇到墙壁前不会停止滚动。当球停下时，可以选择向下一个方向滚动。
// 给你一个大小为 m x n 的迷宫 maze ，以及球的初始位置 start 和目的地 destination ，其中 start = [startrow, startcol] 且 destination = [destinationrow, destinationcol] 。请你判断球能否在目的地停下：如果可以，返回 true ；否则，返回 false 。
//
// 你可以 假定迷宫的边缘都是墙壁（参考示例）。
public class M_490_MazehasPath {
    int[][] dirs={{0,1},{0,-1},{-1,0},{1,0}};
    public boolean hasPath(int[][] maze, int[] start, int[] destination) {
        boolean[][] visited = new boolean[maze.length][maze[0].length];
        Queue<int[]> q = new LinkedList<>();
        q.add(start);
        visited[start[0]][start[1]]=true;
        while (!q.isEmpty()) {
            int[] cur = q.poll();
            if (cur[0]==destination[0]&&cur[1]==destination[1]) {
                return true;
            }
            for (int[] dir : dirs) {
                int x = cur[0] + dir[0];
                int y = cur[1] + dir[1];
                while (x >= 0 && y >= 0 && x < maze.length && y < maze[0].length && maze[x][y] == 0) {
                    x+=dir[0];
                    y+=dir[1];
                }
                x-=dir[0];
                y-=dir[1];
                if (!visited[x][y]) {
                    q.add(new int[]{x,y});
                    visited[x][y]=true;
                }
            }
        }
        return false;
    }
}
