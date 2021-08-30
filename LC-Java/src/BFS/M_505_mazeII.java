package BFS;

import java.awt.image.DataBufferUShort;
import java.util.Arrays;
import java.util.PriorityQueue;

public class M_505_mazeII {
    int[][] dirs = {{0,1},{0,-1},{-1,0},{1,0}};
    public int shortestDistance(int[][] maze, int[] start, int[] destination) {
        int[][] distance = new int[maze.length][maze[0].length];
        for (int[] row : distance) {
            Arrays.fill(row,Integer.MAX_VALUE);
        }
        distance[start[0]][start[1]]=0;
        dijkstra(maze,start,distance);
        return distance[destination[0]][destination[1]] == Integer.MAX_VALUE ? -1 : distance[destination[0]][destination[1]];
    }

    private void dijkstra(int[][] maze, int[] start, int[][] distance) {
        PriorityQueue<int[]> q = new PriorityQueue<>((a, b) -> (a[2] - b[2]));
        q.offer(new int[]{start[0],start[1]});
        while (!q.isEmpty()) {
            int[] cur = q.poll();
            for (int[] dir : dirs) {
                int x = cur[0] + dir[0];
                int y = cur[1] + dir[1];
                int count =0;
                while (x >= 0 && y >= 0 && x < maze.length && y < maze[0].length && maze[x][y] == 0) {
                    x+=dir[0];
                    y+=dir[1];
                    count++;
                }
                x-=dir[0];
                y-=dir[1];
                if (distance[cur[0]][cur[1]]+count< distance[x][y]) {
                    distance[x][y]=distance[cur[0]][cur[1]]+count;
                    q.add(new int[]{x,y,distance[x][y]});
                }
            }
        }
    }
}
