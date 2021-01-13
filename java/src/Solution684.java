import entity.DSU;

/**
 * @author liamcoder
 * @date 2021/1/13 8:43
 */
public class Solution684 {
    public int[] findRedundantConnection(int[][] edges) {
        int n = edges.length;
        DSU dsu = new DSU(n + 1);

        for (int i = 0; i < n; ++i) {
            int[] edge = edges[i];
            int node1 = edge[0], node2 = edge[1];
            if (dsu.find(node1) != dsu.find(node2)) {
                dsu.union(node1, node2);
            } else {
                return edge;
            }
        }
        return new int[0];
    }
}


