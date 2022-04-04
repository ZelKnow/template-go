package template

type unionFind struct {
	fa	[]int
	groups	int
}

func newUnionFind(n int) unionFind {
	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	return unionFind{fa, n}
}

func (u unionFind) find(x int) int {
	if u.fa[x] != x {
		u.fa[x] = u.find(u.fa[x])
	}
	return u.fa[x]
}

func (u *unionFind) merge(from, to int) (isMerged bool) {
	x, y := u.find(from), u.find(to)
	if x == y {
		return false
	}
	u.fa[x] = y
	u.groups--
	return true
}

func (u unionFind) same(x, y int) bool {
	return u.find(x) == u.find(y)
}