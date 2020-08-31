import java.util.List;

public class Solution841 {
    boolean[] visited;
    public boolean canVisitAllRooms(List<List<Integer>> rooms) {
        visited = new boolean[rooms.size()];
        visited[0] = true;
        dfs(rooms, 0);

        // 判断
        for (boolean v : visited) {
            if (!v) {
                return false;
            }
        }
        return true;
    }

    public void dfs(List<List<Integer>> rooms, int key) {
        List<Integer> key_list = rooms.get(key);
        for (int k : key_list) {
            if (!visited[k]) {
                visited[k] = true;
                dfs(rooms, k);
            }
        }
    }

}
