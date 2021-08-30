package SlidingWindow;

import DataStructure.TreeNode;

import java.util.*;

/*给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。

 */
public class _76_minWindow {
    public static void main(String[] args) {
        String s = "ADOBECODEBANC", t = "ABC";
        System.out.println(new _76_minWindow().minWindow(s, t));
        TreeNode root = new TreeNode(1);
        root.left=new TreeNode(2);
        root.right=new TreeNode(3);
        root.left.left=new TreeNode(4);
        root.left.right=new TreeNode(5);
        root.right.right=new TreeNode(7);
        root.right.left=new TreeNode(6);
        root.left.left.left=new TreeNode(8);
        root.left.left.right=new TreeNode(9);

// preOrder: 1 2 4 8 9 5 3 6 7
// inOrder: 8 4 9 2 5 1 6 3 7
// postOrder: 8 9 4 5 2 6 7 3 1

        System.out.println(new _76_minWindow().zigzaLevelTravelsBT(root));
    }

    public String minWindow(String s, String t) {
        if (s == null) {
            return null;
        }
        HashMap<Character, Integer> map = new HashMap<>();
        int left = 0, minStart = 0, minLen = Integer.MAX_VALUE, count = 0;
        for (char c : t.toCharArray()) {
            map.put(c, map.getOrDefault(c, 0) + 1);
        }
        for (int i = 0; i < s.length(); i++) {
            char c = s.charAt(i);
            if (map.containsKey(c)) {
                if (map.get(c) > 0) {
                    count++;
                }
                map.put(c, map.get(c) - 1);
            }
            while (t.length() == count) {
                if (minLen > i - left + 1) {
                    minLen = i - left + 1;
                    minStart = left;
                }
                char leftChar = s.charAt(left);
                if (map.containsKey(leftChar)) {
                    map.put(leftChar, map.get(leftChar) + 1);
                    if (map.get(leftChar) > 0) {
                        count--;
                    }
                }
                left++;
            }
        }
        if (minLen == Integer.MAX_VALUE) {
            return "";
        }
        return s.substring(minStart, minStart + minLen);
    }

    public ArrayList<Integer> preOrder(TreeNode root) {
        if (root == null) {
            return null;
        }
        ArrayList<Integer> res = new ArrayList<>();
        Stack<TreeNode> stack = new Stack<>();
        TreeNode cur = root;
        while (cur != null || !stack.isEmpty()) {
            while (cur != null) {
                res.add(cur.val);
                stack.add(cur);
                cur = cur.left;
            }
            TreeNode node = stack.pop();
            cur = node.right;
        }
        return res;
    }

    public List<Integer> inOrderTraversalBT(TreeNode root) {
        List<Integer> res = new ArrayList<>();
        //空树则直接返回null的res
        if (root == null) {
            return res;
        }
        Stack<TreeNode> stack = new Stack<>();
        // 保存root
        TreeNode cur = root;
        while (cur != null || !stack.isEmpty()) {
            while (cur != null) {
                stack.add(cur);
                cur = cur.left;
            }
            TreeNode node = stack.pop();
            res.add(node.val);
            cur = node.right;

        }

        return res;
    }

    public List<Integer> postOrderTraversalBT(TreeNode root) {
        List<Integer> res = new ArrayList<>();
        //空树则直接返回null的res
        if (root == null) {
            return res;
        }
        Stack<TreeNode> stack = new Stack<>();
        // 保存root
        TreeNode cur = root;
        while (cur != null || !stack.isEmpty()) {
            while (cur != null) {
                res.add(cur.val);
                stack.add(cur);
                cur = cur.right;
            }
            TreeNode node = stack.pop();
            cur = node.left;
        }
        Collections.reverse(res);
        return res;
    }

    public ArrayList<ArrayList<Integer>> zigzaLevelTravelsBT(TreeNode root) {
        ArrayList<ArrayList<Integer>> res = new ArrayList<>();
        if (root == null) {
            return res;
        }
        Queue<TreeNode> queue = new LinkedList<>();
        TreeNode cur = root;
        queue.add(cur);
        boolean reverseFlag = false;
        while (!queue.isEmpty()) {
            ArrayList<Integer> lists = new ArrayList<>();
            int size = queue.size();
            while (size-- > 0) {
                TreeNode node = queue.poll();
                if (node==null) {
                    continue;
                }
                lists.add(node.val);
                queue.add(node.left);
                queue.add(node.right);
            }
            if (reverseFlag){
                Collections.reverse(lists);
            }
            reverseFlag=!reverseFlag;
            if (!lists.isEmpty()) {
                res.add(lists);
            }
        }

        return res;
    }
}
