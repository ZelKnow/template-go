package template

// 自定义条件
func cond(num int) bool {
	if num >= 0 {
		return true
	} else {
		return false
	}
}

// 找第一个满足条件的元素
func lowerBound(a []int) int {
	n := len(a)
	l, r := 0, n - 1
	for l < r {
	// for l <= r { // 如果数组中可能不存在满足条件的元素
		mid := (r - l) >> 1 + l
		if cond(a[mid]) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

// 找最后一个满足条件的元素
func UpperBound(a []int) int {
	n := len(a)
	l, r := 0, n - 1
	for l < r {
	// for l <= r {
		mid := (r - l + 1) >> 1 + l
		if cond(a[mid]) {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return r
}