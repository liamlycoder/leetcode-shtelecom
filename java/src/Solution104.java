import entity.TreeNode;

/**
 * @author liamcoder
 * @date 2020/12/30 14:13
 */
public class Solution104 {
    public int maxDepth(TreeNode root) {
        if (root != null) {
            int leftDepth = maxDepth(root.left);
            int rightDepth = maxDepth(root.right);
            return Math.max(leftDepth, rightDepth) + 1;
        }
        return 0;
    }
}
