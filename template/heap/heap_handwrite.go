package template

type heapData struct {
	// 可自定义
	val	int
}

type hp []heapData

func (h hp) Len() int {
	return len(h)
}

// 向下调整
func (h hp) adjustDown(i int) {
	for i * 2 + 1 < h.Len() {
		child := i * 2 + 1
		if child+1 < h.Len() && h.Less(child+1, child) {
			child++
		}
		if h.Less(i, child) {
			break
		}
		h.Swap(i, child)
		i = child
	}
}

// 向上调整
func (h hp) adjustUp(i int) {
	for ; i > 0 && h.Less(i, (i-1)/2); i=(i-1)/2 {
		h.Swap(i, (i-1)/2)
	}
}

func (h hp) Less(i, j int) bool {
	// 需修改比较规则
	return h[i].val < h[j].val // > 为最大堆
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *hp) push(v heapData) {
	*h = append(*h, v)
	h.adjustUp(h.Len()-1)
}

func (h *hp) pop() heapData {
	h.Swap(0, h.Len()-1)
	ret := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	h.adjustDown(0)
	return ret
}

func newHeap() hp {
	return hp{}
}

func buildHeap(array []heapData) hp {
	h := make(hp, len(array))
	copy(h, array)
	for i:=h.Len()/2-1;i>=0;i-- {
		h.adjustDown(i)
	}
	return h
}