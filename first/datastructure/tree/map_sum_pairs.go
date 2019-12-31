package main

type MapSum struct {
	children [26]*MapSum
	val      int
}

/** Initialize your data structure here. */
func MapSumConstructor() MapSum {
	return MapSum{}
}

// For the method insert, you'll be given a pair of (string, integer). The string represents the key and the integer represents the value.
// If the key already existed, then the original key-value pair will be overridden to the new one.
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

// For the method sum, you'll be given a string representing the prefix,
// and you need to return the sum of all the pairs' value whose key starts with the prefix.
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
