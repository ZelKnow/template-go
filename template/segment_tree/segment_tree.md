# 线段树

- [代码](segment_tree.go)
- [动态开点](segment_tree_v2.go)

线段树简单来讲就是把区间变成一颗完全二叉树，根节点管理整个区间，两个孩子分别管理前半区间和后半区间，以此类推，直到区间长度为 1。每个节点存储这个区间的和。要查询某个区间，就找到线段树中对应的几个节点加起来即可。如果要修改某个值，则从根节点开始搜索，遇到的每个节点都进行修改就行。

如果要支持区间修改，则需要引入懒标记。每当寻找到一个节点，该节点对应的区间是修改区间的子集时，则修改这个节点的值，然后在该节点的懒标记中记录一下修改的值，就直接返回。而查询或修改过程中若发现某节点有懒标记，则需要将懒标记下放到子节点。注意修改过程中每个节点都要根据子节点的值更新一下自己。

- 区间修改、区间查询复杂度 `O(logN)`
- 建树复杂度为 `O(2^(logN+1))`，也即 `O(n)`