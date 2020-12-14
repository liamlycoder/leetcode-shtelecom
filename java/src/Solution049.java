import java.util.*;

/**
 * @author liamcoder
 * @date 2020/12/14 14:01
 */
public class Solution049 {
    public List<List<String>> groupAnagrams(String[] strs) {
        List<List<String>> res = new ArrayList<>();
        HashMap<Long, List<String>> map = new HashMap<>();
        int[] prime = new int[] {2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101};
        for (String str : strs) {
            long sum = 1L;
            for (int j = 0; j < str.length(); ++j) {
                int index = (int) str.charAt(j) - 97;
                sum *= prime[index];
            }
            map.putIfAbsent(sum, new ArrayList<>());
            map.get(sum).add(str);
        }

        for (Long key : map.keySet()) {
            res.add(map.get(key));
        }
        return res;
    }
}
