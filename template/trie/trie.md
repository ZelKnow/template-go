# 字典树

[代码](trie.go)

即用边来代表字母的树，从根节点到树中某一结点的路径就代表一个字符串，如果要记录树中存储哪些字符串，需要在末尾节点打上标记。

用于快速查找字符串，以及判断前缀。

## 题目

- [LC208](https://leetcode-cn.com/problems/implement-trie-prefix-tree/)：练手题