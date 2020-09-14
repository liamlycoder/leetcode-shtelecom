import entity.TreeNode;

import java.util.ArrayList;
import java.util.List;
import java.util.Stack;

public class Solution094 {
    public List<Integer> inorderTraversal(TreeNode root) {
        Stack<TreeNode> treeNodes = new Stack<>();
        ArrayList<Integer> result = new ArrayList<>();
        TreeNode current = root;
        while (!treeNodes.isEmpty() || current != null) {
            while (current != null) {
                treeNodes.push(current);
                current = current.left;
            }
            current = treeNodes.pop();
            result.add(current.val);
            current = current.right;
        }
        return result;
    }
}
