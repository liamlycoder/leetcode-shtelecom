import java.util.HashMap;
import java.util.Map;

public class Solution001 {
    public int[] twoSum(int[] nums, int target) {
        int[] indices = new int[2];
        Map<Integer, Integer> map = new HashMap<Integer, Integer>();
        for (int i = 0; i < nums.length; ++i) {
            if (map.containsKey(nums[i])) {
                indices[0] = i;
                indices[1] = map.get(nums[i]);
                return indices;
            }

            map.put(target-nums[i], i);
        }

        return indices;
    }
}
