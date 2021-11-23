package Tree;

import DataStructure.TreeNode;

import java.lang.reflect.Array;
import java.util.*;

public class _32_PrintFromTopToBottom {
    // 层次打印 按照数字
    public static ArrayList<Integer> printFromToBottom(TreeNode root) {
        Queue<TreeNode> queue = new LinkedList<>();
        ArrayList<Integer> ret = new ArrayList<>();
        queue.add(root);
        while (!queue.isEmpty()) {
            int cnt = queue.size();
            while (cnt-- > 0) {
                TreeNode t = queue.poll();
                if (t==null) {
                    continue;
                }
                ret.add(t.val);
                queue.add(t.left);
                queue.add(t.right);
            }
        }
    return ret;
    }

    // 层次打印 按照列表
    public static ArrayList<ArrayList<Integer>> Print(TreeNode root) {
        ArrayList<ArrayList<Integer>> ret = new ArrayList<>();
        Queue<TreeNode> queue = new LinkedList<>();
        queue.add(root);
        while (!queue.isEmpty()) {
            ArrayList<Integer> list = new ArrayList<>();
            int cnt = queue.size();
            while (cnt-- > 0) {
                TreeNode node = queue.poll();
                if (node==null) {
                    continue;
                }

                list.add(node.val);
                queue.add(node.left);
                queue.add(node.right);

            }
            if (list.size()!=0) {
                ret.add( list);
            }
        }
    return ret;
    }

    // 之字打印二叉树
    public static ArrayList<ArrayList<Integer>> print02(TreeNode root) {
        ArrayList<ArrayList<Integer>> ret = new ArrayList<>();

        Queue<TreeNode> queue= new LinkedList<>();
        queue.add(root);
        boolean reverse = false;
        while (!queue.isEmpty()) {
            ArrayList<Integer> list = new ArrayList<>();
            int cnt = queue.size();
            while (cnt-- > 0) {
                TreeNode node = queue.poll();
                if (node==null) {
                    continue;
                }
                list.add(node.val);
                queue.add(node.left);
                queue.add(node.right);
            }
            if (reverse) {
                Collections.reverse(list);
            }
            reverse=!reverse;
            if (list.size()!=0) {
                ret.add(list);
            }
        }
        return ret;
    }

    public static void main(String[] args) {
        TreeNode root = new TreeNode(1);
        root.left=new TreeNode(2);
        root.right=new TreeNode(3);
        root.left.left=new TreeNode(4);
        root.left.right=new TreeNode(5);
        root.right.right=new TreeNode(7);
        root.right.left=new TreeNode(6);
        root.left.left.left=new TreeNode(8);
        root.left.left.right=new TreeNode(9);

//        ArrayList<Integer> res = printFromToBottom(root);
//        Iterator ite = res.iterator();
//        while (ite.hasNext()) {
//            Object next = ite.next();
//            System.out.println(next);
//        }
//
//        ArrayList<ArrayList<Integer>> ret02 = Print(root);
//        Iterator ite = ret02.iterator();
//        while (ite.hasNext()) {
//            Object next = ite.next();
//            System.out.println(next);
//        }

        ArrayList<ArrayList<Integer>> arrayLists = print06Test(root);
        Iterator ite = arrayLists.iterator();
        while (ite.hasNext()) {
            Object next = ite.next();
            System.out.println(next);
        }

    }

    public static ArrayList<Integer> print03TestIntegers(TreeNode root) {
        ArrayList<Integer> res = new ArrayList<>();
        if (root==null) {
            return res;
        }
        Queue<TreeNode> queue = new LinkedList<>();
        queue.add(root);
        while (!queue.isEmpty()) {
            int count = queue.size();
            while (count-- > 0) {
                TreeNode node = queue.poll();
                if (node==null) {
                    continue;
                }
                res.add(node.val);
//                res.add(queue.poll().val);
                queue.add(node.left);
                queue.add(node.right);
            }
        }
        return res;
    }


    public static ArrayList<ArrayList<Integer>> print04Test (TreeNode root) {
        ArrayList<ArrayList<Integer>> res = new ArrayList<>();
        if (root==null) {
            return res;
        }
        Queue<TreeNode> queue = new LinkedList<>();
        queue.add(root);
        while (!queue.isEmpty()) {
            int count = queue.size();
            ArrayList<Integer> list = new ArrayList<>();
            while (count-- > 0) {
                TreeNode node = queue.poll();
                if (node==null) {
                    continue;
                }
                list.add(node.val);
                queue.add(node.left);
                queue.add(node.right);
            }
            if (!list.isEmpty()) {
                res.add(list);
            }
        }

        return res;
    }


    public static ArrayList<ArrayList<Integer>> print05Test (TreeNode root) {
        ArrayList<ArrayList<Integer>> res = new ArrayList<>();
        if (root==null) {
            return res;
        }
        Queue<TreeNode> queue = new LinkedList<>();
        queue.add(root);
        boolean reverse = false;
        while (!queue.isEmpty()) {
            int count = queue.size();

            ArrayList<Integer> list = new ArrayList<>();
            while (count-- > 0) {
                TreeNode node = queue.poll();
                if (node==null) {
                    continue;
                }
                list.add(node.val);
                queue.add(node.left);
                queue.add(node.right);
            }
            if (reverse) {
                Collections.reverse(list);
            }
            reverse=!reverse;
            if (!list.isEmpty()) {

                res.add(list);
            }

        }
        return res;
    }

    public static ArrayList<ArrayList<Integer>> print06Test (TreeNode root ) {
        ArrayList<ArrayList<Integer>> res = new ArrayList<>();
        if (root==null) {
            return res;
        }
        Queue<TreeNode> queues = new LinkedList<>();
        queues.add(root);
        boolean reverseFlag = false;
        while (!queues.isEmpty()) {
            int count = queues.size();
            ArrayList<Integer> list = new ArrayList<>();
            while (count-- > 0) {
                TreeNode node = queues.poll();
                if (node==null) {
                    continue;
                }
                list.add(node.val);
                queues.add(node.left);
                queues.add(node.right);
            }
            if (reverseFlag) {
                Collections.reverse(list);
            }
            reverseFlag=!reverseFlag;
            if (!list.isEmpty()) {
                res.add(list);
            }


        }
        return res;
    }
}
















