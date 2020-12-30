import entity.TreeNode;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.stream.Collectors;

/**
 * @author liamcoder
 * @date 2020/12/30 15:54
 */
public class Solution654 {
    public TreeNode constructMaximumBinaryTree(int[] nums) {
        return createMaxTree(nums, 0, nums.length - 1);
    }

    private TreeNode createMaxTree(int[] nums, int l, int r) {
        if (l > r) {
            return null;
        }
        int maxIndex = getMaxIndex(nums, l, r);
        TreeNode root = new TreeNode(nums[maxIndex]);
        root.left = createMaxTree(nums, l, maxIndex - 1);
        root.right = createMaxTree(nums, maxIndex + 1, r);
        return root;
    }

    private int getMaxIndex(int[] nums, int l, int r) {
        int maxIndex = -1;
        int maxNum = Integer.MIN_VALUE;
        for (int i = l; i <= r; ++i) {
            if (nums[i] > maxNum) {
                maxIndex = i;
                maxNum = nums[i];
            }
        }
        return maxIndex;
    }
}
