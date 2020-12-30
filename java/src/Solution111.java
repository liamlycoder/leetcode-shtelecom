import entity.TreeNode;

/**
 * @author liamcoder
 * @date 2020/12/30 14:58
 */
public class Solution111 {
    public int minDepth(TreeNode root) {
        if (root == null) {
            return 0;
        }
        int minLeftDepth = minDepth(root.left);
        int minRightDepth = minDepth(root.right);
        return root.left == null || root.right == null ? minLeftDepth + minRightDepth + 1 : Math.min(minLeftDepth, minRightDepth) + 1;
    }
}
