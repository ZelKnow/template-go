package template

type lazyST []struct {
	l, r int
	todo int64
	lazy bool // 支持区间修改为某一个值
	sum  int64 // 根据需要修改
}

// 线段树定义的操作，可修改
func (lazyST) op(a, b int64) int64 {
	return a + b // % mod
}

// 根据子节点修改节点
func (t lazyST) maintain(o int) {
	lo, ro := t[o<<1], t[o<<1|1]
	t[o].sum = t.op(lo.sum, ro.sum)
}

// 根据数组构建线段树，下标从 1 开始
func (t lazyST) build(a []int64, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].sum = a[l-1]
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

// 对节点 o 表示的区间的所有元素加上 add
// 如果要实现区间修改为某一个值而不是加上某一个值的话，将 += 变为 = 即可
func (t lazyST) do(o int, add int64) {
	to := &t[o]
	to.todo += add                     // % mod
	to.lazy = true
	to.sum += int64(to.r-to.l+1) * add // % mod
}

// 懒标记下放
func (t lazyST) spread(o int) {
	if t[o].lazy {
		add := t[o].todo
		t.do(o<<1, add)
		t.do(o<<1|1, add)
		t[o].todo = 0
		t[o].lazy = false
	}
}

// 对 o 表示的区间中属于 [l, r] 的元素进行修改
func (t lazyST) update(o, l, r int, add int64) {
	if l <= t[o].l && t[o].r <= r {
		// 是 [l, r] 的子集，全部修改
		t.do(o, add)
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.update(o<<1, l, r, add)
	}
	if m < r {
		t.update(o<<1|1, l, r, add)
	}
	t.maintain(o)
}

// o=1  [l,r] 1<=l<=r<=n
func (t lazyST) query(o, l, r int) int64 {
	if l <= t[o].l && t[o].r <= r {
		return t[o].sum
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if m < l {
		return t.query(o<<1|1, l, r)
	}
	vl := t.query(o<<1, l, r)
	vr := t.query(o<<1|1, l, r)
	return t.op(vl, vr)
}

func (t lazyST) queryAll() int64 { return t[1].sum }

// a 从 0 开始
func newLazySegmentTree(a []int64) lazyST {
	t := make(lazyST, 4*len(a))
	t.build(a, 1, 1, len(a))
	return t
}
