package Sort;

import org.graalvm.compiler.nodes.calc.ObjectEqualsNode;

import Sort.*;

//sort的抽象方法(空方法)——>使得类必须为抽象类
//比较元素
//交换元素
//输出数组内容
public abstract class AbstractSort {
    
    //排序
    public abstract void sort (Object[] nums);
    
    //比较大小 v<w则返回true
    public boolean compare(Object v,Object w) {
        return ((Comparable)v).compareTo(w) < 0 ;
    }

    //交换元素 输入数组和要交换的元素的下标
    //Object是父类 故可以直接类型转换
    public void swap(Object[] nums, int i, int j) {
        Object temp = nums[i];
        nums[i] = nums[j];
        nums[j] = temp;
    }
    
    //输出元素
    public void showArray(Object[] nums) {
        for(Object num : nums) System.out.print(num+" ");
        System.out.println();
    }
    
}