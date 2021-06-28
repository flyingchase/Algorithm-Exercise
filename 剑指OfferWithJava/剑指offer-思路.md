# 剑指offer-思路

## 41.2 字符流中第一个非重复字符

### 题目描述

请实现一个函数用来找出字符流中第一个只出现一次的字符。例如，当从字符流中只读出前两个字符 "go" 时，第一个只出现一次的字符是 "g"。当从该字符流中读出前六个字符“google" 时，第一个只出现一次的字符是 "l"。

### 解题思路

使用统计数组来统计每个字符出现的次数，本题涉及到的字符为都为 ASCII 码，因此使用一个大小为 128 的整型数组就能完成次数统计任务。

使用队列来存储到达的字符，并在每次有新的字符从字符流到达时移除队列头部那些出现次数不再是一次的元素。因为队列是先进先出顺序，因此队列头部的元素为第一次只出现一次的字符。

### 注意事项⚠️

将字符流逐个入队并判断队首字符出现次数是否超过一次 是则剔除

```java
private int[]cnts=new int[128];
private Queue<Character> queue=new LinkedList>();

public void Insert(char ch){
        cnts[ch]++;
        queue.add(ch);
        while(!queue.isEmpty()&&cnts[queue.peek()]>1)
        queue.poll();
        }

public char FirstAppearingOnce(){
        return queue.isEmpty()?'#':queue.peek();
        }
```

## 59. 滑动窗口的最大值

### 题目描述

给定一个数组和滑动窗口的大小，找出所有滑动窗口里数值的最大值。

### 解题思路

- 维护一个大小为窗口大小的大顶堆，顶堆元素则为当前窗口的最大值。
- 使用一个双端队列，保证队首存放的是窗口最大值的下标。遍历数组，
    - 队尾元素比要入队的元素小，则把其移除（因为不可能成为窗口最大值）。
    - 队首下标对应的元素不在窗口内（即窗口最大值），将其从队列中移除。
    - 把每次滑动值的下标加入队列中（经过步骤1、2，此时加入队列的下标要么是当前窗口最大值的下标，要么是小于窗口最大值的下标）。
    - 滑动窗口的首地址i大于size就写入窗口最大值。

### 注意事项⚠️

- 大根堆 O(NlogM) n为数组长 m为窗口长度

    - 先将第一个窗口入堆并存下最大值堆顶;
    - 再从`i=0,j=i+size;j<nums.length;i++,j++` 逐个滑动清除出窗数据再读入入窗数据并添加堆顶元素

    ```java
        public static ArrayList<Integer> maxInWidows(int[] nums, int size) {
            ArrayList<Integer> ret = new ArrayList<>();
            if(size>nums.length||size<1) {
                return ret;
            }
            PriorityQueue<Integer> heap = new PriorityQueue<>(Comparator.reverseOrder());
            // 将第一个窗口入堆
            for (int i = 0; i < size; i++) {
                heap.add(nums[i]);
            }
    
            // 存下第一个窗口的堆内的最大值
            ret.add(heap.peek());
            for(int i=0,j=i+size;j<nums.length;i++,j++) {
                // 删除窗口滑动中出窗的数
                heap.remove(nums[i]);
                heap.add(nums[j]);
                ret.add(heap.peek());
            }
            return ret;
        }
    ```

- 双端队列

    ```java
    // 双端队列
        public static ArrayList<Integer> maxInWindows02(int[] nums, int size) {
            ArrayList<Integer> reList = new ArrayList<>();
            if(nums==null||nums.length<size||size<1) {
                return reList;
            }
            Deque<Integer> deque = new LinkedList<>();
            for (int i = 0; i < nums.length; i++) {
                // 队尾元素小于入队元素 则移除
                while (!deque.isEmpty()&&nums[deque.getLast()]<=nums[i]) {
                    deque.pollLast();
                }
                // 队首下标对应的元素不在窗口内即剔除
                while (!deque.isEmpty()&&(i-deque.getFirst()+1>size)) {
                    deque.pollFirst();
                }
                // 每次滑动的值加入队列
                deque.add(i);
                // 滑动窗口的首地址i大于size就写入窗口
                if(!deque.isEmpty()&&i+1>=size) {
                    reList.add(nums[deque.getFirst()]);
                }
            }
            return reList;
            
        }
    ```

    