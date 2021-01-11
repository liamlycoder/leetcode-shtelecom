import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.List;

/**
 * @author liamcoder
 * @date 2021/1/11 8:40
 */
public class Solution1202 {
    public String smallestStringWithSwaps(String s, List<List<Integer>> pairs) {
        int len = s.length();

        DSU dsu = new DSU(100000);

        for (List<Integer> pair : pairs) {
            dsu.union(pair.get(0), pair.get(1));
        }

        HashMap<Integer, List<Integer>> map = new HashMap<>();
        for (int i = 0; i < len; ++i) {
            int key = dsu.find(i);
            map.computeIfAbsent(key, unused -> new ArrayList<>()).add(i);
        }

        StringBuffer res = new StringBuffer(s);
        for (List<Integer> value : map.values()) {
            if (value.size() > 1) {
                sort(res, value);
            }
        }
        return res.toString();
    }

    private void sort(StringBuffer res, List<Integer> idx_list) {
        int len = idx_list.size();
        char[] array = new char[len];
        //根据下标集合构建字符集合
        for (int i = 0; i < len; ++i) {
            array[i] = res.charAt(idx_list.get(i));
        }
        //将字符集合排序
        Arrays.sort(array);
        //将字符按序“插入”回StringBuilder
        for (int i = 0; i < len; ++i) {
            res.setCharAt(idx_list.get(i), array[i]);
        }
    }
}

class DSU {
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
