import java.util.HashMap;
import java.util.Map;

/**
 * @author liamcoder
 * @date 2021/1/15 8:36
 */
public class Solution947 {
    public int removeStones(int[][] stones) {
        UnionFind unionFind = new UnionFind();
        for (int[] stone : stones) {
            unionFind.union(stone[0] + 10000, stone[1]);
        }
        return stones.length - unionFind.getCount();
    }
    private class UnionFind {
        private Map<Integer, Integer> parent;
        private int count;
        private int getCount() {
            return this.count;
        }
        public UnionFind() {
            this.parent = new HashMap<>();
            this.count = 0;
        }

        public int find(int x) {
            if (!parent.containsKey(x)) {
                parent.put(x, x);
                count ++;
            }
            if (x != parent.get(x)) {
                parent.put(x, find(parent.get(x)));
            }
            return parent.get(x);
        }

        public void union(int x, int y) {
            int rootX = find(x);
            int rootY = find(y);
            if (rootX != rootY) {
                parent.put(rootX, rootY);
                --count;
            }
        }
    }
}
