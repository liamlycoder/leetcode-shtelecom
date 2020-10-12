import java.util.LinkedList;
import java.util.Queue;

public class Solution_Sword_13 {
    public int movingCount(int m, int n, int k) {
        if (k == 0) {
            return 1;
        }
        Queue<int[]> queue = new LinkedList<>();
        queue.add(new int[]{0, 0});
        boolean[][] visited = new boolean[m][n];
        visited[0][0] = true;
        int count = 1;
        int[] dx = new int[]{0, 1};
        int[] dy = new int[]{1, 0};
        while (!queue.isEmpty()) {
            int[] cell = queue.poll();
            int x = cell[0], y = cell[1];
            for (int i = 0; i < 2; ++i) {
                int nx = dx[i] + x;
                int ny = dy[i] + y;
                if (nx >= 0 && nx < m && ny >= 0 && ny < n && !visited[nx][ny] && digitalSum(nx) + digitalSum(ny) <= k) {
                    ++count;
                    visited[nx][ny] = true;
                    queue.offer(new int[]{nx, ny});
                }
            }

        }
        return count;
    }

    public static int digitalSum(int number) {
        int result = 0;
        while (number != 0) {
            result += number % 10;
            number /= 10;
        }
        return result;
    }
}
