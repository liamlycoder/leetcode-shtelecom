import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * @author liamcoder
 * @date 2021/1/12 9:23
 */
public class Solution1203 {
    public int[] sortItems(int n, int m, int[] group, List<List<Integer>> beforeItems) {
        // 得到每个组有哪些项目
        Map<Integer, List<Integer>> map = new HashMap<>();
        for (int i = 0; i < group.length; ++i) {
            // 没有组的项目，即分组为-1的组，我们对他重新分组，从m开始正序
            if (group[i] == -1) {
                group[i] = m++;
            }
            // 先看看i这个项目的组原来负责了哪些项目，如果原来没负责项目，那我就给你创建一个空列表，增加当前项目进去
            List<Integer> lst = map.computeIfAbsent(group[i], k -> new ArrayList<>());
            lst.add(i);
        }



    }
}
