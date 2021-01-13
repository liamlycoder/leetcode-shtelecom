package entity;

/**
 * @author liamcoder
 * @date 2021/1/13 8:52
 */
public class DSU {
    int[] parent;

    public DSU(int len) {
        parent = new int[len];
        for (int i = 0; i < len; ++i) {
            parent[i] = i;
        }
    }
    public int find(int x) {
        return parent[x] != x ? parent[x] = find(parent[x]) : x;
    }

    public void union(int x, int y) {
        parent[find(x)] = find(y);
    }
}
