package 剑指offer_Java;

import java.util.ArrayList;
import java.util.Stack;

import org.graalvm.compiler.nodes.calc.IntegerTestNode;

public class QueueWithTwoStacks {
    Stack<Integer> stack1 = new Stack<Integer>();
    Stack<Integer> stack2 = new Stack<Integer>();

    public void push(int node){
        stack1.push(node);
    }
    public int pop(){
        if(stack2.isEmpty()){
            if(stack1.isEmpty()) return -1;
            while (!stack1.isEmpty()) {
                stack2.push(stack1.pop());
            }
        }
        return stack2.pop();
    }

    public static void main(String[] args) {
        QueueWithTwoStacks queue = new QueueWithTwoStacks();
        ArrayList<Integer> arr1= new ArrayList<Integer>();
        ArrayList<Integer> arr2= new ArrayList<Integer>();
        Integer[] a={1,2,3,4,5,6};
        Integer[] b ={7,8,9,0};
        // Stack s1=new QueueWithTwoStacks().stack1
        Stack<Integer> s1 = new Stack<Integer>();
        Stack<Integer> s2 = new Stack<Integer>();
        for(Integer i : a) s1.push(i);
        for(Integer j : b) s2.push(j);

        Stack<Integer> s3= new QueueWithTwoStacks().new Stack<Integer>();
        Stack<Integer> s4= new QueueWithTwoStacks().new Stack<Integer>();

        s3=s1;
        s4=s2;
        
    }
}
