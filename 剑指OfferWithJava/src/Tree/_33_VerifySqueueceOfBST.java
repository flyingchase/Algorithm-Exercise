package Tree;

public class _33_VerifySqueueceOfBST {
    // 二叉搜索树的后序遍历序列判断
    public static boolean VerifySqueueOfBST(int[] seques) {
        if (seques==null||seques.length==0) {
            return false;
        }
        return  verify(seques,0,seques.length-1);
    }

    public static boolean verify (int[] seques, int first, int last) {
        if (last-first<=1) {
            return true;
        }
        int rootVal = seques[last];
        int curIndex = first;
        while (curIndex<last&&seques[curIndex] <=rootVal) {
            curIndex++;
        }
        for (int i=curIndex;i<last;i++) {
            if (seques[i]<rootVal) {
                return false;
            }
        }
        return  verify(seques,first,curIndex-1) && verify(seques,curIndex,last-1);
    }

}
