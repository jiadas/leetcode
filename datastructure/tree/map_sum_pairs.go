package main

type MapSum struct {
	children [26]*MapSum
	val      int
}

/** Initialize your data structure here. */
func MapSumConstructor() MapSum {
	return MapSum{}
}

func (this *MapSum) Insert(key string, val int) {
	if this == nil {
		return
	}

	p := this
	for i := 0; i < len(key); i++ {
		index := key[i] - 'a'
		if p.children[index] == nil {
			p.children[index] = new(MapSum)
		}
		p = p.children[index]
	}
	p.val = val
}

func (this *MapSum) Sum(prefix string) int {
	if this == nil {
		return 0
	}

	p := this
	for i := 0; i < len(prefix); i++ {
		index := prefix[i] - 'a'
		if p.children[index] == nil {
			return 0
		}
		p = p.children[index]
	}
	if p != nil {
		sum := p.val
		for _, c := range p.children {
			if c != nil {
				sum += c.Sum("")
			}
		}
		return sum
	}
	return 0
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
