import com.sun.source.tree.Tree;
import entity.TreeNode;

import java.util.HashMap;


public class Solution_Sword_07 {
    // 方法一：哈希+递归   （容易理解）
    public TreeNode buildTree(int[] preorder, int[] inorder) {
        HashMap<Integer, Integer> map = new HashMap<>();
        for (int i = 0; i < inorder.length; ++i) {
            map.put(inorder[i], i);  // 键：中序数组的每个元素值，值：中序书序每个元素值的索引
        }
        return buildTreeHelper(preorder, 0, preorder.length, inorder, 0, inorder.length, map);
    }

    private TreeNode buildTreeHelper(int[] preorder, int pStart, int pEnd,
                                     int[] inorder, int iStart, int iEnd, HashMap<Integer, Integer> map) {
        if (pStart == pEnd) {
            return null;
        }
        int rootVal = preorder[pStart];
        TreeNode root = new TreeNode(rootVal);
        int inorderIndex = map.get(rootVal);
        int leftNum = inorderIndex - iStart;
        root.left = buildTreeHelper(preorder, pStart + 1, pStart + leftNum + 1, inorder, iStart, inorderIndex, map);
        root.right = buildTreeHelper(preorder, pStart + leftNum + 1, pEnd, inorder, inorderIndex + 1, iEnd, map);
        return root;
    }





    // 双指针 + 递归 (无法理解)
    // 链接：https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/solution/xiang-xi-tong-su-de-si-lu-fen-xi-duo-jie-fa-by--22/
    private int prePointer = 0;  // 先序遍历的指针
    private int inPointer = 0;  // 中序遍历的指针
    public TreeNode buildTree2(int[] preorder, int[] inorder) {
        return buildChildTree(preorder, inorder, Integer.MAX_VALUE);
    }

    public TreeNode buildChildTree(int[] preorder, int[] inorder, int stop) {
        // 先序指针范围的判定
        if (prePointer >= preorder.length) {
            return null;
        }

        // 中序指针范围的判定
        if (inorder[inPointer] == stop) {
            inPointer++;
            return null;
        }

        TreeNode root = new TreeNode(preorder[prePointer++]);
        root.left = buildChildTree(preorder, inorder, root.val);
        root.right = buildChildTree(preorder, inorder, stop);
        return root;
    }

}
