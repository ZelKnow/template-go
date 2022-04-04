package template

type MqData struct {
	val	int
	del	int // 懒删除标记，可以理解为原数组中该元素和队列中上一个元素间有多少个元素
}

type monotoneQueue struct {
	data	[]MqData
	size	int
}

func (mq monotoneQueue) less(a, b MqData) bool {
	return a.val >= b.val // 维护区间最大值
	//return a.val <= b.val // 维护区间最小值
}

func (mq *monotoneQueue) push(v int) {
	mq.size++
	d := MqData{v, 1}
	for len(mq.data) > 0 && mq.less(d, mq.data[len(mq.data)-1]) {
		d.del += mq.data[len(mq.data)-1].del // 删除前记录下来，用于 pop
		mq.data = mq.data[:len(mq.data)-1]
	}
	mq.data = append(mq.data, d)
}

func (mq *monotoneQueue) pop() {
	mq.size--
	if mq.data[0].del > 1 {
		mq.data[0].del--
	} else {
		mq.data = mq.data[1:]
	}
}

// 返回区间最值
// 调用前需保证 mq.size > 0
func (mq monotoneQueue) top() int {
	return mq.data[0].val
}

func windowMax(a []int, windowSize int) []int {
	n := len(a)
	q := monotoneQueue{}
	ans := make([]int, 0, n-windowSize+1)
	for i, v := range a {
		q.push(v)
		if q.size > windowSize {
			q.pop()
		}
		if i+1 >= windowSize {
			ans = append(ans, q.top())
		}
	}
	return ans
}