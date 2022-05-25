package template

func computeNext(s []byte) []int {
	next := make([]int, len(s))
	for i, c := 1, 0;i<len(s);i++ {
		v := s[i]
		for c > 0 && v != s[c] {
			c = next[c-1]
		}
		if v == s[c] {
			c++
		}
		next[i] = c
	}
	return next
}

func kmp(text, pattern []byte) (pos []int) {
	next := computeNext(pattern)
	lenP := len(pattern)
	c := 0
	for i, v := range text {
		for c > 0 && pattern[c] != v {
			c = next[c-1]
		}
		if pattern[c] == v {
			c++
		}
		if c == lenP {
			pos = append(pos, i-lenP+1)
			c = next[c-1]
		}
	}
	return
}