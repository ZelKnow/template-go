# 堆

- [模板](heap.go)
- [手写堆](heap_handwrite.go)

主要指二叉堆。满足堆顶比两个儿子都大/小，且两个子树也是堆。

需要实现：

- 向上调整
- 向下调整
- 插入：将元素插到尾部，再向上调整，复杂度 `O(logn)`
- 删除：将堆顶元素与最后一个元素交换，再从堆顶向下调整 `O(logn)`
- 建堆（将数组变为堆）：从后往前对每个元素向下调整，复杂度 `O(n)`