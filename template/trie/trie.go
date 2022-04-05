package template

type trieNode struct {
	children	[26]*trieNode
	isEnd		bool
}

func (o *trieNode) empty() bool {
	for _, child := range o.children {
		if child != nil {
			return false
		}
	}
	return true
}

type trie struct {
	root *trieNode
}

func newTrie() *trie {
	return &trie{&trieNode{}}
}

func ord(c byte) byte {
	return c - 'a'
}

func chr(v byte) byte {
	return v + 'a'
}

func (t *trie) put(s []byte) *trieNode {
	o := t.root
	for _, b := range s {
		b = ord(b)
		if o.children[b] == nil {
			o.children[b] = &trieNode{}
		}
		o = o.children[b]
	}
	o.isEnd = true
	return o
}

func (t *trie) find(s []byte) *trieNode {
	o := t.root
	for _, b := range s {
		o = o.children[ord(b)]
		// 未找到 s，且 s 不是任何字符串的前缀
		if o == nil {
			return nil
		}
	}
	// 未找到 s，但是 s 是某个字符串的前缀
	if !o.isEnd {
		return nil
	}
	return o
}

func (t *trie) delete(s []byte) *trieNode {
	fa := make([]*trieNode, len(s))
	o := t.root
	for i, b := range s {
		fa[i] = o
		o = o.children[ord(b)]
		if o == nil {
			return nil
		}
	}
	if !o.isEnd {
		return nil
	}
	o.isEnd = false
	ret := o
	for i := len(s) - 1; i >= 0; i-- {
		if !o.empty() {
			break
		}
		o = fa[i]
		o.children[ord(s[i])] = nil
	}
	return ret
}