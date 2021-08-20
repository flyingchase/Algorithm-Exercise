# LC 重点





LC 881



## 114 二叉树转换为链表

### 题意

将一棵二叉树按照前序遍历拆解成为一个 `假链表`。所谓的假链表是说，用二叉树的 *right* 指针，来表示链表中的 *next* 指针。

### codes

``` java
public void flatten(TreeNode root) {
    if (root==null) {
        return;
    }
    while (root != null) {
        if (root.left!=null) {
            TreeNode right = root.right;
            root.right=root.left;
            root.left=null;
            TreeNode node = root.right;
            while (node != null && node.right != null) {
                node=node.right;
            }
            node.right=right;

        }
        root=root.right;
    }
}
```





### 注意事项

- 非递归：
  - 不断将左子树拆下来接到右边去。然后将左子树与右子树连接。