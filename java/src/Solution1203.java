import java.util.*;

/**
 * @author liamcoder
 * @date 2021/1/12 9:23
 */
public class Solution1203 {
    public int[] sortItems(int n, int m, int[] group, List<List<Integer>> beforeItems) {
        List<List<Integer>> groupItem = new ArrayList<>();//项目分组
        for (int i = 0; i < n + m; i++) {//初始化小组
            groupItem.add(new ArrayList<>());
        }
        int gId = m;//新的组号从m开始
        for (int i = 0; i < group.length; i++) {
            if (group[i] == -1) group[i] = gId++;//没有id的加上组id
            groupItem.get(group[i]).add(i);//同一组的放在一起
        }
        List<List<Integer>> graphInGroup = new ArrayList<>();//组内拓扑关系
        List<List<Integer>> graphOutGroup = new ArrayList<>();//组间拓扑关系
        for (int i = 0; i < n + m; i++) {//初始化拓扑关系
            graphOutGroup.add(new ArrayList<>());
            if (i >= n) continue;
            graphInGroup.add(new ArrayList<>());
        }
        List<Integer> groupId = new ArrayList<>();//所有组id
        for (int i = 0; i < n + m; i++) {
            groupId.add(i);
        }
        // 需要拓扑排序 所以结点的入度必不可少 两个数组分别维护不同结点的入度
        int[] degInGroup = new int[n];//组内 结点入度 （组内项目入度）
        int[] degOutGroup = new int[n + m];//组间 结点入度（小组入度）

        for (int i = 0; i < beforeItems.size(); i++) {//遍历关系
            int curGroupId = group[i];//当前项目i所属的小组id
            List<Integer> beforeItem = beforeItems.get(i);
            for (Integer item : beforeItem) {
                if (group[item] == curGroupId) {//同一组 修改组内拓扑
                    degInGroup[i]++;// 组内结点的入度+1
                    graphInGroup.get(item).add(i);//item 在 i之前
                } else {
                    degOutGroup[curGroupId]++;// 小组间的结点入度 + 1
                    graphOutGroup.get(group[item]).add(curGroupId);// group[item] 小组 在 curGroupId 之前
                }
            }
        }
        //组间拓扑排序，也就是小组之间的拓扑排序，需要的参数 小组结点的入度degOutGroup，所有的小组groupId，组间的拓扑关系图graphOutGroup
        List<Integer> outGroupTopSort = topSort(degOutGroup, groupId, graphOutGroup);
        if (outGroupTopSort.size() == 0) return new int[0];//无法拓扑排序 返回

        int[] res = new int[n];
        int index = 0;
        for (Integer gid : outGroupTopSort) {//遍历排序后的小组id
            List<Integer> items = groupItem.get(gid);//根据小组id 拿到这一小组中的所有成员
            if (items.size() == 0) continue;
            //组内拓扑排序，需要的参数 组内结点的入度degInGroup，组内的所有的结点groupItem.get(gid)，组内的拓扑关系图graphInGroup
            List<Integer> inGourpTopSort = topSort(degInGroup, groupItem.get(gid), graphInGroup);
            if (inGourpTopSort.size() == 0) return new int[0];//无法拓扑排序 返回
            for (int item : inGourpTopSort) {//排序后，依次的放入答案集合当中
                res[index++] = item;
            }
        }
        return res;
    }

    public List<Integer> topSort(int[] deg, List<Integer> items, List<List<Integer>> graph) {
        Queue<Integer> queue = new LinkedList<>();
        for (Integer item : items) {
            if (deg[item] == 0) {
                queue.offer(item);
            }
        }
        List<Integer> res = new ArrayList<>();
        while (!queue.isEmpty()) {
            int cur = queue.poll();
            res.add(cur);
            for (int neighbor : graph.get(cur)) {
                if (--deg[neighbor] == 0) {
                    queue.offer(neighbor);
                }
            }
        }
        return res.size() == items.size() ? res : new ArrayList<>();
    }
}
