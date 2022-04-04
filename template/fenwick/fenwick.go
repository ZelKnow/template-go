package template

type fenwick struct {
	tree []int
}

func newFenwick(n int) fenwick {
	return fenwick{make([]int, n+1)}
}

func (f fenwick) add(i int, val int) {
	for ; i < len(f.tree); i += i & -i {
		f.tree[i] += val
	}
}

func (f fenwick) sum(i int) (res int) {
	for ; i > 0; i -= i & -i {
		res += f.tree[i]
	}
	return
}

func (f fenwick) query(l, r int) int {
	return f.sum(r) - f.sum(l-1)
}