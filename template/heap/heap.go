package template

import "container/heap"

type heapData struct {
	// 可自定义
	val	int
}

type hp []heapData

func (h hp) Len() int {
	return len(h)
}

func (h hp) Less(i, j int) bool {
	// 需修改比较规则
	return h[i].val < h[j].val // > 为最大堆
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *hp) Push(v interface{}) {
	*h = append(*h, v.(heapData))
}

func (h *hp) Pop() interface{} {
	a := *h;
	v := a[len(a)-1];
	*h = a[:len(a)-1];
	return v
}

func (h *hp) push(v heapData) {
	heap.Push(h, v)
}

func (h *hp) pop() heapData {
	return heap.Pop(h).(heapData)
}

func newHeap() *hp {
	myHeap := new(hp)
	heap.Init(myHeap)
	return myHeap
}